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
 * All import and used structure for MinMax
 *
 * 	TODO:
 * 		NOTHING
**/

//import (
//	"container/heap"
//	"container/list"
//	"fmt"
//	/*"github.com/samuel/gop-opencl/cl"*/
//	"math"
//)

//	Structures
//	*	Position of Pawns
type Pawns struct {
	x, y int
}

//	*	Alignement of Pawns
type AlignP struct {
	pos Pawns
	dir int
}

//	*	Slice of Alignement
type SliceAP []AlignP

//	*	Player data
type Player struct {
	atenum int
	whoiam int
	pawn_p Pawns
	five_w SliceAP
	threef SliceAP
	winpot float32
}

//	*	Board data
type Board [][]int

//	*	MinMax data
type GameData struct {
	facundo, human Player
	//salver is a better name, no ?
	board  Board
	deep   int
	move   Pawns
	prob   int
	turn   int
	whowin int
	//Add probabilty ?
	//wilo	float32
}

//	*	Potential Move
type NextPawns struct {
	pawn_p Pawns
	winpot float32
	test_n int
}

//	*	Slice of MinMax data
type PrioQueue []GameData
