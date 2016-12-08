/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   MinMax.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/27 16:05:19 by tbalea            #+#    #+#             */
/*   Updated: 2016/10/27 16:05:19 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

/**
 * Adding pawn layer and get probabilty of win
 *
 * TODO:
 * 		Seek comments inside functions
 *
 * 		NOTHING
**/

//	Check if pawn are aligned (for winning move (if win and three free)
func CheckAlignement(child GameData, pawn Pawns) {
	var x, y, s int
	var xmax, ymax int
	var check, align, dispo int

	align = 0
	xmax = len(child.board)
	ymax = len(child.board[0])
	for s = 0; s < 4 && align < 5; s++ {
		dispo = 0
		align = 0
		for check = -4; check <= 4 && align < 5; check++ {
			x = pawn.x + check*(s/2+BoolToInt(s == 0))
			y = pawn.y + check*(s%2-BoolToInt(s == 0))
			if x < 0 || y < 0 || x >= xmax || y >= ymax {
				continue
			} else if child.board[x][y] == child.turn {
				dispo++
				align++
			} else {
				if child.board[x][y] <= 0 && align > 0 {
					dispo++
				} else {
					dispo = BoolToInt(child.board[x][y] <= 0)
				}
				align = 0
			}
			//	Maybe dispo == 6
			if align >= 3 && dispo == 5 {
				if child.turn == child.facundo.whoiam {
					AlignPAdd(child.facundo.threef, AlignPInit(x-dispo*(s/2+BoolToInt(s == 0)), y-dispo*(s%2-BoolToInt(s == 0)), s))
				} else if child.turn == child.human.whoiam {
					AlignPAdd(child.human.threef, AlignPInit(x-dispo*(s/2+BoolToInt(s == 0)), y-dispo*(s%2-BoolToInt(s == 0)), s))
				}
			}
			if align == 5 && child.turn == child.facundo.whoiam {
				child.facundo.winpot += (1.0 - child.facundo.winpot) * (0.8 * BoolToFloat32(len(child.facundo.five_w) > 0))
				AlignPAdd(child.facundo.five_w, AlignPInit(x-align*(s/2+BoolToInt(s == 0)), y-align*(s%2-BoolToInt(s == 0)), s))
			} else if align == 5 && child.turn == child.human.whoiam {
				child.human.winpot += (1.0 - child.human.winpot) * (0.8 * BoolToFloat32(len(child.human.five_w) > 0))
				AlignPAdd(child.human.five_w, AlignPInit(x-align*(s/2+BoolToInt(s == 0)), y-align*(s%2-BoolToInt(s == 0)), s))
			}
		}
	}
}

//	Put unauthorized and authorized move
func AddPermissiveMove(child GameData) {
	if len(child.facundo.threef) > 0 {
		AddUnauthorizedMove(child.board, child.facundo.whoiam, child.human.whoiam)
		AddAuthorizedThreef(child.board, child.facundo)
	} else {
		AddAuthorizedMove(child.board, child.facundo)
	}
	if len(child.human.threef) > 0 {
		AddUnauthorizedMove(child.board, child.human.whoiam, child.facundo.whoiam)
		AddAuthorizedThreef(child.board, child.human)
	} else {
		AddAuthorizedMove(child.board, child.human)
	}
}

//	Add alls unauthorized move for player
func AddUnauthorizedMove(board [][]int, curPlayer, othPlayer int) {
	var xmax, ymax int
	var x, y int
	var ix, iy int
	var s int
	var i int

	xmax = len(board)
	ymax = len(board[0])
	for x = 0; x <= xmax; x++ {
		for y = 0; y <= ymax; y++ {
			for s = 0; s < 4 && board[x][y] == curPlayer; s++ {
				ix, iy = x+i*(s/2+BoolToInt(s == 0)), y+i*(s%2-BoolToInt(s == 0))
				for i = 0; ix >= 0 && iy >= 0 && ix < xmax && ymax < ymax && board[ix][iy] == curPlayer; i++ {
					ix, iy = x+i*(s/2+BoolToInt(s == 0)), y+i*(s%2-BoolToInt(s == 0))
				}
				if ix >= 0 && iy >= 0 && ix < xmax && ymax < ymax && board[ix][iy] <= 0 && i > 1 {
					AddUnauthorizedPawn(board, ix, iy, curPlayer, othPlayer)
				}
				ix, iy = x-i*(s/2+BoolToInt(s == 0)), y-i*(s%2-BoolToInt(s == 0))
				if ix >= 0 && iy >= 0 && ix < xmax && ymax < ymax && board[ix][iy] <= 0 && i > 1 {
					AddUnauthorizedPawn(board, ix, iy, curPlayer, othPlayer)
				}
			}
		}
	}
}

//	Add Unauthorized move value for current player
func AddUnauthorizedPawn(board [][]int, x, y, curPlayer, othPlayer int) {
	if board[x][y] == -1*othPlayer {
		board[x][y] -= curPlayer
	} else {
		board[x][y] = -1 * curPlayer
	}
}

//	Let Threef free
func AddAuthorizedThreef(board [][]int, cur Player) {
	var x, y int
	var xmax, ymax int
	var i, j, lenThreef int

	lenThreef = len(cur.threef)
	xmax, ymax = len(board), len(board[0])
	for i = 0; i < lenThreef; i++ {
		for j = 0; j < 5; j++ {
			x = cur.threef[i].pos.x + j*(cur.threef[i].dir/2+BoolToInt(cur.threef[i].dir == 0))
			y = cur.threef[i].pos.y + j*(cur.threef[i].dir%2-BoolToInt(cur.threef[i].dir == 0))
			if x < 0 || x >= xmax || y < 0 || y >= ymax {
				continue
			} else if board[x][y] == -3 || board[x][y] == -1*cur.whoiam {
				board[x][y] += cur.whoiam
			}
		}
	}
}

//	Add Authorized move value for current player
func AddAuthorizedMove(board [][]int, player Player) {
	var x, y int
	var xmax, ymax int

	xmax, ymax = len(board), len(board[0])
	for x = 0; x < xmax; x++ {
		for y = 0; y < ymax; y++ {
			if board[x][y] == -3 || board[x][y] == -1*player.whoiam {
				board[x][y] += player.whoiam
			}
		}
	}
}

//	Check ate pawn (around player.Coord)
func CheckEatPawn(child GameData, pawn Pawns) {
	var x, y, px, py int
	var otherPlayer int

	otherPlayer = GetOtherTurn(child)
	for x = -3; x <= 3; x += 3 {
		for y = -3; y <= 3; y += 3 {
			px, py = pawn.x + x, pawn.y + y
			if (px >= 0 || py >= 0 || px < len(child.board) || py < len(child.board[0])) && child.board[px][py] == child.turn && child.board[px-(x/3)][py-(y/3)] == otherPlayer && child.board[px-(2*x/3)][py-(2*y/3)] == otherPlayer {
				AddAteNumPlayer(child, px-(x/3), py - (y/3),px - (2*x/3), py-(2*y/3))
				CheckAlignement(child, Pawns{
					x: px - (x / 3),
					y: py - (y / 3)})
				CheckAlignement(child, Pawns{
					x: px - (2 * x / 3),
					y: py - (2 * y / 3)})
				//CheckUnauthorizetMove(child, Pawns{
				//	x: pawn.x + (x % 2),
				//	y: pawn.y + (y % 2)})
			}
		}
	}
	if child.turn != child.facundo.whoiam {
		CheckAlignPawnPlayer(child.board, child.facundo)
	} else if child.turn != child.human.whoiam {
		CheckAlignPawnPlayer(child.board, child.human)
	}
}

//	Check all Alignement of Pawn registered
func CheckAlignPawnPlayer(board [][]int, player Player) {
	var lenThreef, i int
	var dispo int

	lenThreef = len(player.threef)
	for i = 0; i < lenThreef; i++ {
		dispo, _ = CheckAlignPawnLocal(board, player.threef[i], player.whoiam)
		if dispo < 5 {
			player.threef[i] = player.threef[len(player.threef)-1]
			player.threef = player.threef[:len(player.threef)-1]
		}
	}
	lenThreef = len(player.five_w)
	for i = 0; i < lenThreef; i++ {
		dispo, _ = CheckAlignPawnLocal(board, player.five_w[i], player.whoiam)
		if dispo < 5 {
			player.five_w[i] = player.five_w[len(player.five_w)-1]
			player.five_w = player.five_w[:len(player.five_w)-1]
		}
	}
}

//	Check Current Alignement of Pawn
func CheckAlignPawnLocal(board [][]int, threef AlignP, whoiam int) (int, int) {
	var x, y, n int
	var dispo, align int

	for dispo, align, n = 0, 0, 0; n < 5; n++ {
		x = threef.pos.x + n*(threef.dir/2+BoolToInt(threef.dir == 0))
		y = threef.pos.y + n*(threef.dir%2-BoolToInt(threef.dir == 0))
		if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || (board[x][y] != whoiam && board[x][y] > 0) {
			return 0, 0
		} else if board[x][y] == whoiam {
			dispo++
			align++
		} else {
			if board[x][y] <= 0 && align > 0 {
				dispo++
			} else {
				dispo = BoolToInt(board[x][y] <= 0)
			}
			align = 0
		}
	}
	return dispo, align
}

//	Increase value of Pawn ate by player
func AddAteNumPlayer(child GameData, x1, y1, x2, y2 int) {
	child.board[x1][y1] = 0
	child.board[x2][y2] = 0
	if child.turn == child.facundo.whoiam {
		child.facundo.atenum++
	} else if child.turn == child.human.whoiam {
		child.human.atenum++
	}
}

//	Add pawn and new authorized place on board
func AddPawnOnBoard(child GameData, pawn Pawns) {
	child.board[pawn.x][pawn.y] = child.turn
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if child.board[pawn.x+x][pawn.y+y] == -4 {
				child.board[pawn.x+x][pawn.y+y] = 0
			}
		}
	}
	if child.move.x == -1 || child.move.y == -1 {
		child.move = pawn
	}
}

//	Check probability (set prob of winning for two party)
func CheckWinLose(player Player, turn int) {
	if turn == player.whoiam && len(player.five_w) > 0 {
		player.winpot = 1.0
	} else {
		player.winpot = (float32)(player.atenum) / 5.0
	}
	player.winpot += (1.0 - player.winpot) * (1.0 - (1.0 / (float32)(len(player.threef)+1)))
	player.winpot += (1.0 - player.winpot) * (0.8 * BoolToFloat32(len(player.five_w) > 0))
}

//	MinMax add couche
func MinMax(childs GameData, pawn NextPawns, link chan GameData) {
	var child GameData

	GameDataCopy(childs, child)
	AddPawnOnBoard(child, pawn.pawn_p)
	CheckWinLose(child.human, child.turn)
	CheckWinLose(child.facundo, child.turn)
	CheckEatPawn(child, pawn.pawn_p)
	CheckAlignement(child, pawn.pawn_p)
	AddPermissiveMove(child)
	child.turn = GetOtherTurn(child)
	child.deep--
	link <- child
}
