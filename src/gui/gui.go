package gui

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)

func appMain(driver gxui.Driver) {
	
	//Path to your image
	file := args[0]

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

	mx := source.Bounds().Max
	window := theme.CreateWindow(mx.X, mx.Y, "Image viewer")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(img)

	// Copy the image to a RGBA format
	// before handing to a gxui.Texture
	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)

	window.OnClose(driver.Terminate)
}

func Test() {
	flag.Parse()
	gl.StartDriver(appMain)
}
