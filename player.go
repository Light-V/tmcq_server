package main

import (
	"container/list"
	"main/define"
	"main/logger"
	"main/util"
	"math"
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
	CanUseBlackCardNum int

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

//TODO 元年相关

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

//TODO 进贡相关

func (this *Player) JinGong(amount int) bool {
	if amount <= 0 || amount > this.PlayerBoard.Money {
		return false
	} else {
		this.changeMoney(-amount)
	}
	return true
}

//TODO 出黑相关

func (this *Player) CanUseBlack(controller *Controller) bool {
	if this.CanUseBlackCardNum > 0 {

		return true
	}

	return false
}

func (this *Player) PutBlackCardToMap(mapTile *MapTile, blackCard BlackCard, controller *Controller) bool {
	//
	var cardType int = blackCard.GetCardType()
	mapData := controller.MapData
	if !this.CanUseBlack(controller) {
		return false
	}

	if cardType == define.B_C_JUN { //人类黑卡
		newLord, isHuman := blackCard.(HumanCard)
		if isHuman {

			if true /*TODO 还有许多条件判定需要增加*/ {
				return this.PutKingHelper(mapTile, newLord, controller)
			}

		}

	} else if cardType == define.B_C_CHEN || cardType == define.B_C_ZU {
		humanCard, isHuman := blackCard.(HumanCard)
		if isHuman {
			if mapData.HaveMyNeighTile(mapTile, this.Country) /*TODO 还有许多条件判定需要增加*/ {
				if mapTile.OwnerCountry == define.C_EMPTY /*空地图格*/ {
					this.PutHumanCardHelper(mapTile, humanCard)
					humanCard.TriggerEffect(controller)
					humanCard.GetGold(controller)
					return true

				} else if mapTile.HumanCards.Len() == 0 && mapTile.GroundCards.Len() == 1 /*有地形卡无人类卡*/ {
					if (mapTile.QianJunType != 0 || mapTile.ZhongJunType != 0 || mapTile.XiaJunType != 0) &&
						mapTile.OwnerCountry != this.Country { /*有白卡并且不是属于自己的地图格*/

						//不能下在该地图格
						return false
					} else { /*没有白卡 或者 该地图格不属于自己   此时由于该地图格只有一张地形卡，所以也可以下*/
						if util.GetListElementAt(mapTile.GroundCards, 0).Value.(BlackCard).GetCardType() == define.B_C_DU {
							//若该卡为国都卡，则不能下
							return false
						}
						this.PutHumanCardHelper(mapTile, humanCard)
						humanCard.TriggerEffect(controller)
						humanCard.GetGold(controller)

						return true
					}

				} else if mapTile.HumanCards.Len() > 0 && mapTile.GroundCards.Len() == 1 /*有地形卡有人类卡*/ {
					if mapTile.RemainPutHuman > 0 && mapTile.OwnerCountry == this.Country {
						/*该地图格可以容纳多一张人类黑卡并且该地图格属于自己*/
						this.PutHumanCardHelper(mapTile, humanCard)
						humanCard.TriggerEffect(controller)
						humanCard.GetGold(controller)

						return true
					}
				}
			}
		}

	} else if cardType == define.B_C_DI { //地形卡
		groundCard, isGround := blackCard.(GroundCard)
		if isGround {
			if mapData.HaveMyNeighTile(mapTile, this.Country) /*TODO 还有许多条件判定需要增加*/ {
				if mapTile.OwnerCountry == define.C_EMPTY /*空地图格*/ {
					this.PutGroundCardHelper(mapTile, groundCard)
					groundCard.TriggerEffect(controller)
					return true

				} else if mapTile.HumanCards.Len() > 0 && mapTile.GroundCards.Len() == 0 /*有人类卡无地形卡*/ {
					if mapTile.OwnerCountry == this.Country && mapTile.RemainPutGround > 0 {
						this.PutGroundCardHelper(mapTile, groundCard)
						groundCard.TriggerEffect(controller)
						return true
					}

				}
			}
		}
	}
	return false
}

/*出人类卡帮助函数 出国君情况*/

func (this *Player) PutKingHelper(mapTile *MapTile, newLord HumanCard, controller *Controller) bool {
	if mapTile != this.Lord && mapTile != this.Capital { //选择的地图格与当前君主的地图格不相等并且与当前国都的地图格也不相等
		return false
	} else if mapTile == this.Lord { //选择的地图格就是旧国君所在地图格
		for i := this.Lord.HumanCards.Front(); i != nil; i = i.Next() { //从该地图格中删除旧国君
			if i.Value.(BlackCard).GetCardType() == define.B_C_JUN {
				oldLord := i.Value.(BlackCard) //旧国君
				oldLord_human, _ := oldLord.(HumanCard)
				oldLord_base := oldLord_human.(*HumanAndGroundBase)

				//发动旧国君退场的效果
				oldLord.TriggerLeaveMapEffect(controller)

				//进入祭坛，  属性复原

				this.PlayerBoard.MyBlackCardsInAlter.PushBack(oldLord)

				oldLord_base.CanMove = false
				oldLord_base.MoveSteps = 0
				oldLord_base.CanUseSkill = false
				oldLord_base.X = -1
				oldLord_base.Y = -1

				//发动进入祭坛效果
				oldLord.TriggerEnterAlterEffect(controller)

				//地图格操作
				this.Lord.HumanCards.Remove(i)
				this.Lord.RemainPutHuman += 1
				break

			}
		}

		//新国君打入场上并发动入场效果
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

				//发动旧国君退场的效果
				oldLord.TriggerLeaveMapEffect(controller)

				//进入祭坛，  属性复原

				this.PlayerBoard.MyBlackCardsInAlter.PushBack(oldLord)

				oldLord_base.CanMove = false
				oldLord_base.MoveSteps = 0
				oldLord_base.CanUseSkill = false
				oldLord_base.X = -1
				oldLord_base.Y = -1

				//发动进入祭坛效果
				oldLord.TriggerEnterAlterEffect(controller)

				//旧国君所在地图格操作
				oldLordTile.HumanCards.Remove(i)              //剔除旧国君
				oldLordTile.RemainPutHuman += 1               //可放置人类卡数量+1
				oldLordTile.CanAttack = false                 //旧地图格不可进攻
				oldLordTile.UpdateMapTileOwner()              //更新旧君主所在地图格主人
				if oldLordTile.OwnerCountry != this.Country { //如果旧君主所在的地图格不属于我，那么将该地图格白卡收回我的手牌
					oldLordTile.GetWhiteCardFromThisTileTODeck(this)
				}
				break

			}
		}

		//新国君打入场上并发动入场效果
		this.PutHumanCardHelper(mapTile, newLord)
		newLord.TriggerEffect(controller)
		newLord.GetGold(controller)
		this.Lord = mapTile
		return true

	}

	return false

}

/*出人类卡帮助函数 出非国君情况*/

func (this *Player) PutHumanCardHelper(mapTile *MapTile, humanCard HumanCard) {
	logger.GetLogger().Println("")
	//输入：mapTile可放置人类黑卡的地图格   blackCard人类黑卡

	HumanAndGroundBase := humanCard.(*HumanAndGroundBase)
	mapTile.HumanCards.PushBack(humanCard)
	mapTile.RemainPutHuman -= 1
	mapTile.OwnerCountry = this.Country //下人类黑卡到一个地图格，该地图格的拥有者一定是该人类黑卡对应的国家
	mapTile.CanAttack = false

	HumanAndGroundBase.CanMove = true
	HumanAndGroundBase.MoveSteps = 1
	HumanAndGroundBase.CanUseSkill = true
	HumanAndGroundBase.X = mapTile.X
	HumanAndGroundBase.Y = mapTile.Y

	this.CanUseBlackCardNum -= 1

}

func (this *Player) PutGroundCardHelper(mapTile *MapTile, groundCard GroundCard) {
	//输入 mapTile可放置地形黑卡的地图格   groundCard地形黑卡
	HumanAndGroundBase := groundCard.(*HumanAndGroundBase)
	mapTile.GroundCards.PushBack(groundCard)
	mapTile.RemainPutGround -= 1
	mapTile.OwnerCountry = this.Country
	mapTile.CanAttack = false

	HumanAndGroundBase.CanMove = false
	HumanAndGroundBase.MoveSteps = 0
	HumanAndGroundBase.CanUseSkill = true
	HumanAndGroundBase.X = mapTile.X
	HumanAndGroundBase.Y = mapTile.Y

	this.CanUseBlackCardNum -= 1
}

func (this *Player) UseStrategyOrUnitOrToolCard(controller *Controller, blackCard BlackCard) bool {

	if true /*TODO 补充出黑判定*/ {
		blackCard.TriggerEffect(controller)
		blackCard.TriggerEnterAlterEffect(controller)
		this.CanUseBlackCardNum -= 1
		return true
	}
	return false
}

//TODO 买白相关

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
	} else {
		return
	}

}

//TODO 调遣相关

func (this *Player) MoveBlackCardToNewMapTile(newMapTile *MapTile, blackCard BlackCard, controller *Controller) bool {
	cardType := blackCard.GetCardType()

	if cardType == 1 || cardType == 2 || cardType == 3 { //TODO 人类卡
		humanCard, _ := blackCard.(HumanCard)
		humanBase := humanCard.(*HumanAndGroundBase)
		mapData := controller.MapData
		currMapTile := mapData.GetTileAt(humanBase.X, humanBase.Y)
		needSteps := int(math.Abs(float64(newMapTile.X-currMapTile.X)) + math.Abs(float64(newMapTile.Y-currMapTile.Y)))
		if humanBase.CanMove && needSteps <= humanBase.MoveSteps /*TODO 是否能够移动到目标地图格的相关判定 以及 某些黑卡移动条件的判定？（是否实现在这里）*/ {
			if newMapTile.OwnerCountry == define.C_EMPTY /*新格子为空地图格*/ {

				//TODO 旧地图格相关修改
				currMapTile.RemainPutHuman += 1
				currMapTile.HumanCards.Remove(GetBlackCardFromList(currMapTile.HumanCards, blackCard))
				currMapTile.CanAttack = false
				currMapTile.UpdateMapTileOwner()
				if currMapTile.OwnerCountry != this.Country {
					currMapTile.GetWhiteCardFromThisTileTODeck(this)
				}

				//TODO 新地图格相关修改
				newMapTile.RemainPutHuman -= 1
				newMapTile.HumanCards.PushBack(blackCard)
				newMapTile.OwnerCountry = this.Country
				newMapTile.CanAttack = false

				//TODO 该人类卡相关修改
				humanBase.X = newMapTile.X
				humanBase.Y = newMapTile.Y
				humanBase.CanUseSkill = false
				humanBase.MoveSteps -= needSteps
				if humanBase.MoveSteps <= 0 {
					humanBase.CanMove = false
				}

				return true

			} else if newMapTile.HumanCards.Len() == 0 && newMapTile.GroundCards.Len() == 1 /*新格子有地形卡无人类卡*/ {
				if (newMapTile.QianJunType != 0 || newMapTile.ZhongJunType != 0 || newMapTile.XiaJunType != 0) &&
					newMapTile.OwnerCountry != this.Country { /*有白卡并且不是属于自己的地图格*/

					//不能移动到该地图格
					return false
				} else { /*没有白卡 或者 该地图格不属于自己   此时由于该地图格只有一张地形卡，所以也可以移动到该地图格*/
					if util.GetListElementAt(newMapTile.GroundCards, 0).Value.(BlackCard).GetCardType() == define.B_C_DU {
						//若该卡为国都卡，则不能移动到该格
						return false
					}

					//TODO 旧地图格相关修改
					currMapTile.RemainPutHuman += 1
					currMapTile.HumanCards.Remove(GetBlackCardFromList(currMapTile.HumanCards, blackCard))
					currMapTile.CanAttack = false
					currMapTile.UpdateMapTileOwner()
					if currMapTile.OwnerCountry != this.Country {
						currMapTile.GetWhiteCardFromThisTileTODeck(this)
					}

					//TODO 新地图格相关修改
					newMapTile.RemainPutHuman -= 1
					newMapTile.HumanCards.PushBack(blackCard)
					newMapTile.OwnerCountry = this.Country
					newMapTile.CanAttack = false

					//TODO 该人类卡相关修改
					humanBase.X = newMapTile.X
					humanBase.Y = newMapTile.Y
					humanBase.CanUseSkill = false
					humanBase.MoveSteps -= needSteps
					if humanBase.MoveSteps <= 0 {
						humanBase.CanMove = false
					}

					return true
				}

			} else if newMapTile.HumanCards.Len() > 0 && newMapTile.GroundCards.Len() == 1 /*有地形卡有人类卡*/ {
				if newMapTile.RemainPutHuman > 0 && newMapTile.OwnerCountry == this.Country {
					/*该地图格可以容纳多一张人类黑卡并且  该地图格属于自己*/
					//TODO 旧地图格相关修改
					currMapTile.RemainPutHuman += 1
					currMapTile.HumanCards.Remove(GetBlackCardFromList(currMapTile.HumanCards, blackCard))
					currMapTile.CanAttack = false
					currMapTile.UpdateMapTileOwner()
					if currMapTile.OwnerCountry != this.Country {
						currMapTile.GetWhiteCardFromThisTileTODeck(this)
					}

					//TODO 新地图格相关修改
					newMapTile.RemainPutHuman -= 1
					newMapTile.HumanCards.PushBack(blackCard)
					newMapTile.OwnerCountry = this.Country
					newMapTile.CanAttack = false

					//TODO 该人类卡相关修改
					humanBase.X = newMapTile.X
					humanBase.Y = newMapTile.Y
					humanBase.CanUseSkill = false
					humanBase.MoveSteps -= needSteps
					if humanBase.MoveSteps <= 0 {
						humanBase.CanMove = false
					}

					return true
				}
			}
		}

	} else if cardType == define.B_C_DU { /*移动国都*/
		return this.MoveCapital(newMapTile, blackCard)

	}

	return false

}

//移动国都卡
func (this *Player) MoveCapital(newMapTile *MapTile, capital BlackCard) bool {
	capitalCard, _ := capital.(GroundCard)
	capitalBase := capitalCard.(*HumanAndGroundBase)
	oldCapital := this.Capital
	needSteps := int(math.Abs(float64(newMapTile.X-oldCapital.X)) + math.Abs(float64(newMapTile.Y-oldCapital.Y)))
	if capitalBase.CanMove && needSteps == 1 && this.PlayerBoard.Money >= 2 /*TODO 是否能够移动到目标地图格的相关判定*/ {
		if newMapTile.OwnerCountry == define.C_EMPTY /*新格子为空地图格*/ {

			//TODO 旧地图格相关修改
			oldCapital.RemainPutGround += 1
			oldCapital.HumanCards.Remove(GetBlackCardFromList(oldCapital.HumanCards, capital))
			oldCapital.CanAttack = false
			oldCapital.UpdateMapTileOwner()
			if oldCapital.OwnerCountry != this.Country {
				oldCapital.GetWhiteCardFromThisTileTODeck(this)
			}

			//TODO 新地图格相关修改
			newMapTile.RemainPutGround -= 1
			newMapTile.HumanCards.PushBack(capital)
			newMapTile.OwnerCountry = this.Country
			newMapTile.CanAttack = false

			//TODO 该国都卡相关修改
			capitalBase.X = newMapTile.X
			capitalBase.Y = newMapTile.Y
			capitalBase.CanUseSkill = false
			capitalBase.MoveSteps -= needSteps
			if capitalBase.MoveSteps <= 0 {
				capitalBase.CanMove = false
			}

			//TODO 玩家面板相关修改
			this.changeMoney(-2)

			this.Capital = newMapTile
			return true

		} else if newMapTile.HumanCards.Len() > 0 && newMapTile.GroundCards.Len() == 0 /*新格子有人类卡卡无地形卡*/ {
			if newMapTile.OwnerCountry == this.Country && newMapTile.RemainPutGround > 0 {

				//TODO 旧地图格相关修改
				oldCapital.RemainPutHuman += 1
				oldCapital.HumanCards.Remove(GetBlackCardFromList(oldCapital.HumanCards, capital))
				oldCapital.CanAttack = false
				oldCapital.UpdateMapTileOwner()
				if oldCapital.OwnerCountry != this.Country {
					oldCapital.GetWhiteCardFromThisTileTODeck(this)
				}

				//TODO 新地图格相关修改
				newMapTile.RemainPutHuman -= 1
				newMapTile.HumanCards.PushBack(capital)
				newMapTile.OwnerCountry = this.Country
				newMapTile.CanAttack = false

				//TODO 该国都卡相关修改
				capitalBase.X = newMapTile.X
				capitalBase.Y = newMapTile.Y
				capitalBase.CanUseSkill = false
				capitalBase.MoveSteps -= needSteps
				if capitalBase.MoveSteps <= 0 {
					capitalBase.CanMove = false
				}

				//TODO 玩家面板相关修改
				this.changeMoney(-2)

				this.Capital = newMapTile
				return true
			}

		}
	}

	return false
}

func (this *Player) PutWhiteCardToMapTile(whiteCardType int, mapTile *MapTile, location int) bool {
	if mapTile.canAddWhite /*TODO 还有许多其他判定条件*/ {
		if location == 1 { //放置前军
			this.ReceiveWhiteCard(mapTile.QianJunType, 1)
			mapTile.QianJunType = whiteCardType
			return true

		} else if location == 2 { //放置中军
			this.ReceiveWhiteCard(mapTile.ZhongJunType, 1)
			mapTile.ZhongJunType = whiteCardType
			return true

		} else if location == 3 { //放置下军
			this.ReceiveWhiteCard(mapTile.XiaJunType, 1)
			mapTile.XiaJunType = whiteCardType
			return true
		}

	}

	return false
}

func (this *Player) UseBlackCardSkill(blackCard BlackCard, controller *Controller) bool {

	return blackCard.UseSkill(controller)
}

func (this *Player) TriggerBlackCardEffect(blackCard BlackCard, controller *Controller) {
	blackCard.TriggerEffect(controller)
}

//TODO 休整相关

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

func GetBlackCardFromList(l *list.List, blackCard BlackCard) *list.Element {

	var i *list.Element
	for i = l.Front(); i != nil; i = i.Next() {
		if i.Value.(BlackCard).GetCardID() == blackCard.GetCardID() {
			return i
		}

	}
	return nil

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
