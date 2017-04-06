package gui

import(
  "algo"
  "fmt"
)

//Play until GameDataGain != 0
func play(posY, posX int) int{
    GData.PrintBoard()

    fmt.Println("HUMAN RESULT:")
    if (GData.PlayerTurn(algo.PawnsInit(posX, posY)) == 1) {
      return 1
    }

    GData.PrintBoard()
    fmt.Println("HUMAN RESULT:", *(GData.GetHuman()))
    GData.Pathfinding(3, 1)
    GData.PrintBoard()
    return 0
}
/*
func updateBoard(wp_img image.Image, bp_img image.Image, rgba *image.RGBA) {
}
*/
