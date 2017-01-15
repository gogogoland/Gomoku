package gui

import (
	"fmt"
	"algo"
)

func BasicDisplay(whobegin int) {
	Init := algo.GameDataInit(whobegin)
	//Example := [][]int  {{1,2,3,4},
	//						 {1,2,3,4},
	//					 {1,2,3,4}}
	//PrintBoard(Example)
	test := Init.Board
	fmt.Println(test)
}

func PrintBoard(GameBoard [][]int) {

}
