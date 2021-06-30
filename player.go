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
	//BlackCardsDeck []BlackCard
	BlackCardsDeck *list.List

	//白卡手牌
	BuNum   int
	GongNum int
	QiNum   int
	CheNum  int

	//该玩家在地图格上的黑卡
	//BlackCardsInMap []BlackCard

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
		this.Lord = c.MapData.GetTileAt(0, 0)
	} else {
		//势力
		this.Country = define.C_ZHENG
		//都城
		this.Capital = c.MapData.GetTileAt(0, 4)
		//TODO 君主
		this.Lord = c.MapData.GetTileAt(0, 4)

	}

}

func (this *Player) PutBlackCardToMap(mapTile *MapTile, blackCard BlackCard, controller *Controller) bool {
	//
	var cardType int = blackCard.GetCardType()

	if cardType == define.B_C_JUN { //人类黑卡
		newLord, isHuman := blackCard.(HumanCard)
		if isHuman {

			if true /*TODO 还有许多条件判定需要增加*/ {
				return this.PutKingHelper(mapTile, newLord, controller)
			}

		}

	} else if cardType == define.B_C_CHEN {
		humanCard, isHuman := blackCard.(HumanCard)
		if isHuman {
			if true /*TODO 还有许多条件判定需要增加*/ {
				this.PutHumanCardHelper(mapTile, humanCard)
				humanCard.GetGold(controller)
			}
		}

	} else if cardType == define.B_C_ZU {
		humanCard, isHuman := blackCard.(HumanCard)
		if isHuman {
			if true /*TODO 还有许多条件判定需要增加*/ {
				this.PutHumanCardHelper(mapTile, humanCard)
				humanCard.GetGold(controller)
			}
		}

	} else if cardType == define.B_C_DI { //地形卡
		groundCard, isGround := blackCard.(GroundCard)
		if isGround {
			this.PutGroundCardHelper(mapTile, groundCard)
			if true /*TODO 还有许多条件判定需要增加*/ {
				this.PutGroundCardHelper(mapTile, groundCard)
			}
		}
	}
	return false
}

func (this *Player) PutKingHelper(mapTile *MapTile, newLord HumanCard, controller *Controller) bool {
	if mapTile != this.Lord && mapTile != this.Capital { //选择的地图格与当前君主的地图格不相等并且与当前国都的地图格也不相等
		return false
	} else if mapTile == this.Lord { //选择的地图格就是旧国君所在地图格
		for i := this.Lord.HumanCards.Front(); i != nil; i = i.Next() { //从该地图格中删除旧国君
			if i.Value.(BlackCard).GetCardType() == define.B_C_JUN {
				oldLord := i.Value.(BlackCard) //旧国君
				oldLord_human, _ := oldLord.(HumanCard)
				oldLord_base := oldLord_human.(*HumanAndGroundBase)

				//TODO 发动旧国君退场的效果
				oldLord.TriggerLeaveMapEffect(controller)

				//进入祭坛，  属性复原

				this.PlayerBoard.MyBlackCardsInAlter.PushBack(oldLord)

				oldLord_base.CanMove = false
				oldLord_base.MoveSteps = 0
				oldLord_base.CanUseSkill = false
				oldLord_base.X = -1
				oldLord_base.Y = -1

				//TODO 发动进入祭坛效果
				oldLord.TriggerEnterAlterEffect(controller)

				//地图格操作
				this.Lord.HumanCards.Remove(i)
				this.Lord.RemainPutHuman += 1
				break

			}
		}

		//TODO 新国君打入场上并发动入场效果
		this.PutHumanCardHelper(mapTile, newLord)
		newLord.TriggerEffect(controller)
		newLord.GetGold(controller)
		//this.Lord = mapTile
		return true
	} else if mapTile == this.Capital { //选择的地图格是当前国都的地图格
		if mapTile.RemainPutHuman < 1 {
			return false
		}

		oldLordTile := this.Lord
		for i := oldLordTile.HumanCards.Front(); i != nil; i = i.Next() { //从该地图格中删除旧国君
			if i.Value.(BlackCard).GetCardType() == define.B_C_JUN {
				oldLord := i.Value.(BlackCard) //旧国君
				oldLord_human, _ := oldLord.(HumanCard)
				oldLord_base := oldLord_human.(*HumanAndGroundBase)

				//TODO 发动旧国君退场的效果
				oldLord.TriggerLeaveMapEffect(controller)

				//进入祭坛，  属性复原

				this.PlayerBoard.MyBlackCardsInAlter.PushBack(oldLord)

				oldLord_base.CanMove = false
				oldLord_base.MoveSteps = 0
				oldLord_base.CanUseSkill = false
				oldLord_base.X = -1
				oldLord_base.Y = -1

				//TODO 发动进入祭坛效果
				oldLord.TriggerEnterAlterEffect(controller)

				//地图格操作
				oldLordTile.HumanCards.Remove(i)              //剔除旧国君
				mapTile.RemainPutHuman += 1                   //可放置人类卡数量+1
				oldLordTile.UpdateMapTileOwner()              //更新旧君主所在地图格主人
				if oldLordTile.OwnerCountry != this.Country { //如果旧君主所在的地图格不属于我，那么将该地图格白卡收回我的手牌
					oldLordTile.GetWhiteCardFromThisTileTODeck(this)
				}
				break

			}
		}

		//TODO 新国君打入场上并发动入场效果
		this.PutHumanCardHelper(mapTile, newLord)
		newLord.TriggerEffect(controller)
		newLord.GetGold(controller)
		this.Lord = mapTile
		return true

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

func (this *Player) UseStrategyOrUnitOrToolCard(controller *Controller, blackCard BlackCard) bool {

	if true /*TODO 补充出黑判定*/ {
		blackCard.TriggerEffect(controller)
		blackCard.TriggerEnterAlterEffect(controller)
		return true
	}
	return false
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

func (this *Player) MoveBlackCardToOtherMapTile(otherMapTile *MapTile, blackCard BlackCard, controller *Controller) {
	cardType := blackCard.GetCardType()
	if true /*TODO 是否能够移动到目标地图格的相关判定*/ {
		if cardType == 1 || cardType == 2 || cardType == 3 {
			humanCard, _ := blackCard.(HumanCard)
			humanBase := humanCard.(*HumanAndGroundBase)
			currMapTile := controller.MapData.GetTileAt(humanBase.X, humanBase.Y)

			//TODO 旧地图格相关修改
			currMapTile.RemainPutHuman -= 1

		}
	}

}

func (this *Player) MoveCapital(maptile *MapTile) {

	this.Capital = maptile
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
