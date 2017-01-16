package gui

import (
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
	window.OnClose(driver.Terminate)
}

/*
** I'm not using it for the moment.
*/
/*
func appMain2(driver gxui.Driver) {
	//Path to your image
	file := "nop"

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
*/

func Test() {
	gl.StartDriver(appMain)
}
