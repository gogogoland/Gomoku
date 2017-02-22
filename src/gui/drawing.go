package gui

import(
	//"fmt"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/gxui"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

func drawP(canvas gxui.Canvas, center math.Point, radius float32, points int, color gxui.Color) {
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
	drawP(canvas, center, 23, 38, gxui.White)

	canvas.Complete()

	image := theme.CreateImage()
	image.SetCanvas(canvas)
	window.AddChild(image)
}
