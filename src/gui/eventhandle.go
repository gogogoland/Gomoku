package gui

import(
	"github.com/google/gxui"
	"fmt"
	"os"
)

/*
** TYPE : Private
** Handle all event with Gxui
*/

func handleEvent(driver gxui.Driver, window gxui.Window) {
	window.OnKeyDown(func(ev gxui.KeyboardEvent) {
		if ev.Key == gxui.KeyEscape || ev.Key == gxui.KeyKpEnter {
			fmt.Println("[LOG] Game exit, you pressed 'esc'.")
			window.Close()
		}
	})

	window.OnClick(func(me gxui.MouseEvent) {
		if me.Button == 0 {
			fmt.Println(me.WindowPoint)
		}
	})

	window.OnResize(func(){
		fmt.Println("[ERROR] : someone try to resize the window.")
		os.Exit(1)
	})
}
