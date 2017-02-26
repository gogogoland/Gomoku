package gui

import (
	"container/list"
	"github.com/google/gxui/math"
	"fmt"
)

type Map struct {
	PosX, PosY, TabPosX, TabPosY int
}

func Debug() *list.List{
	//Init my list
	list := list.New()

	Px, Py := (width - 44) / 19, (height - 200) / 19

	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			list.PushBack(Map{(Px * x) + 52, (Py * y) + 38, x, y})
		}
	}
	//for e := list.Front(); e != nil; e = e.Next() {
	//	DrawPawn(driver, window, gxui.Black, math.Point{e.Value.(Map).PosX, e.Value.(Map).PosY})
	//}
	return list
}

func getClickPosInTab(list *list.List, pos math.Point) {
	for e := list.Front(); e != nil; e = e.Next() {
		if (pos.X >= e.Value.(Map).PosX - 15 && pos.X <= e.Value.(Map).PosX + 15 && pos.Y >= e.Value.(Map).PosY - 15 && pos.Y <= e.Value.(Map).PosY + 15){
			fmt.Println(e.Value)
		}
	}
}
