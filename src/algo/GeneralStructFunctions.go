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

func AlignPAdd(lst *([]AlignP), new AlignP) {
	var cur, lenLst int

	lenLst = len(*lst)
	for cur = 0; cur < lenLst; cur++ {
		if AlignPCompare((*lst)[cur], new) {
			return
		}
	}
	*lst = append(*lst, new)
}

func AlignPCopy(tocopy AlignP) AlignP {
	return AlignP{
		pos: PawnsCopy(tocopy.pos),
		dir: tocopy.dir,
	}
}

func AlignPSliceCopy(tocopy []AlignP) []AlignP {
	var theone []AlignP
	var lenAlP, curAlP int

	lenAlP = len(tocopy)
	theone = make([]AlignP, lenAlP)
	for curAlP = 0; curAlP < lenAlP; curAlP++ {
		theone[curAlP] = AlignPCopy(tocopy[curAlP])
	}
	return theone
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

func PawnsCopy(tocopy Pawns) Pawns {
	return PawnsInit(tocopy.x, tocopy.y)
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

func NextPawnsCopy(tocopy NextPawns) NextPawns {
	return NextPawns{
		pawn_p: PawnsCopy(tocopy.pawn_p),
		winpot: tocopy.winpot,
		test_n: tocopy.test_n,
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

func BoardIntCopy(tocopy [][]int) [][]int {
	var theone [][]int
	var ylen, xlen int
	var y, x int

	ylen = len(tocopy)
	theone = make([][]int, ylen)
	for y = 0; y < ylen; y++ {
		xlen = len(tocopy[y])
		theone[y] = make([]int, xlen)
		for x = 0; x < xlen; x++ {
			theone[y][x] = tocopy[y][x]
		}
	}
	return theone
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
		winpot: 0.0,
	}
}

func PlayerCopy(tocopy Player) Player {
	return Player{
		atenum: tocopy.atenum,
		whoiam: tocopy.whoiam,
		pawn_p: PawnsCopy(tocopy.pawn_p),
		five_w: AlignPSliceCopy(tocopy.five_w),
		threef: AlignPSliceCopy(tocopy.threef),
		winpot: tocopy.winpot,
	}
}

/*
 * Functions for GameData structures
 */

func GameDataInit(whobegin int) GameData {
	return GameData{
		facundo: PlayerInit(2),
		human:   PlayerInit(1),
		Board:   BoardIntInit(19, 19, -4),
		deep:    0,
		move:    PawnsInit(-1, -1),
		prob:    0,
		turn:    whobegin,
		whowin:  0,
	}
}

func GameDataCopy(tocopy /*, theone */ GameData) GameData {
	//theone = tocopy
	return GameData{
		facundo: PlayerCopy(tocopy.facundo),
		human:   PlayerCopy(tocopy.human),
		Board:   BoardIntCopy(tocopy.Board),
		deep:    tocopy.deep,
		move:    PawnsCopy(tocopy.move),
		prob:    tocopy.prob,
		turn:    tocopy.turn,
		whowin:  tocopy.whowin,
	}
}

func GameDataDeep(tocopy GameData, deepness int) GameData {
	var theone GameData

	theone = GameDataCopy(tocopy)
	theone.deep = deepness
	return theone
}

func GameDataGain(gain GameData) int {
	gain.whowin = gain.facundo.whoiam * BoolToInt(gain.facundo.winpot >= 1.0)
	gain.whowin += gain.human.whoiam * BoolToInt(gain.human.winpot >= 1.0)
	return gain.whowin
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
