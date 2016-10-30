/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Structures.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/03/22 18:42:35 by tbalea            #+#    #+#             */
/*   Updated: 2016/03/22 18:42:35 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

/**
 * 	TODO:
 * 		NOTHING
**/

import (
	"container/heap"
	"container/list"
	"fmt"
	"github.com/samuel/go-opencl/cl"
	"math"
)

//	Structures
//	*	Position of Pawns
type Pawns struct {
	x, y int
}

//	*	Win Alignement
type AlingnW struct {
	pos Pawns
	dir int
}

//	*	Player data
type Player struct {
	atenum int
	whoiam int
	pawn_p Pawns
	five_w []AlignW
	threef bool
	haswin bool
	winpot float32
}

//	*	MinMax data
type GameData struct {
	facundo, human Player
	//salver is a better name, no ?
	board [][]int
	deep  int
	move  Pawns
	prob  int
	turn  int
}

//	*	Potential Move
type NextPawns struct {
	pawn_p Pawns
	winpot float32
	test_n int
}

//	*	Slice of MinMax
type PrioQueue []MinMax
