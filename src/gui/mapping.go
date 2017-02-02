package gui

import (
	"fmt"
	"container/list"
)

type Map struct {
	PosX, PosY, PposX, PposY int
}

func MappingBoard() *List{
	//Init my list
	list := list.New()

	Px, Py := height / 19, width / 19

	for x := 0; x < 19; x++{
		for y := 0; y < 19; y++{
			list.PushBack(Map{Px + 1 * x, Py + 1 * y, Px * x, Py * y})
		}
	}
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	return list
}
