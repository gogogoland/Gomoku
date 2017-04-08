package gui

import (
	"algo"
	"fmt"
)

//Play until GameDataGain != 0
func play(posY, posX int, vs bool) int {

	if GData.PlayerTurn(algo.PawnsInit(posX, posY)) == 1 {
		GData.PrintBoard()
		fmt.Println("Forbidden Move.")
		fmt.Println("HUMAN RESULT:", *(GData.GetHuman()))
		fmt.Println("FACUNDO RESULT:", *(GData.GetFacundo()))
		return 1
	}

	GData.PrintBoard()
	fmt.Println("HUMAN RESULT:", *(GData.GetHuman()))
	fmt.Println("FACUNDO RESULT:", *(GData.GetFacundo()))
	if GData.Gain() != 0 {
		return 0
	}
	GData.Pathfinding(3, 1, vs)
	GData.PrintBoard()
	fmt.Println("HUMAN RESULT:", *(GData.GetHuman()))
	fmt.Println("FACUNDO RESULT:", *(GData.GetFacundo()))
	return 0
}

/*
func updateBoard(wp_img image.Image, bp_img image.Image, rgba *image.RGBA) {
}
*/
