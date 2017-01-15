package gui

import (
	"fmt"
	"algo"
)

func BasicDisplay(whobegin int) {
	Init := algo.GameDataInit(whobegin)
	PrintBoard(Init.Board)
}

func PrintBoard(GameBoard [][]int) {
	for i := 0; i < len(GameBoard); i++{
		fmt.Println(GameBoard[i])
	}
}
