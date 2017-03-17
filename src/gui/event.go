package gui

import(
	"github.com/google/gxui/math"
	"github.com/google/gxui"
	"fmt"
	"os"
)


/*
** Handle all event with Gxui for gomoku
*/

func Event(driver gxui.Driver, window gxui.Window) {
	window.OnKeyDown(func(ev gxui.KeyboardEvent) {
		if ev.Key == gxui.KeyEscape || ev.Key == gxui.KeyKpEnter {
			fmt.Println("Close")
			window.Close()
		}
	})

	//Init List
	list := InitMap()

	window.OnClick(func(me gxui.MouseEvent) {
		if me.Button == 0 {
			fmt.Println(me.WindowPoint)

			value := getClickPosInTab(list, me.WindowPoint)
			if (value.PosX != -1 && value.PosY != -1){
				DrawPawn(driver, window, gxui.White, math.Point{value.PosX,value.PosY})
			}
		}
	})

	window.OnResize(func(){
		fmt.Println("[ERROR] : someone try to resize the window.")
		os.Exit(1)
	})
}
