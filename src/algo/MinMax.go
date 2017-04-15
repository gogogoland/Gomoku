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
 * 		Check victory probability for better IA
 *
 * 		NOTHING
**/

/**
 * 	"Check Alignement"
 *
 * 	Check if new pawn made an alignement with other pawns (in 8 direction)
 * 	(limAlign is for aligned pawn with isolate space between them)
 * 	(dispo is for available space)
 * 	(winAlign is for uninterruptable alignement)
 * 	If so save it for:
 * 		threef (2-3 aligned and 5 available) SliceAP
 * 		five_w (5 aligned) SliceAP
**/
func (child *GameData) CheckAlignement(pawn Pawns) {
	var x, y, s int
	var check, winAlign, limAlign, dispo, eat int

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
					eat = BoolToInt(child.board[x][y] > 0)
					limAlign = 0
				}
				winAlign = 0
			}

			if eat == 1 && winAlign == 2 && (check == 0 || check == 1) {
				child.GetOtherPlayer(child.turn).eating.Add(PawnsInit(x-(limAlign+1)*(s/2+BoolToInt(s == 0)), y-(limAlign+1)*(s%2-BoolToInt(s == 0))))
			}
			if limAlign == 3 && dispo == 5 {
				child.GetPlayer(child.turn).threef.Add(AlignPInit(x-(limAlign+1)*(s/2+BoolToInt(s == 0)), y-(limAlign+1)*(s%2-BoolToInt(s == 0)), s))
			}
			if limAlign == 4 && dispo == 5 {
				child.GetPlayer(child.turn).four_w.Add(AlignPInit(x-(limAlign)*(s/2+BoolToInt(s == 0)), y-(limAlign)*(s%2-BoolToInt(s == 0)), s))
			}
			if winAlign == 5 {
				child.GetPlayer(child.turn).five_w.Add(AlignPInit(x-(winAlign-1)*(s/2+BoolToInt(s == 0)), y-(winAlign-1)*(s%2-BoolToInt(s == 0)), s))
				child.GetPlayer(child.turn).threef.Del(AlignPInit(x-(winAlign-1)*(s/2+BoolToInt(s == 0)), y-(winAlign-1)*(s%2-BoolToInt(s == 0)), s))
			}
		}
	}
}

/**
 * 	Check Ate Pawn
 *
 * 	Check if newly placed pawn eat other player's pawns
 * 	If so, add number of pawns ate for current player
 * 	And remove other player's pawns
**/
func (child *GameData) CheckEatPawn(pawn Pawns) {
	var x, y, px, py int
	var otherPlayer int

	otherPlayer = child.GetOtherTurn()
	for x = -3; x <= 3; x += 3 {
		for y = -3; y <= 3; y += 3 {
			px, py = pawn.x+x, pawn.y+y
			if px >= 0 && py >= 0 && px < child.maxx && py < child.maxy && child.board[px-(x/3)][py-(y/3)] == otherPlayer && child.board[px-(2*x/3)][py-(2*y/3)] == otherPlayer {
				if child.board[px][py] == child.turn {
					child.AddAteNumPlayer(px-(x/3), py-(y/3), px-(2*x/3), py-(2*y/3))
					child.CheckAlignement(Pawns{
						x: px - (x / 3),
						y: py - (y / 3)})
					child.CheckAlignement(Pawns{
						x: px - (2 * x / 3),
						y: py - (2 * y / 3)})
					if x != 0 && y != 0 {
						child.DiagonalEatRemovePawn(pawn, px, py)
					}
				} else if child.board[px][py] <= 0 {
					child.GetPlayer(child.turn).eating.Add(PawnsInit(px, py))
				}
			}
		}
	}
}

/**
 * 	Remove Unavailable Box After Diagonal Ate"
 *
 * 	Remove box availability for the six that are between diagonale ater pawns
 * 	Using "Remove Unavailable Box" (see below)
**/
func (child *GameData) DiagonalEatRemovePawn(p Pawns, x, y int) {
	child.RemoveUnavailableMove(p.x+2-BoolToInt(p.x > x)*4, p.y)
	child.RemoveUnavailableMove(p.x, p.y+2-BoolToInt(p.y > y)*4)
	child.RemoveUnavailableMove(p.x+1-BoolToInt(p.x > x)*2, p.y+2-BoolToInt(p.y > y)*4)
	child.RemoveUnavailableMove(x+1-BoolToInt(x > p.x)*2, y+2-BoolToInt(y > p.y)*4)
	child.RemoveUnavailableMove(x+2-BoolToInt(x > p.x)*4, y)
	child.RemoveUnavailableMove(x, y+2-BoolToInt(y > p.y)*4)
}

/**
 * 	"Remove Unavailable Box"
 *
 * 	Remove box that it's not close to a pawn
**/
func (child *GameData) RemoveUnavailableMove(px, py int) {
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

/**
 * 	"Check Saved Alignement Of Player"
 *
 * 	Verify saved alignement with "Check Current Alignement" (see below)
 * 	If avaible box is under five, delete it
 * 	If eating pawn is unavailable, delete it
**/
func (player *Player) CheckAlignPawnPlayer(Board [][]int) {
	var lenSlice, i int
	var dispo, align int

	lenSlice = len(player.threef)
	for i = 0; i < lenSlice; i++ {
		dispo, align = CheckAlignPawnLocal(Board, player.threef[i], player.whoami)
		if dispo < 5 && align < 3 {
			player.threef[i] = player.threef[len(player.threef)-1]
			player.threef = player.threef[:len(player.threef)-1]
			lenSlice--
		}
	}

	lenSlice = len(player.four_w)
	for i = 0; i < lenSlice; i++ {
		dispo, align = CheckAlignPawnLocal(Board, player.four_w[i], player.whoami)
		if dispo < 5 {
			player.four_w[i] = player.four_w[len(player.four_w)-1]
			player.four_w = player.four_w[:len(player.four_w)-1]
			lenSlice--
		}
	}

	lenSlice = len(player.five_w)
	for i = 0; i < lenSlice; i++ {
		dispo, align = CheckAlignPawnLocal(Board, player.five_w[i], player.whoami)
		if align < 5 {
			player.five_w[i] = player.five_w[len(player.five_w)-1]
			player.five_w = player.five_w[:len(player.five_w)-1]
			lenSlice--
		}
	}

	lenSlice = len(player.eating)
	for i = 0; i < lenSlice; i++ {
		if player.eating[i].x > 18 || player.eating[i].x < 0 || player.eating[i].y > 18 || player.eating[i].y < 0 || Board[player.eating[i].x][player.eating[i].y] > 0 {
			player.eating[i] = player.eating[len(player.eating)-1]
			player.eating = player.eating[:len(player.eating)-1]
			lenSlice--
		}
	}
}

/**
 * 	"Check Current Alignement"
 *
 * 	Check avaible and occuped box for specific alignement
 * 	Return avaible box and concret alignement
**/
func CheckAlignPawnLocal(Board [][]int, threef AlignP, whoami int) (int, int) {
	var x, y, n int
	var dispo, align int

	for dispo, align, n = 0, 0, 0; n < 5; n++ {
		x = threef.pos.x + n*(threef.dir/2+BoolToInt(threef.dir == 0))
		y = threef.pos.y + n*(threef.dir%2-BoolToInt(threef.dir == 0))
		if x < 0 || y < 0 || x >= len(Board) || y >= len(Board[0]) || (Board[x][y] != whoami && Board[x][y] > 0) {
			return 0, 0
		} else if Board[x][y] == whoami {
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

/**
 * 	"Add Ate Number Of Player"
 *
 * 	Increase value of Pawn eaten by player
**/
func (child *GameData) AddAteNumPlayer(x1, y1, x2, y2 int) {
	child.board[x1][y1] = 0
	child.board[x2][y2] = 0
	if child.turn == child.facundo.whoami {
		child.facundo.atenum++
	} else if child.turn == child.human.whoami {
		child.human.atenum++
	}
}

/**
 * 	"Add Pawn On Board"
 *
 * 	If available box, put a pawn from current player on
 * 	And then, put new authorized place on Board for Facundo
 * 	If player's first move, save pawn's coordinate (usefull for Facundo)
 * 	Return false if box is unavailable else return true.
**/
func (child *GameData) AddPawnOnBoard(pawn Pawns) bool {

	if pawn.x < 0 || pawn.x > 18 || pawn.y < 0 || pawn.y > 18 || child.board[pawn.x][pawn.y] > 0 {
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

/**
 * "Check Win/Lose Probability"
 *
 * 	With "Check Player Win/Lose Probability" function (see below),
 * 	Check each player victory, if so, put other player at minus
**/
func (child *GameData) CheckWinLose() {
	var humanWin, facundoWin float32

	humanWin = child.human.CheckPlayerWinLose(child.turn)
	facundoWin = child.facundo.CheckPlayerWinLose(child.turn)
	child.human.winpot -= facundoWin
	child.facundo.winpot -= humanWin
}

/**
 * 	"Check Player Win/Lose Probability"
 *
 * 	Calcul win/lose probability for each player given the below requirement:
 * 	Five pawns aligned from other player (Win for other player)
 * 	Five peer of pawns ate (Win for current player)
 *
 * 	Check also other data :
 * 	(number of threef (three free),peer of pawns eaten, eatable pawns, four pawns)
**/
func (player *Player) CheckPlayerWinLose(turn int) float32 {
	if turn != player.whoami && len(player.five_w) > 0 {
		player.winpot = 1.0
	} else {
		player.winpot = (float32)(player.atenum) / 5.0
	}
	player.winpot += (1.0 - player.winpot) * (0.75 * BoolToFloat32(len(player.five_w) > 0))
	player.winpot += (1.0 - player.winpot) * (0.45 - (0.45 / (float32)(len(player.threef)+1)))
	player.winpot += (1.0 - player.winpot) * (0.3 - (0.3 / (float32)(len(player.threef)+1)))
	player.winpot += (1.0 - player.winpot) * (0.15 - (0.15 / (float32)(len(player.eating)+1)))
	return 2 * BoolToFloat32(player.winpot >= 1.0)
}

/**
 * 	"MinMax"
 *
 * 	Execute "Turn Process" (see below) on new child
 * 	Return newly made child with applyed move and result of "Turn Process"
 *
 * 	(the link part leads to an use of channels for goroutine)
**/
func MinMax(childs GameData, pawn NextPawns /*, link chan GameData*/) (GameData, int) {
	var child GameData
	var denied int

	child = childs.Copy()
	denied = (&child).TurnProcess(pawn.pawn_p)
	child.deep--
	//link//link <- child, denied
	return child, denied
}

/**
 * 	"Player Turn"
 *
 * 	Execute "Turn Process" (see below) on copied GameData
 * 	If move is authorized, apply "Turn Process" to original GameData
 * 	Return 1 if the move is denied, else 0
**/
func (Mom *GameData) PlayerTurn(pawn Pawns) int {
	var testMove GameData
	var denied int

	testMove = Mom.Copy()
	denied = (&testMove).TurnProcess(pawn)
	if denied == 0 {
		denied = Mom.TurnProcess(pawn)
	}
	return denied
}

/**
 * 	"Turn Process"
 *
 * 	Execute move of one player and apply rules (upper functions)
 * 	Calcul victory condition for each players
 * 	Return 1 if the move is denied (unavailable box or multiple free three)
 * 	Else return 0
**/
func (child *GameData) TurnProcess(pawn Pawns) int {
	var FreeThree int

	if child.AddPawnOnBoard(pawn) == false {
		return 1
	}

	FreeThree = len(child.GetPlayer(child.turn).threef)
	if FreeThree <= 0 {
		FreeThree = 1
	}
	child.CheckAlignement(pawn)
	if len(child.GetPlayer(child.turn).threef) > FreeThree && FreeThree > 0 {
		return 1
	} else if child.GetPlayer(child.turn).pawn_p.Compare(PawnsInit(-1, -1)) {
		child.GetPlayer(child.turn).pawn_p = pawn
	}
	child.CheckEatPawn(pawn)
	child.GetPlayer(child.GetOtherTurn()).CheckAlignPawnPlayer(child.board)
	child.CheckWinLose()
	return 0
}
