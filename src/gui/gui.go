package gui

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)


/*
** Little test for key usage with gxui
*/

func appMain(driver gxui.Driver){
	theme := flags.CreateTheme(driver)

	window := theme.CreateWindow(800, 600, "Open file...")
	window.OnKeyDown(func(ev gxui.KeyboardEvent) {
		if ev.Key == gxui.KeyEscape || ev.Key == gxui.KeyKpEnter {
			fmt.Println("Close")
			window.Close()
		}
	})

	window.OnClick(func(me gxui.MouseEvent) {
		if me.Button == 0 {
			fmt.Println("Right Click.")
		}
	})

	window.OnClose(driver.Terminate)
}

func Test() {
	gl.StartDriver(appMain)
}
