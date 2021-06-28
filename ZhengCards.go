package main

type ZhengZhuangGong struct {
	HumanAndGroundBase
}

//TODO 实现Human接口的方法
func (this *ZhengZhuangGong) IsSatisfySkillCondition(controller *Controller) bool {

	return false
}

func (this *ZhengZhuangGong) UseSkill(controller *Controller) bool {

	return false
}

func (this *ZhengZhuangGong) IsSatisfyEffectCondition(controller *Controller) bool {

	return false
}

//发动被动效果
func (this *ZhengZhuangGong) TriggerEffect(controller *Controller) bool {
	if this.IsSatisfyEffectCondition(controller) {

		//TODO

		return true
	}
	return false
}

func (this *ZhengZhuangGong) GetCardType() int {

	return this.HumanAndGroundBase.CardType
}

func (this *ZhengZhuangGong) GetGold(controller *Controller) {
	players := controller.Players

	for i := 0; i < len(players); i++ {
		if players[i].Country == this.Country {
			players[i].PlayerBoard.Money += this.HumanAndGroundBase.Gold
		}
	}
}
