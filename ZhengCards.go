package main

import (
	"main/define"
)

type ZhengZhuangGong struct {
	MakeChoiceMsg chan *Message
	UseBlackMsg   chan *Message
	HumanAndGroundBase
}

func NewThisCard() *ZhengZhuangGong {
	MKMsg := make(chan *Message, 10)
	UBMsg := make(chan *Message, 10)
	thisCard := &ZhengZhuangGong{MakeChoiceMsg: MKMsg, UseBlackMsg: UBMsg, HumanAndGroundBase: HumanAndGroundBase{
		ID:          0,
		Fight:       3,
		Gold:        7,
		Faith:       0,
		CanUseSkill: false,
		SkillTimes:  1,
		CanMove:     false,
		MoveSteps:   0,
		CardType:    define.B_C_JUN,
		Country:     define.C_ZHENG,
		CanPutToMap: true,
		X:           -1,
		Y:           -1,
	}}
	return thisCard
}

//TODO 实现Human接口的方法
func (this *ZhengZhuangGong) IsSatisfySkillCondition(controller *Controller) bool {
	zhengPlayer := controller.GetPlayerByCountry(define.C_ZHENG)
	if this.CanUseSkill && this.SkillTimes >= 1 && zhengPlayer.PlayerBoard.Money >= 2 {
		return true
	}
	return false
}

func (this *ZhengZhuangGong) UseSkill(controller *Controller) bool {
	if this.IsSatisfySkillCondition(controller) {
		zhengPlayer := controller.GetPlayerByCountry(define.C_ZHENG)

		/*TODO 选择发动技能选项 */

		msg1 := *(<-this.MakeChoiceMsg)
		makeChoiceMsg := msg1.(MakeChoiceMessage)
		choice := makeChoiceMsg.Choice
		if choice == 0 { //选择发动第一种技能效果 2金出一张非大臣卡
			if this.CanUseSkillEffectOne(controller, zhengPlayer) {

				/*TODO 选择一个非大臣卡*/
				msg2 := *(<-this.UseBlackMsg)
				useBlackCardMsg := msg2.(UseBlackMessage)
				bc := zhengPlayer.CheckBlackCardDeck(useBlackCardMsg.BlackCardID)
				cardType := bc.GetCardType()
				/*打出该大臣卡*/
				if bc.GetCardType() != define.B_C_CHEN {
					//非大臣人来卡或者地形卡
					if cardType == define.B_C_JUN || cardType == define.B_C_ZU || cardType == define.B_C_DI {
						mt := controller.MapData.GetTileAt(useBlackCardMsg.DestX, useBlackCardMsg.DestY)
						res := zhengPlayer.PutBlackCardToMap(mt, bc, controller)
						if res == true {
							this.CanUseSkill = false
							this.SkillTimes -= 1

							zhengPlayer.changeMoney(-2)

							thisMapTile := controller.MapData.GetTileAt(this.HumanAndGroundBase.X, this.HumanAndGroundBase.Y)
							thisMapTile.CanAttack = false
							return res

						} else /*TODO 无法将黑卡下在该地图格*/ {

							return res
						}
					} else if cardType == define.B_C_CE || cardType == define.B_C_QI || cardType == define.B_C_MENG {
						res := zhengPlayer.UseStrategyOrUnitOrToolCard(controller, bc)
						if res == true {
							this.CanUseSkill = false
							this.SkillTimes -= 1

							zhengPlayer.changeMoney(-2)

							thisMapTile := controller.MapData.GetTileAt(this.HumanAndGroundBase.X, this.HumanAndGroundBase.Y)
							thisMapTile.CanAttack = false
							return true
						} else /*TODO 不满足出该黑卡的条件*/ {

							return res
						}
					}

				}
			}

		} else if choice == 1 { //选择第二种技能效果 2金买一车

			zhengPlayer.CheNum += 1
			zhengPlayer.changeMoney(-2)
			thisMapTile := controller.MapData.GetTileAt(this.HumanAndGroundBase.X, this.HumanAndGroundBase.Y)
			thisMapTile.CanAttack = false

			return true

		}

	}

	return false
}

func (this *ZhengZhuangGong) CanUseSkillEffectOne(controller *Controller, zhengPlayer *Player) bool {
	//判断是否能使用第一种技能效果
	count := 0
	for i := zhengPlayer.BlackCardsDeck.Front(); i != nil; i = i.Next() {
		if i.Value.(BlackCard).GetCardType() != define.B_C_CHEN {
			count++
		}
	}

	return count >= 1

}

func (this *ZhengZhuangGong) IsSatisfyEffectCondition(controller *Controller) bool {

	return false
}

//发动被动效果
func (this *ZhengZhuangGong) TriggerEffect(controller *Controller) {
	if this.IsSatisfyEffectCondition(controller) {
		//TODO

	}
	return

}

func (this *ZhengZhuangGong) GetCardType() int {

	return this.HumanAndGroundBase.CardType
}

func (this *ZhengZhuangGong) TriggerLeaveMapEffect(controller *Controller) {

}

func (this *ZhengZhuangGong) TriggerEnterAlterEffect(controller *Controller) {

}

func (this *ZhengZhuangGong) GetGold(controller *Controller) {
	players := controller.Players

	for i := 0; i < len(players); i++ {
		if players[i].Country == this.Country {
			players[i].PlayerBoard.Money += this.HumanAndGroundBase.Gold
		}
	}
}
