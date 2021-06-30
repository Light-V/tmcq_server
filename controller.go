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
		logger.GetLogger().Printf("第%d年开始\n", this.Round)
		this.RunNormalYear()
		logger.GetLogger().Printf("第%d年结束\n", this.Round)
	}
}

func (this *Controller) RunNormalYear() {
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
	logger.GetLogger().Println("人事阶段结束")
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
			BuNum:          0,
			GongNum:        0,
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
	round := 1
	c := &Controller{
		Round:              round,
		MapData:            mapData,
		Players:            players,
		DestiniesInPool:    destinyPool,
		AvailableCountries: availableCountries,
		Countries:          availableCountries,
	}
	return c
}

func test() {
	a := Controller{}
	go a.Run()
}
