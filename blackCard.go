package main

type BlackCard interface {
	IsSatisfySkillCondition(controller *Controller) bool
	UseSkill(controller *Controller) bool
	IsSatisfyEffectCondition(controller *Controller) bool
	TriggerEffect(controller *Controller) bool
}
