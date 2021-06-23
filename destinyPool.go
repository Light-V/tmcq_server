package main

import (
	"container/list"
	"main/util"
	"math/rand"
)

type DestinyPool struct {
	// LIST of *Destiny
	AvailableDestinies *list.List
	// LIST of *Destiny
	UsedDestinies *list.List
}

func NewDestinyPool() *DestinyPool {
	available := list.New()
	used := list.New()
	//todo: 初始化天命
	available.PushBack(&Destiny{})
	return &DestinyPool{available, used}
}

func (this *DestinyPool) GetNextDestiny() *Destiny {
	length := this.AvailableDestinies.Len()
	n := rand.Intn(length)
	destinyElement := util.GetListElementAt(this.AvailableDestinies, n)
	this.AvailableDestinies.Remove(destinyElement)
	this.UsedDestinies.PushBack(destinyElement.Value)
	r, _ := destinyElement.Value.(*Destiny)
	return r
}
