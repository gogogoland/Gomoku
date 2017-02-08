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

import "fmt"

func PrintHuman(popipo GameData) {
	fmt.Println("Human Data :", popipo.human)
}

/**
 * Adding pawn layer and get probabilty of win
 *
 * TODO:
 * 		Seek comments inside functions
 *		Check unauthorized move left by diagonal ate move
 *		Change Permissive (freethree) Move for when the pawn is set
 *		Check space left between pawns
 *		See CheckEatPawn()
 *
 * 		NOTHING
**/

//	Check if pawn are aligned (for winning move (if win and three free)
func (child *GameData) CheckAlignement(pawn Pawns) {
	var x, y, s int
	var check, winAlign, limAlign, dispo int

	winAlign, limAlign = 0, 0
	for s = 0; s < 4 && winAlign < 5; s++ {
		dispo, winAlign, limAlign = 0, 0, 0
		for check = -4; check <= 4 && winAlign < 5; check++ {
			x = pawn.x + check*(s/2+BoolToInt(s == 0))
			y = pawn.y + check*(s%2-BoolToInt(s == 0))
			if x < 0 || y < 0 || x >= child.maxx || y >= child.maxy {
				continue
			} else if child.board[x][y] == child.turn {
				dispo++
				winAlign++
				limAlign++
			} else {
				if child.board[x][y] <= 0 && winAlign > 0 {
					dispo++
				} else {
					dispo = BoolToInt(child.board[x][y] <= 0)
					limAlign = 0
				}
				winAlign = 0
			}
			//	Maybe dispo == 6
			if limAlign >= 3 && dispo == 5 {
				if child.turn == child.facundo.whoiam {
					child.facundo.threef.Add(AlignPInit(x-(limAlign+1)*(s/2+BoolToInt(s == 0)), y-(limAlign+1)*(s%2-BoolToInt(s == 0)), s))
				} else if child.turn == child.human.whoiam {
					child.human.threef.Add(AlignPInit(x-(limAlign+1)*(s/2+BoolToInt(s == 0)), y-(limAlign+1)*(s%2-BoolToInt(s == 0)), s))
				}
			}
			if winAlign == 5 {
				if child.turn == child.facundo.whoiam {
					child.facundo.winpot += (1.0 - child.facundo.winpot) * (0.8 * BoolToFloat32(len(child.facundo.five_w) > 0))
					child.facundo.five_w.Add(AlignPInit(x-winAlign*(s/2+BoolToInt(s == 0)), y-winAlign*(s%2-BoolToInt(s == 0)), s))
				} else if child.turn == child.human.whoiam {
					child.human.winpot += (1.0 - child.human.winpot) * (0.8 * BoolToFloat32(len(child.human.five_w) > 0))
					child.human.five_w.Add(AlignPInit(x-winAlign*(s/2+BoolToInt(s == 0)), y-winAlign*(s%2-BoolToInt(s == 0)), s))
				}
			}
		}
	}
}

//	Put unauthorized and authorized move
func (child *GameData) AddPermissiveMove() {
	if len(child.facundo.threef) > 0 {
		child.AddUnauthorizedMove(child.facundo.whoiam, child.human.whoiam)
		child.AddAuthorizedThreef(child.facundo)
	} else {
		child.AddAuthorizedMove(child.facundo.whoiam)
	}
	if len(child.human.threef) > 0 {
		child.AddUnauthorizedMove(child.human.whoiam, child.facundo.whoiam)
		child.AddAuthorizedThreef(child.human)
	} else {
		child.AddAuthorizedMove(child.human.whoiam)
	}
}

//	Add alls unauthorized move for player
func (child *GameData) AddUnauthorizedMove(curPlayer, othPlayer int) {
	var x, y int
	var ix, iy int
	var s int
	var i int

	for x = 0; x < child.maxx; x++ {
		for y = 0; y < child.maxy; y++ {
			for s = 0; s < 4 && child.board[x][y] == curPlayer; s++ {
				ix, iy = x+i*(s/2+BoolToInt(s == 0)), y+i*(s%2-BoolToInt(s == 0))
				for i = 0; ix >= 0 && iy >= 0 && ix < child.maxx && child.maxy < child.maxy && child.board[ix][iy] == curPlayer; i++ {
					ix, iy = x+i*(s/2+BoolToInt(s == 0)), y+i*(s%2-BoolToInt(s == 0))
				}
				if ix >= 0 && iy >= 0 && ix < child.maxx && child.maxy < child.maxy && child.board[ix][iy] <= 0 && i > 1 {
					child.AddUnauthorizedPawn(ix, iy, curPlayer, othPlayer)
				}
				ix, iy = x-i*(s/2+BoolToInt(s == 0)), y-i*(s%2-BoolToInt(s == 0))
				if ix >= 0 && iy >= 0 && ix < child.maxx && child.maxy < child.maxy && child.board[ix][iy] <= 0 && i > 1 {
					child.AddUnauthorizedPawn(ix, iy, curPlayer, othPlayer)
				}
			}
		}
	}
}

//	Add Unauthorized move value for current player
func (child *GameData) AddUnauthorizedPawn(x, y, curPlayer, othPlayer int) {
	if child.board[x][y] == -1*othPlayer || child.board[x][y] == 0 {
		child.board[x][y] -= curPlayer
	}
}

//	Let Threef free
func (child *GameData) AddAuthorizedThreef(player Player) {
	var x, y int
	var i, j, lenThreef int

	lenThreef = len(player.threef)
	for i = 0; i < lenThreef; i++ {
		for j = 0; j < 5; j++ {
			x = player.threef[i].pos.x + j*(player.threef[i].dir/2+BoolToInt(player.threef[i].dir == 0))
			y = player.threef[i].pos.y + j*(player.threef[i].dir%2-BoolToInt(player.threef[i].dir == 0))
			if x < 0 || x >= child.maxx || y < 0 || y >= child.maxy {
				continue
			} else if child.board[x][y] == -3 || child.board[x][y] == -1*player.whoiam {
				child.board[x][y] += player.whoiam
			}
		}
	}
}

//	Add Authorized move value for current player
func (child *GameData) AddAuthorizedMove(whoiam int) {
	var x, y int

	for x = 0; x < child.maxx; x++ {
		for y = 0; y < child.maxy; y++ {
			if child.board[x][y] == -1*whoiam || child.board[x][y] == -3 {
				child.board[x][y] += whoiam
			}
		}
	}
}

//	Check ate pawn (around player.Coord)
func (child *GameData) CheckEatPawn(pawn Pawns) {
	var x, y, px, py int
	var otherPlayer int

	otherPlayer = child.GetOtherTurn()
	for x = -3; x <= 3; x += 3 {
		for y = -3; y <= 3; y += 3 {
			px, py = pawn.x+x, pawn.y+y
			if px >= 0 && py >= 0 && px < child.maxx && py < child.maxy && child.board[px][py] == child.turn && child.board[px-(x/3)][py-(y/3)] == otherPlayer && child.board[px-(2*x/3)][py-(2*y/3)] == otherPlayer {
				child.AddAteNumPlayer(px-(x/3), py-(y/3), px-(2*x/3), py-(2*y/3))
				child.CheckAlignement(Pawns{
					x: px - (x / 3),
					y: py - (y / 3)})
				child.CheckAlignement(Pawns{
					x: px - (2 * x / 3),
					y: py - (2 * y / 3)})
				//child.DiagonalEatRemovePawn(pawn, Pawns{
				//	x: px,
				//	y: py})
				//CheckUnauthorizetMove(child, Pawns{
				//	x: pawn.x + (x % 2),
				//	y: pawn.y + (y % 2)})
			}
		}
	}
	if child.turn != child.facundo.whoiam {
		child.facundo.CheckAlignPawnPlayer(child.board)
	} else if child.turn != child.human.whoiam {
		child.human.CheckAlignPawnPlayer(child.board)
	}
}

//	To RemoveUnabailbeMove for each four box around diagonal pawns ate
//	TODO:
//		Check diagonal unaivable box
func (child *GameData) DiagonalEatRemovePawn(p1, p2 Pawns) {
	if p2.x == 0 || p2.y == 0 {
		return
	}
	child.RemoveUnavaibleMove(p1.x+1-BoolToInt(p1.x < p2.x)*2, p1.y)
	child.RemoveUnavaibleMove(p1.x, p1.y)
	child.RemoveUnavaibleMove(p2.x, p2.y)
	child.RemoveUnavaibleMove(p2.x, p2.y)
}

//	Remove avaible box after diagonal ate
func (child *GameData) RemoveUnavaibleMove(px, py int) {
	var x, y int

	if px < 0 || px >= child.maxx || py < 0 || py >= child.maxy {
		return
	}
	for x = -1; x <= 1; x++ {
		for y = -1; y <= 1; y++ {
			if px+x >= 0 && px+x < child.maxx && py+y >= 0 && py+y < child.maxy && child.board[px+x][py+y] > 0 {
				return
			}
		}
	}
	child.board[px][py] = -4
}

//	Check all Alignement of Pawn registered
func (player *Player) CheckAlignPawnPlayer(Board [][]int) {
	var lenThreef, i int
	var dispo int

	lenThreef = len(player.threef)
	for i = 0; i < lenThreef; i++ {
		dispo, _ = CheckAlignPawnLocal(Board, player.threef[i], player.whoiam)
		if dispo < 5 {
			player.threef[i] = player.threef[len(player.threef)-1]
			player.threef = player.threef[:len(player.threef)-1]
			lenThreef--
		}
	}
	lenThreef = len(player.five_w)
	for i = 0; i < lenThreef; i++ {
		dispo, _ = CheckAlignPawnLocal(Board, player.five_w[i], player.whoiam)
		if dispo < 5 {
			player.five_w[i] = player.five_w[len(player.five_w)-1]
			player.five_w = player.five_w[:len(player.five_w)-1]
			lenThreef--
		}
	}
}

//	Check Current Alignement of Pawn
func CheckAlignPawnLocal(Board [][]int, threef AlignP, whoiam int) (int, int) {
	var x, y, n int
	var dispo, align int

	for dispo, align, n = 0, 0, 0; n < 5; n++ {
		x = threef.pos.x + n*(threef.dir/2+BoolToInt(threef.dir == 0))
		y = threef.pos.y + n*(threef.dir%2-BoolToInt(threef.dir == 0))
		/*	} else if child.board[x][y] == child.turn {
				dispo++
				winAlign++
				limAlign++
			} else {
				if child.board[x][y] <= 0 && winAlign > 0 {
					dispo++
				} else {
					dispo = BoolToInt(child.board[x][y] <= 0)
					limAlign = 0
				}
				winAlign = 0
			}*/
		if x < 0 || y < 0 || x >= len(Board) || y >= len(Board[0]) || (Board[x][y] != whoiam && Board[x][y] > 0) {
			return 0, 0
		} else if Board[x][y] == whoiam {
			dispo++
			align++
		} else {
			if Board[x][y] <= 0 && align > 0 {
				dispo++
			} else {
				dispo = BoolToInt(Board[x][y] <= 0)
			}
			align = 0
		}
	}
	return dispo, align
}

//	Increase value of Pawn ate by player
func (child *GameData) AddAteNumPlayer(x1, y1, x2, y2 int) {
	child.board[x1][y1] = 0
	child.board[x2][y2] = 0
	if child.turn == child.facundo.whoiam {
		child.facundo.atenum++
	} else if child.turn == child.human.whoiam {
		child.human.atenum++
	}
}

//	Add pawn and new authorized place on Board
func (child *GameData) AddPawnOnBoard(pawn Pawns) bool {
	if pawn.x < 0 || pawn.x > 18 || pawn.y < 0 || pawn.y > 18 || child.board[pawn.x][pawn.y] < -2 || child.board[pawn.x][pawn.y] == -1*child.turn || child.board[pawn.x][pawn.y] > 0 {
		return false
	}
	child.board[pawn.x][pawn.y] = child.turn
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if !(pawn.x+x < 0 || pawn.x+x > 18 || pawn.y+y < 0 || pawn.y+y > 18) && child.board[pawn.x+x][pawn.y+y] == -4 {
				child.board[pawn.x+x][pawn.y+y] = 0
			}
		}
	}
	if child.move.x == -1 || child.move.y == -1 {
		child.move = pawn.Copy()
	}
	return true
}

//	Check probability (set prob of winning for two party)
func (player *Player) CheckWinLose(turn int) {
	if turn != player.whoiam && len(player.five_w) > 0 {
		// Check if already five aligned, if so it's a win
		player.winpot = 1.0
	} else {
		// Else get closer to victory by pair of pawns ate
		player.winpot = (float32)(player.atenum) / 5.0
		//fmt.Println("Player :", player.whoiam, "PawnPos :", player.pawn_p, "Ate :", player.atenum, "WinPot :", player.winpot)
	}
	// For the remain potential, get pre-aligned or newly aligned pawns
	player.winpot += (1.0 - player.winpot) * (1.0 - (1.0 / (float32)(len(player.threef)+1)))
	player.winpot += (1.0 - player.winpot) * (0.8 * BoolToFloat32(len(player.five_w) > 0))
}

//	MinMax add couche
func MinMax(childs GameData, pawn NextPawns /*, link chan GameData*/) GameData {
	var child GameData

	child = childs.Copy()
	(&child).TurnProcess(pawn.pawn_p)
	child.deep--
	//link <- child//GOLD
	return child
}

//	Turn process
func (child *GameData) TurnProcess(pawn Pawns) int {
	// Place new move
	if child.AddPawnOnBoard(pawn) == false {
		return 1
	}
	//	Check Ate pawns during this turn
	child.CheckEatPawn(pawn)
	//	Check Alignement for pawns
	child.CheckAlignement(pawn)
	//	Change box value for new place
	child.AddPermissiveMove()
	//	Set new potential of victory
	child.human.pawn_p, child.facundo.pawn_p = child.move.Copy(), child.move.Copy()
	child.human.CheckWinLose(child.turn)
	child.facundo.CheckWinLose(child.turn)
	//	Set turn value for other player
	child.turn = child.GetOtherTurn()
	return 0
}
