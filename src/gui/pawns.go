package gui

import (
  "image"
  "image/draw"
)

func drawPawns(color int, posX int, posY int, wp_img image.Image, bp_img image.Image, rgba *image.RGBA, size_board image.Point) {
  if (color == 0){
    draw.Draw(rgba, image.Rect(posX, posY, size_board.X, size_board.Y), wp_img, image.ZP, draw.Over)
  } else {
    draw.Draw(rgba, image.Rect(posX, posY, size_board.X, size_board.Y), bp_img, image.ZP, draw.Over)
  }
}
