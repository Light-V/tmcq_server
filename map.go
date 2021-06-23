package main

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
			tiles[i][j] = &MapTile{"", i, j, -1}
		}
	}
	return &Map{5, 5, tiles}
}

type MapTile struct {
	Name    string
	X       int
	Y       int
	OwnerID int
}

func (this *Map) GetTileAt(x, y int) *MapTile {
	return this.Tiles[x][y]
}
