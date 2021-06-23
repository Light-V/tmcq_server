package main

import "main/define"

type Player struct {
	//玩家ID
	ID int
	//势力
	Country int
	//都城
	Capital *MapTile
	//君主
	//手牌
	Cards []*Card
	//金钱
	Money int
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
