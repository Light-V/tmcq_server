package main

import (
	"container/list"
	"main/define"
	"main/util"
)

type Map struct {
	Height int
	Width  int
	Tiles  [][]*MapTile
}

func NewMap() *Map {
	tiles := make([][]*MapTile, 5)
	for i := 0; i < 5; i++ {
		tiles[i] = make([]*MapTile, 5)
		for j := 0; j < 5; j++ {
			tiles[i][j] = &MapTile{"", i, j, 0, 1, 1,
				-1, define.C_EMPTY, nil, nil, 0, 0, 0,
				false, false, false,
			}
		}
	}
	return &Map{5, 5, tiles}
}

type MapTile struct {
	Name string
	X    int
	Y    int

	//地图格类型 无 临山 临水 大漠  农田
	TileType        int
	RemainPutHuman  int
	RemainPutGround int
	OwnerID         int
	OwnerCountry    int
	HumanCards      *list.List
	GroundCards     *list.List //都城和地形卡都存放在这里
	//Traders			*list.List //宋国某些卡所在的地图格专用
	QianJunType  int
	ZhongJunType int
	XiaJunType   int

	//CanBePutBlacCard bool

	CanAttackZhou bool

	CanAttack bool

	canAddWhite bool //进攻后不可增兵

	//CanBeAttacked bool
}

func (this *Map) IsValidTile(x int, y int) bool {
	if x >= 0 && x < this.Width && y >= 0 && y < this.Height {
		return true
	}

	return false
}

func (this *Map) GetTileAt(x, y int) *MapTile {
	return this.Tiles[x][y]
}

func (this *MapTile) GetWhiteCardFromThisTileTODeck(player *Player) {
	if this.QianJunType != 0 {
		player.ReceiveWhiteCard(this.QianJunType, 1)
		this.QianJunType = 0
	}
	if this.ZhongJunType != 0 {
		player.ReceiveWhiteCard(this.ZhongJunType, 1)
		this.ZhongJunType = 0
	}
	if this.XiaJunType != 0 {
		player.ReceiveWhiteCard(this.XiaJunType, 1)
		this.XiaJunType = 0
	}

}

func (this *MapTile) UpdateMapTileOwner() {
	if this.HumanCards.Len() == 0 && this.GroundCards.Len() == 0 {
		this.OwnerCountry = define.C_EMPTY
	} else if this.HumanCards.Len() == 0 && this.GroundCards.Len() == 1 {
		this.OwnerCountry = util.GetListElementAt(this.GroundCards, 0).Value.(BlackCard).GetCountry()
	}

}

func (this *MapTile) UpdateMapTileRemainPut() {
	this.UpdateMapTileOwner()
	if this.OwnerCountry == define.C_EMPTY {
		this.RemainPutHuman = 1
		this.RemainPutGround = 1
	} else if this.OwnerID == define.C_QIN {
		this.RemainPutHuman = 3 - this.HumanCards.Len() - this.GroundCards.Len()
		this.RemainPutGround = 3 - this.HumanCards.Len() - this.GroundCards.Len()
	} else if true /*TODO 其他国家情况 */ {

	}

}

func (this *Map) HaveMyNeighTile(mapTile *MapTile, myCountry int) bool {
	x := mapTile.X
	y := mapTile.Y
	return (this.IsValidTile(x+1, y) && this.GetTileAt(x+1, y).OwnerCountry == myCountry) ||
		(this.IsValidTile(x-1, y) && this.GetTileAt(x-1, y).OwnerCountry == myCountry) ||
		(this.IsValidTile(x, y+1) && this.GetTileAt(x, y+1).OwnerCountry == myCountry) ||
		(this.IsValidTile(x, y-1) && this.GetTileAt(x, y-1).OwnerCountry == myCountry)

}

func (this *MapTile) IsThisTileCanBeAttacked() bool {
	//TODO  人类黑卡   地形卡  国都卡
	if this.HumanCards.Len() == 0 && this.GroundCards.Len() == 0 {
		return false
	}

	return false
}
