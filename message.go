package main

const (
	MSG_SEND = iota
	MSG_JINGONG
	MSG_USEBLACKCARD
	MSG_BUYWHITECARD
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
