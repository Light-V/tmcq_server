package main

type GroundCard struct {
	//战力
	Fight int
	//进场给国库的金币
	GetGold int
	//是否能使用礼器
	CanUseSkill bool
	//礼器数量
	SkillTimes int
	//是否能够移动
	CanMove bool
	//可以移动的步数
	MoveSteps int
	//地形卡
	CardType int
	//国家
	Country int
	//下到地图上之后的坐标
	X int
	Y int
}

func (this *HumanCard) IsSatisfySkillCondition(controller *Controller) bool {

	return false
}

func (this *HumanCard) UseSkill(controller *Controller) bool {

	return false
}

func (this *HumanCard) IsSatisfyEffectCondition(controller *Controller) bool {

	return false
}

//发动被动效果
func (this *GroundCard) TriggerEffect(controller *Controller) bool {
	if this.IsSatisfyEffectCondition(controller) {
		players := controller.Players

		for i := 0; i < len(players); i++ {
			if players[i].Country == this.Country {
				players[i].PlayerBoard.Money += this.GetGold
			}
		}

		//TODO

		return true
	}
	return false
}
