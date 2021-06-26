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
				-1, define.C_EMPTY, nil, nil, nil,
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
	WhiteCards      *list.List

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
