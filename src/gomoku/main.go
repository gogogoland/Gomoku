//tempory example for building
package main

import (
	"os"
	"tempory_show"
)

/*
 * 	TODO:
 * 		AI:
 * 			1)Choose who begin (1 = AI/Facundo; 2 = Human/Player)
 * 			2) Init GameData structure by calling GameDataInit(beginner value)
 *			3) On Human turn wait for correct entry and launch TurnProcess
 *			4) On AI turn launch Pathfinding(GameData, Deep int, AI_type int)
 * 			5) Result give you new GameData + AI entry + Current Board
 *			6) Check Victory by calling GameDataGain(GameData), it return :
 * 				0 for no winner, 1 for AI/Facundo win, 2 for Player/Human win
 * 			7) Repeat at 3)
 *
 * 		GAMEDATA STRUCT:
 * 			See GameData structure in src/algo/Structures.go for game's data
 */

func run() int {
	return 0
}

func main() {
	tempory_show.Example()
	os.Exit(run())
}
