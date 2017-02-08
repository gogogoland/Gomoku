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
)

func drawStar(canvas gxui.Canvas, center math.Point, radius float32, points int, color gxui.Color) {
	p := make(gxui.Polygon, points*2)
	for i := 0; i < points*2; i++ {

		frac := float32(i) / float32(points*2)
		α := frac * math.TwoPi
		r := []float32{radius, radius}[i&1]

		p[i] = gxui.PolygonVertex{
			Position: math.Point{
				X: center.X + int(r*math.Cosf(α)),
				Y: center.Y + int(r*math.Sinf(α)),
			},
			RoundedRadius: []float32{0, 50}[i&1],
		}
	}
	canvas.DrawPolygon(p, gxui.CreatePen(1.25, gxui.Black), gxui.CreateBrush(color))
}

func DrawPawn(driver gxui.Driver, window gxui.Window, color gxui.Color, center math.Point) {
	theme := flags.CreateTheme(driver)

	canvas := driver.CreateCanvas(math.Size{W: width, H: height})
	drawStar(canvas, center, 25, 38, gxui.White)

	canvas.Complete()

	image := theme.CreateImage()
	image.SetCanvas(canvas)
	window.AddChild(image)
}

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
			DrawPawn(driver, window, gxui.White, me.WindowPoint)
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
