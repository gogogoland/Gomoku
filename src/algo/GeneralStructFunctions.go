/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   GeneralStructFunctions.go                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/30 16:31:14 by tbalea            #+#    #+#             */
/*   Updated: 2016/10/30 16:31:14 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

/**
 * Functions for Structures and other
 *
 * TODO:
 * 		copy of GameData to DO ?
 * 		NOTHING
**/

/*
 * Functions for AlignP structures
 */
func AlignPInit(x, y, dir int) AlignP {
	return AlignP{
		pos: PawnsInit(x, y),
		dir: dir,
	}
}

func AlignPCompare(ap1, ap2 AlignP) bool {
	return (PawnsCompare(ap1.pos, ap2.pos) && ap1.dir == ap2.dir)
}

func AlignPAdd(lst []AlignP, new AlignP) {
	var cur, lenLst int

	lenLst = len(lst)
	for cur = 0; cur < lenLst; cur++ {
		if AlignPCompare(lst[cur], new) {
			return
		}
	}
	lst = append(lst, new)
}

/*
 * Functions for Pawns structures
 */
func PawnsInit(x, y int) Pawns {
	return Pawns{
		x: x,
		y: y,
	}
}

func PawnsCompare(p1, p2 Pawns) bool {
	return (p1.x == p2.x && p1.y == p2.y)
}

/*
 * Functions for NextPawns structures
 */

func NextPawnsInit(x, y, test_n int, winpot float32) NextPawns {
	return NextPawns{
		pawn_p: PawnsInit(x, y),
		winpot: winpot,
		test_n: test_n,
	}
}

/*
 * Functions Board
 */

func BoardIntInit(height, width, value int) [][]int {
	var board [][]int

	board = make([][]int, height)
	for y := 0; y < height; y++ {
		board[y] = make([]int, width)
		for x := 0; x < width; x++ {
			board[y][x] = value
		}
	}
	/*
	 * Beg Special rules
	 */
	board[height/2+height%2][width/2+width%2] = 0
	/*
	 * End Special rules
	 */
	return board
}

/*
 * Functions for Player structures
 */

func PlayerInit(whoareyou int) Player {
	return Player{
		atenum: 0,
		whoiam: whoareyou,
		pawn_p: PawnsInit(-1, -1),
		five_w: nil,
		threef: nil,
		haswin: false,
		winpot: 0.0,
	}
}

/*
 * Functions for GameData structures
 */

func GameDataInit(whobegin int) GameData {
	return GameData{
		facundo: PlayerInit(1),
		human:   PlayerInit(2),
		board:   BoardIntInit(19, 19, -4),
		deep:    0,
		move:    PawnsInit(-1, -1),
		prob:    0,
		turn:    whobegin,
	}
}

func GameDataCopy(tocopy, theone GameData) {
	theone = tocopy
}

/*
 * Other functions
 */
func GetOtherTurn(gd GameData) int {
	var newTurn int

	if gd.turn == gd.human.whoiam {
		newTurn = gd.facundo.whoiam
	} else if gd.turn == gd.facundo.whoiam {
		newTurn = gd.human.whoiam
	}
	return newTurn
}

func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func BoolToFloat64(b bool) float64 {
	if b {
		return 1
	} else {
		return 0
	}
}

func BoolToFloat32(b bool) float32 {
	if b {
		return 1
	} else {
		return 0
	}
}
