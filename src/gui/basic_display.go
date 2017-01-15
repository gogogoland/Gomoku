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
	for i := 0; i < len(GameBoard) - 1 ; i++{
		fmt.Println(GameBoard[i])
	}
}
