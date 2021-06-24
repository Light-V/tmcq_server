package main

import (
	"container/list"
	"main/define"
)

type Player struct {
	//玩家ID
	ID int
	//势力
	Country int
	//都城
	Capital *MapTile
	//君主

	//黑卡手牌
	BlackCardsDeck *list.List

	//白卡手牌
	WhiteCardsDeck *list.List

	//该玩家在地图格上的黑卡
	BlackCardsInMap *list.List

	//能否出黑
	CanUseBlackCard bool

	//能否买白
	CanBuyWhite bool

	//玩家面板
	PlayerBoard PlayerBoard
}

type PlayerBoard struct {

	//金钱
	Money int

	//士气
	Morale int

	//武器

	//祭坛中的己方黑卡
	MyBlackCardsInAlter *list.List

	//祭坛中的非己方黑卡
	OtherBlackCardsInAlter *list.List

	//间谍 越国 郑旦
	Spy *list.List
	//器物
	Tools *list.List
}

func (this *Player) GetBlackCards(County int) {
	//TODO 根据不同的国家来初始化玩家的手牌
}

func (this *Player) PutHumanCardToMap(mapTile *MapTile, blackCard *BlackCard) bool {

	return false
}

// Select 选择势力&都城&君主
func (this *Player) Select(c *Controller) {
	//todo
	if this.ID == 0 {
		//势力
		this.Country = define.C_QIN
		//都城
		this.Capital = c.MapData.GetTileAt(0, 0)
		//君主

	} else {
		//势力
		this.Country = define.C_ZHENG
		//都城
		this.Capital = c.MapData.GetTileAt(0, 4)
		//君主

	}

}
