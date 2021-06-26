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

// Select 选择势力&都城&君主
func (this *Player) Select(c *Controller) {
	//todo
	if this.ID == 0 {
		//势力
		this.Country = define.C_QIN
		//都城
		this.Capital = c.MapData.GetTileAt(0, 0)
		//TODO 君主

	} else {
		//势力
		this.Country = define.C_ZHENG
		//都城
		this.Capital = c.MapData.GetTileAt(0, 4)
		//TODO 君主

	}

}

func (this *Player) PutHumanCardToMap(mapTile *MapTile, humanCard *HumanCard) {

}

func (this *Player) HumanCardHelper(mapTile *MapTile, humanCard *HumanCard) {
	//输入：mapTile可放置人类黑卡的地图格   blackCard人类黑卡
	mapTile.HumanCards.PushBack(humanCard)
	mapTile.RemainPutHuman -= 1
	mapTile.OwnerCountry = this.Country
	humanCard.CanMove = true
	humanCard.MoveSteps = 1
	humanCard.CanUseSkill = true
	humanCard.X = mapTile.X
	humanCard.Y = mapTile.Y
}

func (this *Player) PutGroundCardToMap(mapTile *MapTile, groundCard *GroundCard) {

}

func (this *Player) GroundCardHelper(mapTile *MapTile, groundCard *GroundCard) {
	//输入 mapTile可放置地形黑卡的地图格   groundCard地形黑卡
	mapTile.GroundCards.PushBack(groundCard)
	mapTile.RemainPutGround -= 1
	groundCard.CanMove = false
	groundCard.MoveSteps = 0
	groundCard.CanUseSkill = true
	groundCard.X = mapTile.X
	groundCard.Y = mapTile.Y
}

func (this *Player) BuyWhiteCard(cardType int, number int) bool {
	var cost int = cardType * number
	if this.PlayerBoard.Money >= cost {
		for i := 0; i < number; i++ {
			this.WhiteCardsDeck.PushBack(WhiteCard{cardType, cardType, cardType})
		}
		this.PlayerBoard.Money -= cost
		return true
	}

	return false
}

func (this *Player) MoveBlackCardToMapTile(mapTile *MapTile, blackCard *BlackCard) {
	//输入：
}
