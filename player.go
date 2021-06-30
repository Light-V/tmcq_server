package main

import (
	"container/list"
	"main/define"
	"main/logger"
)

type Player struct {
	//玩家ID
	ID int
	//势力
	Country int
	//都城所在地图格子
	Capital *MapTile
	//君主所在地图格子
	Lord *MapTile

	//黑卡手牌
	BlackCardsDeck *list.List

	//白卡手牌
	BuNum   int
	GongNum int
	QiNum   int
	CheNum  int

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

	//兵种武器
	BuWeaponNum   int
	GongWeaponNum int
	QiWeaponNum   int
	CheWeaponNum  int

	//祭坛中的己方黑卡
	MyBlackCardsInAlter *list.List

	//祭坛中的非己方黑卡
	OtherBlackCardsInAlter *list.List

	//间谍 越国 郑旦
	Spy *list.List
	//器物
	Tools *list.List
}

func (this Player) changeMoney(num int) {
	// todo
	this.PlayerBoard.Morale += num
}

func (this *Player) GetBlackCardsToMyDeck(County int) {
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

func (this *Player) PutBlackCardToMap(mapTile *MapTile, blackCard BlackCard, controller *Controller) bool {
	var cardType int = blackCard.GetCardType()
	if cardType == define.B_C_JUN { //人类黑卡
		humanCard, isHuman := blackCard.(HumanCard)
		if isHuman {
			//TODO 还有许多条件判定需要增加

			if mapTile != this.Lord && mapTile != this.Capital { //选择的地图格与当前君主的地图格
				return false
			} else if mapTile == this.Lord { //替换当前国君

				this.PutHumanCardHelper(mapTile, humanCard)
				humanCard.GetGold(controller)
				return true
			} else if mapTile == this.Capital {
				this.PutHumanCardHelper(mapTile, humanCard)
				humanCard.GetGold(controller)
				return true

			}

		}

	} else if cardType == define.B_C_CHEN {
		humanCard, isHuman := blackCard.(HumanCard)
		if isHuman {
			//TODO 还有许多条件判定需要增加
			this.PutHumanCardHelper(mapTile, humanCard)
		}

	} else if cardType == define.B_C_ZU {
		humanCard, isHuman := blackCard.(HumanCard)
		if isHuman {
			//TODO 还有许多条件判定需要增加
			this.PutHumanCardHelper(mapTile, humanCard)
		}

	} else if cardType == define.B_C_DI { //地形卡
		groundCard, isGround := blackCard.(GroundCard)
		if isGround {
			//TODO 还有许多条件判定需要增加
			this.PutGroundCardHelper(mapTile, groundCard)
		}
	}
	return false
}

func (this *Player) PutHumanCardHelper(mapTile *MapTile, humanCard HumanCard) {
	logger.GetLogger().Println("")
	//输入：mapTile可放置人类黑卡的地图格   blackCard人类黑卡
	HumanAndGroundBase := humanCard.(*HumanAndGroundBase)
	mapTile.HumanCards.PushBack(humanCard)
	mapTile.RemainPutHuman -= 1
	mapTile.OwnerCountry = this.Country
	HumanAndGroundBase.CanMove = true
	HumanAndGroundBase.MoveSteps = 1
	HumanAndGroundBase.CanUseSkill = true
	HumanAndGroundBase.X = mapTile.X
	HumanAndGroundBase.Y = mapTile.Y
}

func (this *Player) PutGroundCardHelper(mapTile *MapTile, groundCard GroundCard) {
	//输入 mapTile可放置地形黑卡的地图格   groundCard地形黑卡
	HumanAndGroundBase := groundCard.(*HumanAndGroundBase)
	mapTile.GroundCards.PushBack(groundCard)
	mapTile.RemainPutGround -= 1
	HumanAndGroundBase.CanMove = false
	HumanAndGroundBase.MoveSteps = 0
	HumanAndGroundBase.CanUseSkill = true
	HumanAndGroundBase.X = mapTile.X
	HumanAndGroundBase.Y = mapTile.Y
}

func (this *Player) BuyWhiteCard(cardType int, number int) bool {
	var cost int = cardType * number
	if this.PlayerBoard.Money >= cost {
		this.ReceiveWhiteCard(cardType, number)
		this.changeMoney(-cost)
		return true
	}

	return false
}

func (this *Player) ReceiveWhiteCard(cardType int, number int) {
	if cardType == define.W_C_BU {
		this.BuNum += number
	} else if cardType == define.W_C_GONG {
		this.GongNum += number
	} else if cardType == define.W_C_QI {
		this.QiNum += number
	} else if cardType == define.W_C_CHE {
		this.CheNum += number
	}

}

func (this *Player) MoveBlackCardToMapTile(mapTile *MapTile, blackCard *BlackCard) {
	//输入：
}

func (this *Player) JinGong(amount int) bool {
	if amount <= 0 || amount > this.PlayerBoard.Money {
		return false
	} else {
		this.changeMoney(-amount)
	}
	return true
}

func (this *Player) CheckBlackCardDeck(id int) BlackCard {
	for i := this.BlackCardsDeck.Front(); i != nil; i = i.Next() {
		bc := i.Value.(BlackCard)
		if bc.GetCardID() == id {
			return bc
		}
	}
	return nil
}

func (this *Player) BuyWeapon(weaponType int) bool {
	switch weaponType {
	case define.WEAPON_BU:
		if this.PlayerBoard.BuWeaponNum > 0 {
			return false
		}
		if this.PlayerBoard.Money >= define.PRICE_WEAPON_BU {
			this.changeMoney(-define.PRICE_WEAPON_BU)
			this.PlayerBoard.BuWeaponNum += 1
			return true
		}

	case define.WEAPON_GONG:
		if this.PlayerBoard.GongWeaponNum > 0 {
			return false
		}
		if this.PlayerBoard.Money >= define.PRICE_WEAPON_GONG {
			this.changeMoney(-define.PRICE_WEAPON_GONG)
			this.PlayerBoard.GongWeaponNum += 1
			return true
		}

	case define.WEAPON_QI:
		if this.PlayerBoard.QiWeaponNum > 0 {
			return false
		}
		if this.PlayerBoard.Money >= define.PRICE_WEAPON_QI {
			this.changeMoney(-define.PRICE_WEAPON_QI)
			this.PlayerBoard.QiWeaponNum += 1
			return true
		}

	case define.WEAPON_CHE:
		if this.PlayerBoard.CheWeaponNum > 0 {
			return false
		}
		if this.PlayerBoard.Money >= define.PRICE_WEAPON_CHE {
			this.changeMoney(-define.PRICE_WEAPON_CHE)
			this.PlayerBoard.CheWeaponNum += 1
			return true
		}

	default:
		return false

	}
	return true
}

func (this Player) SellWeapon(weaponType int) bool {
	switch weaponType {
	case define.WEAPON_BU:
		if this.PlayerBoard.BuWeaponNum == 0 {
			return false
		}
		if this.PlayerBoard.Money >= define.PRICE_WEAPON_BU {
			this.changeMoney(define.PRICE_WEAPON_BU)
			this.PlayerBoard.BuWeaponNum -= 1
			return true
		}

	case define.WEAPON_GONG:
		if this.PlayerBoard.GongWeaponNum == 0 {
			return false
		}
		if this.PlayerBoard.Money >= define.PRICE_WEAPON_GONG {
			this.changeMoney(define.PRICE_WEAPON_GONG)
			this.PlayerBoard.GongWeaponNum -= 1
			return true
		}

	case define.WEAPON_QI:
		if this.PlayerBoard.QiWeaponNum == 0 {
			return false
		}
		if this.PlayerBoard.Money >= define.PRICE_WEAPON_QI {
			this.changeMoney(define.PRICE_WEAPON_QI)
			this.PlayerBoard.QiWeaponNum -= 1
			return true
		}

	case define.WEAPON_CHE:
		if this.PlayerBoard.CheWeaponNum == 0 {
			return false
		}
		if this.PlayerBoard.Money >= define.PRICE_WEAPON_CHE {
			this.changeMoney(define.PRICE_WEAPON_CHE)
			this.PlayerBoard.CheWeaponNum -= 1
			return true
		}

	default:
		return false

	}
	return true
}
