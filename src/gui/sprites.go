package gui

import (
  "os"
  "fmt"
  "image"
)

/*
** TYPE : PRIVATE FUNCTION
**
** OPEN ALL SPRITES, DECODE and RETURN io.stream
*/

func getSprites() (image.Image, image.Image, image.Image){
  board_fd, err := os.Open(board)
  if err != nil {
    fmt.Printf("[ERROR] Failed to open image => %v\n", err)
    os.Exit(-1)
  }

  board_dc, _, err := image.Decode(board_fd)
  if err != nil {
		fmt.Printf("[ERROR] Failed to read image => %v\n", err)
		os.Exit(-1)
	}

  white_pawns_fd, err := os.Open(white_pawns)
  if err != nil {
    fmt.Printf("[ERROR] Failed to open image => %v\n", err)
    os.Exit(-1)
  }

  white_pawns_dc, _, err := image.Decode(white_pawns_fd)
  if err != nil {
		fmt.Printf("[ERROR] Failed to read image => %v\n", err)
		os.Exit(-1)
	}

  black_pawns_fd, err := os.Open(black_pawns)
  if err != nil {
    fmt.Printf("[ERROR] Failed to open image => %v\n", err)
    os.Exit(-1)
  }

  black_pawns_dc, _, err := image.Decode(black_pawns_fd)
  if err != nil {
		fmt.Printf("[ERROR] Failed to read image => %v\n", err)
		os.Exit(-1)
	}
  return board_dc, white_pawns_dc, black_pawns_dc
}
