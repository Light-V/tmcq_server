package main

import (
	"container/list"
	"main/util"
	"math/rand"
	"time"
)

type DestinyPool struct {
	// LIST of *Destiny
	AvailableDestinies *list.List
	// LIST of *Destiny
	UsedDestinies *list.List
}

func NewDestinyPool() *DestinyPool {
	rand.Seed(int64(time.Now().Nanosecond()))
	available := list.New()
	used := list.New()
	//todo: 初始化天命
	for i := 0; i < 10; i++ {
		available.PushBack(&Destiny{})
	}
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
