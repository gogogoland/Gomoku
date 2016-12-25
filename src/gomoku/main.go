//tempory example for building
package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
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

var winTitle string = "Go-SDL2 Events"
var winWidth, winHeight int = 800, 600

func run() int {
	var window *sdl.Window
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	return 0
}

func main() {
	os.Exit(run())
}
