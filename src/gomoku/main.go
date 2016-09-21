//tempory example for building
package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

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
