package gui

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
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
	X, Y := Input()
	fmt.Println(X, Y)
}

/*
 * Get X, Y and parse it.
 * TODO 
 * 	-> Do some check for the input.
 *	-> Replace Value in board by 2 for player turn
 *	-> Use AI
*/

func Input() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter X Y value: ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text,"\n", "", -1)
	getValue := strings.Split(text, " ")
	fmt.Println(getValue)

	X, _:= strconv.Atoi(getValue[0])
	Y, _:= strconv.Atoi(getValue[1])
	return X, Y
}
