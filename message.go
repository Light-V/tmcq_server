package main

const (
	MSG_SEND = iota
	MSG_JINGONG
	MSG_USEBLACKCARD
	MSG_BUYWHITECARD
	MSG_BUYWEAPON
	MSG_BUYSHIQI
	MSG_SELLWEAPON
	MSG_SELLBLACKCARD
	MSG_MAKECHOICE
)

type Message interface {
	GetType() int
}

type StMessage struct {
	Type int
}

func (m *StMessage) GetType() int {
	return m.Type
}

type JinGongMessage struct {
	Message
	Amount int
}

// 行政阶段

type UseBlackMessage struct {
	Message
	BlackCardID int
	DestX       int
	DestY       int
}

type BuyWhiteMessage struct {
	Message
	WhiteCardType int
	WhiteCardNum  int
}

// 休整阶段

type BuyWeaponMessage struct {
	Message
	WeaponType int
}

type BuyShiQiMessage struct {
	Message
}

type SellWeaponMessage struct {
	Message
	WeaponType int
}

type SellBlackCardMessage struct {
	Message
}

type MakeChoiceMessage struct {
	Message
	Choice int
}
