package main

type BlackCard interface {
	IsSatisfySkillCondition(controller *Controller) bool
	UseSkill(controller *Controller) bool
	IsSatisfyEffectCondition(controller *Controller) bool
	TriggerEffect(controller *Controller)
	TriggerLeaveMapEffect(controller *Controller)
	TriggerEnterAlterEffect(controller *Controller)
	GetCardType() int
	GetCardID() int
}

type HumanAndGroundBase struct {
	ID int
	//战力
	Fight int
	//进场给国库的金币
	Gold int
	//忠诚
	Faith int
	//是否能使用礼器
	CanUseSkill bool
	//礼器数量
	SkillTimes int
	//是否能够移动
	CanMove bool
	//可以移动的步数
	MoveSteps int
	//人类卡类型 君主 族裔 大臣
	CardType int
	//国家
	Country int
	//是否能打到地图上 烛之武 要离等不可直接打到地图上
	CanPutToMap bool
	//下到地图上之后的坐标
	X int
	Y int
}

func (this *HumanAndGroundBase) IsSatisfySkillCondition(controller *Controller) bool {

	return false
}

func (this *HumanAndGroundBase) UseSkill(controller *Controller) bool {

	return false
}

func (this *HumanAndGroundBase) IsSatisfyEffectCondition(controller *Controller) bool {

	return false
}

//发动被动效果
func (this *HumanAndGroundBase) TriggerEffect(controller *Controller) {
	if this.IsSatisfyEffectCondition(controller) {

	}

}

func (this *HumanAndGroundBase) TriggerLeaveMapEffect(controller *Controller) {

}

func (this *HumanAndGroundBase) TriggerEnterAlterEffect(controller *Controller) {

}

func (this *HumanAndGroundBase) GetCardType() int {

	return this.CardType
}

func (this *HumanAndGroundBase) GetCardID() int {

	return this.ID
}

func (this *HumanAndGroundBase) GetGold(controller *Controller) {

}

type HumanCard interface {
	GetGold(controller *Controller)

	BlackCard
}

type GroundCard interface {
	BlackCard
}
