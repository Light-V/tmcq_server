package main

import (
	"container/list"
	"main/define"
	"main/logger"
)

const (
	PlayerNum = 4
)

type Controller struct {
	Round              int
	Players            []*Player
	Countries          []int
	AvailableCountries []int
	DestiniesInPool    *DestinyPool
	MapData            *Map
	SeqOfCurrentRound  []int

	// 消息队列
	MessageChan chan *Message

	// List of Register
	AtBeginOfYear         *list.List
	AtEndOfYear           *list.List
	AtBeginOfDestiny      *list.List
	BeforeActiveOfDestiny *list.List
	AfterActiveOfDestiny  *list.List
	BeforeTax             *list.List
	AfterTax              *list.List
	AtEndOfDestiny        *list.List
	AtBeginOfRenshi       *list.List
	AtEndOfRenshi         *list.List
	BeforeAttack          *list.List
	AfterAttack           *list.List
}

func (this *Controller) Run() {
	logger.GetLogger().Println("元年开始")
	logger.GetLogger().Println("玩家选择势力,都城,君主")
	//元年
	for _, player := range this.Players {
		logger.GetLogger().Printf("玩家%d开始选择\n", player.ID)
		player.Select(this)
		logger.GetLogger().Printf("玩家%d选择完成\n", player.ID)
	}
	//正常循环
	for {
		this.RunNormalYear()
	}
}

func (this *Controller) RunNormalYear() {
	logger.GetLogger().Printf("第%d年开始\n", this.Round)
	//天时
	//抽取天命
	logger.GetLogger().Println("天时阶段开始")
	destiny := this.DestiniesInPool.GetNextDestiny()
	//发动天命
	logger.GetLogger().Println("天命生效开始")
	destiny.Activate()
	logger.GetLogger().Println("天命生效结束")
	logger.GetLogger().Println("天时阶段结束")
	logger.GetLogger().Println("地利阶段开始")
	//地利
	for i := 0; i < PlayerNum; i++ {
		logger.GetLogger().Printf("玩家%d获取税金开始", i)
		this.GetTax(i)
		logger.GetLogger().Printf("玩家%d获取税金结束", i)
	}
	logger.GetLogger().Println("地利阶段结束")
	//人事
	logger.GetLogger().Println("人事阶段开始")
	for idx, id := range this.SeqOfCurrentRound {
		currentPlayer := this.Players[id]
		currentStage := define.STAGE_XINGZHENG
		logger.GetLogger().Printf("本年第%d个人事阶段开始,属于玩家%d", idx, id)
		// 默认未进贡
		hasJinGongOrPunished := false
		// 默认未执行出黑/买白
		execBlackOrWhite := define.NOT_EXEC_BLAKC_OR_WHITE

		var msg Message
		for {
			msg = *(<-this.MessageChan)
			msgType := msg.GetType()
			// 进贡阶段
			if msgType == MSG_JINGONG {
				if hasJinGongOrPunished {
					// todo 通知客户端
					logger.GetLogger().Printf("玩家%d进贡失败，重复进贡或已受惩罚", id)
				}
				jm := msg.(JinGongMessage)
				amount := jm.Amount
				// 进贡0金视为直接接受惩罚
				if amount == 0 {
					// todo 惩罚
					continue
				}
				res := currentPlayer.JinGong(amount)
				if res {
					//todo 通知客户端
					logger.GetLogger().Printf("玩家%d进贡%金成功", id, amount)
				} else {
					//todo 通知客户端
					logger.GetLogger().Printf("玩家%d进贡%金失败", id, amount)
				}
				hasJinGongOrPunished = true
				continue
			}

			// 未进贡惩罚
			if !hasJinGongOrPunished {
				logger.GetLogger().Printf("玩家%d由于未进贡将接受惩罚", id)
				// todo 惩罚
				hasJinGongOrPunished = true
			}
			// 行政
			if currentStage <= define.STAGE_XINGZHENG {
				// 进入行政阶段
				switch msg.GetType() {
				case MSG_USEBLACKCARD:
					// 已经买白,无法出黑
					if execBlackOrWhite == define.WHITE_STAGE {
						logger.GetLogger().Printf("本回合已经买白,无法出黑")
						//todo 通知客户端
						continue
					}
					um := msg.(UseBlackMessage)
					bc := currentPlayer.CheckBlackCardDeck(um.BlackCardID)
					mt := this.MapData.GetTileAt(um.DestX, um.DestY)
					currentPlayer.PutBlackCardToMap(mt, bc, this)
					execBlackOrWhite = define.BLACK_STAGE
					// 执行成功进入行政阶段
					currentStage = define.STAGE_XINGZHENG

				case MSG_BUYWHITECARD:
					// 已经出黑,无法买白
					if execBlackOrWhite == define.BLACK_STAGE {
						logger.GetLogger().Printf("本回合已经出黑,无法买白")
						//todo 通知客户端
						continue
					}
					bw := msg.(BuyWhiteMessage)
					res := currentPlayer.BuyWhiteCard(bw.WhiteCardType, bw.WhiteCardNum)
					if res {
						// 购买成功
						execBlackOrWhite = define.WHITE_STAGE
						// todo 通知客户端
					} else {
						// 购买失败
						// todo 通知客户端
						continue
					}
					// 执行成功进入行政阶段
					currentStage = define.STAGE_XINGZHENG

				default:
					break
				}
			}
			// 调遣阶段
			if currentStage <= define.STAGE_DIAOQIAN {
				switch msg.GetType() {

				}
				// 执行成功 进入调遣阶段
				currentStage = define.STAGE_DIAOQIAN
			}

			// 休整阶段
			if currentStage <= define.STAGE_XIUZHENG {
				switch msg.GetType() {
				case MSG_BUYWEAPON:
					bw := msg.(BuyWeaponMessage)
					res := currentPlayer.BuyWeapon(bw.WeaponType)
					if !res {

						//购买失败
						break
					}
					//执行成功 进入修整阶段
					currentStage = define.STAGE_XINGZHENG
					continue
				case MSG_BUYSHIQI:
					// todo
					//执行成功 进入修整阶段
					currentStage = define.STAGE_XINGZHENG
				case MSG_SELLBLACKCARD:
					// todo
					//执行成功 进入修整阶段
					currentStage = define.STAGE_XINGZHENG
				case MSG_SELLWEAPON:
					sw := msg.(SellWeaponMessage)
					res := currentPlayer.SellWeapon(sw.WeaponType)
					if !res {

						//出售失败
						break
					}
					//执行成功 进入修整阶段
					currentStage = define.STAGE_XINGZHENG
					continue
				default:

				}

			}

		}

		logger.GetLogger().Printf("本年第%d个人事阶段结束,属于玩家%d", idx, id)
	}

	logger.GetLogger().Println("人事阶段结束")

	logger.GetLogger().Printf("第%d年结束\n", this.Round)
	//下一年
	this.Round++
}

func (this *Controller) GetTax(playerID int) int {
	tax := define.BASE_TAX
	//TODO
	return tax
}

func NewController() *Controller {
	availableCountries := []int{define.C_QIN, define.C_ZHENG}
	mapData := NewMap()
	destinyPool := NewDestinyPool()
	players := make([]*Player, PlayerNum)
	for i := 0; i < PlayerNum; i++ {
		players[i] = &Player{
			ID:             i,
			Country:        define.C_EMPTY,
			Capital:        nil,
			BlackCardsDeck: nil,
			GongNum:        0,
			BuNum:          0,
			QiNum:          0,
			CheNum:         0,
			//BlackCardsInMap: nil,
			CanUseBlackCard: false,
			CanBuyWhite:     false,

			PlayerBoard: PlayerBoard{
				Money:                  0,
				Morale:                 0,
				MyBlackCardsInAlter:    nil,
				OtherBlackCardsInAlter: nil,
				Tools:                  nil,
			},
		}
	}
	//初始化消息队列
	MessageChan := make(chan *Message, 100)
	SeqOfInitialTurn := make([]int, PlayerNum)
	for i := 0; i < PlayerNum; i++ {
		SeqOfInitialTurn[i] = i
	}
	c := &Controller{
		Round:              1,
		SeqOfCurrentRound:  SeqOfInitialTurn,
		MapData:            mapData,
		MessageChan:        MessageChan,
		Players:            players,
		DestiniesInPool:    destinyPool,
		AvailableCountries: availableCountries,
		Countries:          availableCountries,
	}
	return c
}
