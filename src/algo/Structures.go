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
	//four_w SliceAP
	threef SliceAP
	//tofree SliceAP
	//four_p SliceAP
	//threep SliceAP
	winpot float32
}

//	*	Board data
type Board [][]int

//	*	MinMax data
type GameData struct {
	facundo, human Player
	//salver is a better name, no ?
	board  Board
	maxx   int
	maxy   int
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
