package main

import "main/define"

const (
	PlayerNum = 4
)

type Controller struct {
	Round              int
	Players            []*Player
	Countries          []int
	AvailableCountries []int
	DestiniesInPool    *DestinyPool
	MapData            *Map
}

func (this *Controller) Run() {
	//元年
	for _, player := range this.Players {
		player.Select(this)
	}
	//正常循环
	for {
		this.RunNormalYear()
	}
}

func (this *Controller) RunNormalYear() {
	//天时
	//抽取天命
	destiny := this.DestiniesInPool.GetNextDestiny()
	//发动天命
	destiny.Activate()
	//地利
	for i := 0; i < PlayerNum; i++ {
		this.GetTax(i)
	}
	//人事
}

func (this *Controller) GetTax(playerID int) int {
	tax := define.BASE_TAX
	for i := 0; i < this.MapData.Height; i++ {
		for j := 0; j < this.MapData.Width; j++ {
			mt := this.MapData.GetTileAt(i, j)
			if mt.OwnerID == playerID {
				tax += define.TAX_PER_TILE
			}
		}
	}
	return tax
}

func NewController() *Controller {
	availableCountries := []int{define.C_QIN, define.C_ZHENG}
	mapData := NewMap()
	destinyPool := NewDestinyPool()
	players := make([]*Player, PlayerNum)
	for i := 0; i < PlayerNum; i++ {
		players[i] = &Player{
			ID:      i,
			Country: define.C_EMPTY,
			Cards:   nil,
			Money:   0,
		}
	}
	round := 0
	c := &Controller{
		Round:              round,
		MapData:            mapData,
		Players:            players,
		DestiniesInPool:    destinyPool,
		AvailableCountries: availableCountries,
		Countries:          availableCountries,
	}
	return c
}

func test() {
	a := Controller{}
	go a.Run()
}
