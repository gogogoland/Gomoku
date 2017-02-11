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

//Remove void data in slice
func DeleteElemFromSliceOfString(str []string, del string) []string {
	var ret []string
	var cur string

	for _, cur = range str {
		if cur != del {
			ret = append(ret, cur)
		}
	}
	return ret
}

//Play until GameDataGain != 0
func Play(whobegin int) {
	var retry int

	GData := algo.GameDataInit(whobegin)
	GData.PrintBoard()
	for GData.Gain() == 0 {
		retry = 1
		for retry > 0 {
			retry = GData.TurnProcess(Input())
		}
		GData.PrintBoard()
		fmt.Println(*(GData.GetHuman()))
		GData.Pathfinding(3, 1)
		GData.PrintBoard()
	}
}

//Wait for player input X and Y and return a Pawns Struct from algo pkg
func Input() algo.Pawns {
	var getValue []string
	var i, nbrValue int

	reader := bufio.NewReader(os.Stdin)
	for getValue = nil; getValue == nil || len(getValue) != 2; {
		fmt.Print("Enter X Y value: ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		getValue = strings.Split(text, " ")
		getValue = DeleteElemFromSliceOfString(getValue, "")
		for nbrValue, i = len(getValue), 0; i < nbrValue; i++ {
			print("\"", getValue[i], "\" ")
		}
		print("\n")
	}

	X, _ := strconv.Atoi(getValue[0])
	Y, _ := strconv.Atoi(getValue[1])
	return algo.PawnsInit(X, Y)
}
