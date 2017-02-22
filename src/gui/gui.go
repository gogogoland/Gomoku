package gui

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"image"
	"image/draw"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

var (
	Title         string = "Gomoku"
	width, height int    = 1024, 1024
	InitX, InitY int = 52, 38
)

func appMain(driver gxui.Driver) {
	//Path to Image
	file := "sprites/GoBoard.png"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Failed to open image '%s': %v\n", file, err)
		os.Exit(1)
	}

	source, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Failed to read image '%s': %v\n", file, err)
		os.Exit(1)
	}
	theme := flags.CreateTheme(driver)
	img := theme.CreateImage()

	window := theme.CreateWindow(height, width, Title)
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(img)

	window.OnKeyDown(func(ev gxui.KeyboardEvent) {
		if ev.Key == gxui.KeyEscape || ev.Key == gxui.KeyKpEnter {
			fmt.Println("Close")
			window.Close()
		}
	})

	window.OnClick(func(me gxui.MouseEvent) {
		if me.Button == 0 {
			fmt.Println("Right Click.")
			fmt.Println(me.WindowPoint)
			DrawPawn(driver, window, gxui.White, math.Point{52, 38})
			DrawPawn(driver, window, gxui.Black, math.Point{52 + width/19 , 38})
		}
	})

	// Copy the image to a RGBA format before handing to a gxui.Texture
	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)

	window.OnClose(driver.Terminate)
}

func Test() {
	gl.StartDriver(appMain)
}
