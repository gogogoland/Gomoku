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

func (data AlignP) Compare(cmp AlignP) bool {
	return (data.pos.Compare(cmp.GetPos()) && data.dir == cmp.GetDir())
}

func (data AlignP) Copy() AlignP {
	return AlignP{
		pos: data.pos.Copy(),
		dir: data.dir,
	}
}

/*
 * Get functions for AlignP
 */

func (get *AlignP) GetPos() Pawns {
	return get.pos
}

func (get *AlignP) GetDir() int {
	return get.dir
}

/*
 * For Slice of AlignP
 */

func (data *SliceAP) Add(new AlignP) {
	var cur, lenSlice int

	lenSlice = len(*data)
	for cur = 0; cur < lenSlice; cur++ {
		if (*data)[cur].Compare(new) {
			return
		}
	}
	*data = append(*data, new)
}

func (data SliceAP) Copy() SliceAP {
	var theone SliceAP
	var lenAlP, curAlP int

	lenAlP = len(data)
	theone = make(SliceAP, lenAlP)
	for curAlP = 0; curAlP < lenAlP; curAlP++ {
		theone[curAlP] = data[curAlP].Copy()
	}
	return theone
}

/*
 * End Functions for AlignP
 */

/*
 * Functions for Pawns structures
 */
func PawnsInit(x, y int) Pawns {
	return Pawns{
		x: x,
		y: y,
	}
}

func (data Pawns) Compare(cmp Pawns) bool {
	return (data.x == cmp.GetX() && data.y == cmp.GetY())
}

func (data Pawns) Copy() Pawns {
	return PawnsInit(data.x, data.y)
}

/*
 * Get functions for Pawns
 */

func (get *Pawns) GetX() int {
	return get.x
}

func (get *Pawns) GetY() int {
	return get.y
}

/*
 * End Functions for Pawns
 */

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

func (data NextPawns) Copy() NextPawns {
	return NextPawns{
		pawn_p: data.pawn_p.Copy(),
		winpot: data.winpot,
		test_n: data.test_n,
	}
}

/*
 * Get functions for NextPawns
 */

func (get *NextPawns) GetPawn_P() Pawns {
	return get.pawn_p
}

func (get *NextPawns) GetWinPot() float32 {
	return get.winpot
}

func (get *NextPawns) GetTest_N() int {
	return get.test_n
}

/*
 * End Functions for NextPawns
 */

/*
 * Functions Board
 */

func BoardIntInit(height, width, value int) Board {
	var board Board

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

func (data Board) Copy() Board {
	var theone Board
	var ylen, xlen int
	var y, x int

	ylen = len(data)
	theone = make([][]int, ylen)
	for y = 0; y < ylen; y++ {
		xlen = len(data[y])
		theone[y] = make([]int, xlen)
		for x = 0; x < xlen; x++ {
			theone[y][x] = data[y][x]
		}
	}
	return theone
}

/*
 * End Functions for Board
 */

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

func (data *Player) Copy() Player {
	return Player{
		atenum: data.atenum,
		whoiam: data.whoiam,
		pawn_p: data.pawn_p.Copy(),
		five_w: data.five_w.Copy(),
		threef: data.threef.Copy(),
		winpot: data.winpot,
	}
}

/*
 * Get functions for Player
 */

func (get *Player) GetAteNum() int {
	return get.atenum
}

func (get *Player) GetWhoIAm() int {
	return get.whoiam
}

func (get *Player) GetPawn_P() Pawns {
	return get.pawn_p
}

func (get *Player) GetFive_W() SliceAP {
	return get.five_w
}

func (get *Player) GetFour_W() SliceAP {
	return get.four_w
}

func (get *Player) GetThreeF() SliceAP {
	return get.threef
}

func (get *Player) GetToFree() SliceAP {
	return get.tofree
}

func (get *Player) GetFour_P() SliceAP {
	return get.four_p
}

func (get *Player) GetThreeP() SliceAP {
	return get.threep
}

func (get *Player) GetWinPot() float32 {
	return get.winpot
}

/*
 * End Functions for Player
 */

/*
 * Functions for GameData structures
 */

func GameDataInit(whobegin int) GameData {
	return GameData{
		facundo: PlayerInit(2),
		human:   PlayerInit(1),
		board:   BoardIntInit(19, 19, -4),
		maxx:    19,
		maxy:    19,
		deep:    0,
		move:    PawnsInit(-1, -1),
		prob:    0,
		turn:    whobegin,
		whowin:  0,
	}
}

func (data *GameData) Copy() GameData {
	return GameData{
		facundo: data.facundo.Copy(),
		human:   data.human.Copy(),
		board:   data.board.Copy(),
		maxx:    data.maxx,
		maxy:    data.maxy,
		deep:    data.deep,
		move:    data.move.Copy(),
		prob:    data.prob,
		turn:    data.turn,
		whowin:  data.whowin,
	}
}

func (data *GameData) Deep(deepness int) GameData {
	var theone GameData

	theone = data.Copy()
	theone.deep = deepness
	return theone
}

func (data *GameData) Gain() int {
	data.whowin = data.facundo.whoiam * BoolToInt(data.facundo.winpot >= 1.0)
	data.whowin += data.human.whoiam * BoolToInt(data.human.winpot >= 1.0)
	return data.whowin
}

/*
 * Get function for GameData value
 */

func (get *GameData) GetHuman() *Player {
	return &get.human
}

func (get *GameData) GetFacundo() *Player {
	return &get.facundo
}

func (get *GameData) GetPlayer(whoiam int) *Player {
	if get.human.whoiam == whoiam {
		return get.GetHuman()
	} else if get.facundo.whoiam == whoiam {
		return get.GetFacundo()
	}
	return nil
}

func (get *GameData) GetBoard() Board {
	return get.board
}

func (get *GameData) GetMaxX() int {
	return get.maxx
}

func (get *GameData) GetMaxY() int {
	return get.maxy
}

func (get *GameData) GetWhoWin() int {
	return get.whowin
}

func (get *GameData) GetMove() Pawns {
	return get.move
}

func (get *GameData) GetTurn() int {
	return get.turn
}

func (get *GameData) GetOtherTurn() int {
	if get.turn == get.human.whoiam {
		return get.facundo.whoiam
	} else if get.turn == get.facundo.whoiam {
		return get.human.whoiam
	}
	return 0
}

func (get *GameData) GetProb() int {
	return get.prob
}

func (get *GameData) PrintBoard() {
	var x, y int

	for y = 0; y < get.maxy; y++ {
		for x = 0; x < get.maxx; x++ {
			if get.board[y][x] >= 0 {
				print("  ", get.board[y][x])
			} else if get.board[y][x] != -4 {
				print(" ", get.board[y][x])
			} else if get.board[y][x] == -4 {
				print("  .")
			}
		}
		print("\n")
	}
}

/*
 * End Functions for GameData
 */

/*
 * Other functions
 */
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
