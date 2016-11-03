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
 * 		Put function in kernel
 *
 * 		NOTHING
**/

//	Check if pawn are aligned (for winning move (if win and three free)
func CheckAlignement(child GameData, pawn Pawn) {
	var x, y, s int
	var xfive, yfive int
	var xmax, ymax int
	var check, align, dispo int

	align = 0
	xmax = len(child.board)
	ymax = len(child.board[0])
	for s = 0; s < 4 && align < 5; s++ {
		dispo = 0
		align = 0
		for check = -4; check <= 4 && align < 5; check++ {
			x = pawn.x + check*(s/2+!s)
			y = pawn.y + check*(s%2-!s)
			if x < 0 || y < 0 || x >= xmax || y >= ymax {
				continue
			} else if child.board[x][y] == child.turn {
				dispo++
				align++
			} else {
				if child.board[x][y] <= 0 && align > 0 {
					dispo++
				} else {
					dispo = (child.board[x][y] <= 0)
				}
				align = 0
			}
			if align >= 3 && dispo == 5 {
				if child.turn == child.facundo.whoiam {
					AddAlignP(child.facundo, AlignPInit(x-dispo*(s/2+!s), y-dispo*(s%2-!s), s))
				} else if child.turn == child.human.whoiam {
					AddAlignP(child.human, AlignPInit(x-dispo*(s/2+!s), y-dispo*(s%2-!s), s))
				}
			}
			if align == 5 && child.turn == child.facundo.whoiam {
				child.facundo.winpot += (1.0 - player.winpot) * (0.8 * (len(player.five_w) > 0))
				AddAlignP(child.facundo, AlignPInit(x-align*(s/2+!s), y-align*(s%2-!s), s))
			} else if align == 5 && child.turn == child.human.whoiam {
				child.human.winpot += (1.0 - player.winpot) * (0.8 * (len(player.five_w) > 0))
				AddAlignP(child.human, AlignPInit(x-align*(s/2+!s), y-align*(s%2-!s), s))
			}
		}
	}
}

//	Put unauthorized and authorized move
func AddPermissiveMove(child GameData) {
	if len(player.facundo.threef) > 0 {
		AddUnauthorizedMove(board, child.facundo.whoiam, child.human.whoiam)
		AddAuthorizedThreef(board, child.facundo)
	} else {
		AddAuthorizedMove(board, child.facundo)
	}
	if len(player.human.threef) > 0 {
		AddUnauthorizedMove(board, child.human.whoiam, child.facundo.whoiam)
		AddAuthorizedThreef(board, child.human)
	} else {
		AddAuthorizedMove(board, child.human)
	}
}

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
				ix, iy = x+i*(s/2+!s), y+i*(s%2-!s)
				for i = 0; ix >= 0 && iy >= 0 && ix < xmax && ymax < ymax && board[ix][iy] == curPlayer; i++ {
					ix, iy = x+i*(s/2+!s), y+i*(s%2-!s)
				}
				if ix >= 0 && iy >= 0 && ix < xmax && ymax < ymax && board[ix][iy] <= 0 && i > 1 {
					AddUnauthorizedPawn(board, ix, iy, curPlayer, othPlayer)
				}
				ix, iy = x-i*(s/2+!s), y-i*(s%2-!s)
				if ix >= 0 && iy >= 0 && ix < xmax && ymax < ymax && board[ix][iy] <= 0 && i > 1 {
					AddUnauthorizedPawn(board, ix, iy, curPlayer, othPlayer)
				}
			}
		}
	}
}

func AddUnauthorizedPawn(board [][]int, x, y, curPlayer, othPlayer int) {
	if board[x][y] == -1*othPlayer {
		board[x][y] -= curPlayer
	} else {
		board[x][y] = -1 * curPlayer
	}
}

func AddAuthorizedThreef(board [][]int, cur Player) {
	var x, y int
	var xmax, ymax int
	var i, j, len int

	len = len(cur.threef)
	xmax, ymax = len(board), len(board[0])
	for i = 0; i < len; i++ {
		for j = 0; j < 5; j++ {
			x = cur.threef[i].pos.x + j*(cur.threef[i].dir/2+!cur.threef[i].dir)
			y = cur.threef[i].pos.y + j*(cur.threef[i].dir%2-!cur.threef[i].dir)
			if x < 0 || x >= xmax || y < 0 || y >= ymax {
				continue
			} else if board[x][y] == -3 || board[x][y] == -1*cur.whoiam {
				board[x][y] += cur.whoiam
			}
		}
	}
}

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
func CheckEatPawn(child GameData, pawn Pawn) {
	var ate bool
	var x, y int
	var otherPlayer int

	ate = false
	otherPlayer = GetOtherTurn(child)
	for x = -2; x <= 2; x += 2 {
		for y = -2; y <= 2; y += 2 {
			if (x >= 0 || y >= 0 || x < len(child.board) || y < len(child.board[0])) && child.board[pawn.x+x][pawn.y+y] == child.turn && child.board[pawn.x+(x%2)][pawn.y+(y%2)] == otherPlayer {
				AddAteNumPlayer(child, pawn.x+(x%2), pawn.y+(y%2))
				ate = true
				CheckAlignement(child, Pawn{
					x: pawn.x + (x % 2),
					y: pawn.y + (y % 2)})
				CheckUnauthorizetMove(child, Pawn{
					x: pawn.x + (x % 2),
					y: pawn.y + (y % 2)})
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
	var x, y int
	var xmax, ymax int
	var len, i int
	var dispo, align int

	xmax, ymax = len(board), len(board[0])
	len = len(player.threef)
	for i = 0; i < len; i++ {
		dispo, align = CheckAlignPawnLocal(board, player.threef[i], player.wohiam)
		if dispo < 5 {
			player.threef[i] = player.threef[len(player.threef)-1]
			player.threef[len(player.threef)-1] = nil
			player.threef = player.threef[:len(player.threef)-1]
		}
	}
	len = len(player.five_w)
	for i = 0; i < len; i++ {
		dispo, align = CheckAlignPawnLocal(board, player.five_w[i], player.wohiam)
		if dispo < 5 {
			player.five_w[i] = player.five_w[len(player.five_w)-1]
			player.five_w[len(player.five_w)-1] = nil
			player.five_w = player.five_w[:len(player.five_w)-1]
		}
	}
}

func CheckAlignPawnLocal(board [][]int, threef AlignP, wohiam int) (int, int) {
	var x, y, n int
	var dispo, align int

	for dispo, align, n = 0, 0, 0; n < 5; n++ {
		x = threef.pos.x + j*(threef.dir/2+!threef.dir)
		y = threef.pos.y + j*(threef.dir%2-!threef.dir)
		if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || (board[x][y] != whoiam && board[x][y] > 0) {
			return 0, 0
		} else if board[x][y] == whoiam {
			dispo++
			align++
		} else {
			if child.board[x][y] <= 0 && align > 0 {
				dispo++
			} else {
				dispo = (child.board[x][y] <= 0)
			}
			align = 0
		}
	}
	return dispo, align
}

func AddAteNumPlayer(child GameData, x, y int) {
	child.board[pawn.x+x][pawn.y+y] = 0
	if child.turn == child.facundo.whoiam {
		child.facundo.ate++
	} else if child.turn == child.human.whoiam {
		child.human.ate
	}
}

//	Add pawn and new authorized place on board
func AddPawnOnBoard(child GameData, pawn Pawn) {
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
		player.winpot = player.atenum / 5
	}
	player.winpot += (1.0 - player.winpot) * (1.0 - (1.0 / (len(player.threef) + 1)))
	player.winpot += (1.0 - player.winpot) * (0.8 * (len(player.five_w) > 0))
}

//	MinMax add couche
func MinMax(childs GameData, pawn NextPawns, link chan GameData) {
	var child GameData

	CopyGameData(childs, child)
	AddPawnOnBoard(child, pawn)
	CheckWinLose(child.human, child.turn)
	CheckWinLose(child.facundo, child.turn)
	CheckEatPawn(child, pawn)
	CheckAlignement(child, pawn)
	AddPermissiveMove(child.board, child.facundo)
	AddPermissiveMove(child.board, child.human)
	child.turn = GetOtherTurn(child)
	child.deep--
	link <- child
}

//	Check time remaining for calcul
/*func CheckTimeLeft(remain, occur float64) float64 {
	remain -= occu
	if remain < 0.0 {
		remain = 0.0
	}
	return remain
}*/
