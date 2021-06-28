package util

import (
	"container/list"
)

func GetListElementAt(l *list.List, index int) *list.Element {
	cnt := 0
	var i *list.Element
	for i = l.Front(); i != nil && cnt < index; i = i.Next() {
		cnt++
	}
	return i
}
