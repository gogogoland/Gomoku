package gui

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
	"algo"
)

/*
 * TODO 
 * 	-> Do some check for the input.
*/

// Main Function
func BasicDisplay(whobegin int) {
	Play(whobegin)
}

//Play until GameDataGain != 0
func Play(whobegin int) {
	GData := algo.GameDataInit(whobegin)
	PrintBoard(GData.Board)
	for algo.GameDataGain(GData) == 0 {
		GData = algo.TurnProcess(GData, Input())
		PrintBoard(GData.Board)
		GData, _, _ = algo.Pathfinding(GData, 1, 1)
		PrintBoard(GData.Board)
	}
}

func PrintBoard(GameBoard [][]int) {
	for i := 0; i < len(GameBoard); i++{
		fmt.Println(GameBoard[i])
	}
}

//Wait for player input X and Y and return a Pawns Struct from algo pkg
func Input() algo.Pawns {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter X Y value: ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text,"\n", "", -1)
	getValue := strings.Split(text, " ")
	fmt.Println(getValue)

	X, _:= strconv.Atoi(getValue[0])
	Y, _:= strconv.Atoi(getValue[1])
	return algo.PawnsInit(X, Y)
}
