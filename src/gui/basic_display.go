package gui

import (
	"algo"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var retry int

	GData := algo.GameDataInit(whobegin)
	PrintBoard(GData.GetBoard())
	for GData.Gain() == 0 {
		retry = 1
		for retry > 0 {
			retry = GData.TurnProcess(Input())
		}
		GData.Pathfinding(3, 1)
		PrintBoard(GData.GetBoard())
	}
}

func PrintBoard(GameBoard [][]int) {
	for y, ymax := 0, len(GameBoard); y < ymax; y++ {
		for x, xmax := 0, len(GameBoard[y]); x < xmax; x++ {
			if GameBoard[y][x] >= 0 {
				print("  ", GameBoard[y][x])
			} else {
				print(" ", GameBoard[y][x])
			}
		}
		print("\n")
	}
}

//Wait for player input X and Y and return a Pawns Struct from algo pkg
func Input() algo.Pawns {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter X Y value: ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	getValue := strings.Split(text, " ")
	fmt.Println(getValue)

	X, _ := strconv.Atoi(getValue[0])
	Y, _ := strconv.Atoi(getValue[1])
	return algo.PawnsInit(X, Y)
}
