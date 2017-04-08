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

	"fmt"
)

func appMain(driver gxui.Driver) {

	board_img, wp_img, bp_img := getSprites()

	theme := flags.CreateTheme(driver)
	img := theme.CreateImage()

	size_board := board_img.Bounds().Max
	window := theme.CreateWindow(size_board.X, size_board.Y, title)

	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(img)

	rgba := image.NewRGBA(board_img.Bounds())
	draw.Draw(rgba, board_img.Bounds(), board_img, image.ZP, draw.Src)

	handleEvent(driver, window)

	//Get position of the click and check if the move is ok.
	window.OnClick(func(me gxui.MouseEvent) {
		if me.Button == 0 {
			pos := getCursorsPosition(me.WindowPoint)
			if GData.Gain() == 1 {
				fmt.Println("YOU WIN.")
			} else if GData.Gain() == 2 {
				fmt.Println("YOU LOSE.")
			} else {
				if play(pos.X, pos.Y, false) == 0 {
					draw.Draw(rgba, board_img.Bounds(), board_img, image.ZP, draw.Src)
					board := GData.GetBoard()
					if GData.Gain() == 1 {
						fmt.Println("YOU WIN.")
					} else if GData.Gain() == 2 {
						fmt.Println("YOU LOSE.")
					}
					for i := 0; i < 19; i += 1 {
						for j := 0; j < 19; j += 1 {
							if board[i][j] > 0 {
								drawPawns(board[i][j], j, i, wp_img, bp_img, rgba, size_board)
								texture := driver.CreateTexture(rgba, 1)
								img.SetTexture(texture)
							}
						}
					}
				}
			}
		}
	})

	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)

	window.OnClose(driver.Terminate)
}

func Gui() {
	gl.StartDriver(appMain)
}
