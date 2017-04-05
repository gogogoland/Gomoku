package gui

import(
  "github.com/google/gxui/math"
)

func getCursorsPosition(mouse math.Point) math.Point{
    posX, posY := float64(mouse.X + 7) / 53.36, float64(mouse.Y + 7) / 53.36
    diffX, diffY := posX - float64(int(posX)), posY - float64(int(posY))
    cursX, cursY := -1, -1


    if (diffX >= 0.6) {
      cursX = int(posX)
    } else if (diffX <= 0.4 && diffX >= 0) {
      cursX = int(posX) - 1
    }

    if (diffY >= 0.6) {
      cursY = int(posY)
    } else if (diffY <= 0.4 && diffY >= 0) {
      cursY = int(posY) - 1
    }

    return math.Point{ X: cursX, Y: cursY }
}
