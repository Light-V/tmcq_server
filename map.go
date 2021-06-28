package main

import (
	"container/list"
	"main/define"
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
	GroundCards     *list.List
	QianJunType     int
	ZhongJunType    int
	XiaJunType      int

	CanAttackZhou bool

	CanAttack bool

	CanBeAttacked bool
}

func (this *Map) isValidTile(x int, y int) bool {
	if x >= 0 && x < this.Width && y >= 0 && y < this.Height {
		return true
	}

	return false
}

func (this *Map) GetTileAt(x, y int) *MapTile {
	return this.Tiles[x][y]
}

func (this *Map) GetWhiteCardFromTileTODeck(tile *MapTile, player *Player) {
	if tile.QianJunType != 0 {
		player.ReceiveWhiteCard(tile.QianJunType, 1)
		tile.QianJunType = 0
	}
	if tile.ZhongJunType != 0 {
		player.ReceiveWhiteCard(tile.ZhongJunType, 1)
		tile.ZhongJunType = 0
	}
	if tile.XiaJunType != 0 {
		player.ReceiveWhiteCard(tile.XiaJunType, 1)
		tile.XiaJunType = 0
	}

}
