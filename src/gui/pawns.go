package gui

import (
  "image"
  "image/draw"
)

/*
** TYPE : Private
** Function who print pawns.
*/

func drawPawns(color int, posX int, posY int, wp_img image.Image, bp_img image.Image, rgba *image.RGBA, size_board image.Point) {

  pawnsX := 6 + 53 * posX + posX % 2
  pawnsY := 6 + 53 * posY + posY % 2

  if (color == 1) {
    draw.Draw(rgba, image.Rect(pawnsX, pawnsY, size_board.X, size_board.Y), wp_img, image.ZP, draw.Over)
  } else if (color == 2) {
    draw.Draw(rgba, image.Rect(pawnsX, pawnsY, size_board.X, size_board.Y), bp_img, image.ZP, draw.Over)
  }
}
