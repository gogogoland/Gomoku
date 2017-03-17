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
	"github.com/google/gxui/samples/flags"
)


var (
	title         string = "Gomoku"
	width, height int    = 1024, 1024
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

	window := theme.CreateWindow(height, width, title)
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(img)

	Event(driver, window)

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
