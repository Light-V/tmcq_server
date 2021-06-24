package main

type BlackCard interface {
	UseSkill(controller *Controller) bool
	IsSatisfyCondition(controller *Controller) bool
	TriggerEffect(controller *Controller) bool
}
