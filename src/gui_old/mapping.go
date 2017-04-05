package gui

import (
	"container/list"
	"github.com/google/gxui/math"
)

type Map struct {
	PosX, PosY, TabPosX, TabPosY int
}

func InitMap() *list.List{
	//Init my list
	list := list.New()

	Px, Py := (width - 44) / 19, (height - 200) / 19

	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			list.PushBack(Map{(Px * x) + 52, (Py * y) + 45, x, y})
		}
	}
	return list
}

func getClickPosInTab(list *list.List, pos math.Point) Map {
	for e := list.Front(); e != nil; e = e.Next() {
		if (pos.X >= e.Value.(Map).PosX - 15 && pos.X <= e.Value.(Map).PosX + 15 && pos.Y >= e.Value.(Map).PosY - 15 && pos.Y <= e.Value.(Map).PosY + 15){
			return e.Value.(Map)
		}
	}
	return Map{-1, -1, -1, -1}
}
