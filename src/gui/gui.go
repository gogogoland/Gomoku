package gui

/*
** MAIN of the package.
*/

import (
	"image"
	"image/draw"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)

func appMain(driver gxui.Driver) {

	board_img, wp_img, bp_img := getSprites()

	theme := flags.CreateTheme(driver)
	img := theme.CreateImage()

	size_board := board_img.Bounds().Max
	window := theme.CreateWindow(size_board.X, size_board.Y, title)

	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(img)

	handleEvent(driver, window)

	rgba := image.NewRGBA(board_img.Bounds())

	draw.Draw(rgba, board_img.Bounds(), board_img, image.ZP, draw.Src)

	for i, j := 6, 0; i < 1000; i, j = i + 53 + j % 2, j + 1 {
		for k, l := 6, 0; k < 1000; k, l = k + 53 + l % 2, l + 1 {
			drawPawns(1, i, k, wp_img, bp_img, rgba, size_board)
		}
	}
	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)

	window.OnClose(driver.Terminate)
}

func Gui() {
	gl.StartDriver(appMain)
}
