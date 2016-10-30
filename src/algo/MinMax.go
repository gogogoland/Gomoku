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
 * TODO:
 * 		Redo check free three
 * 		Add Check victory (check aligned 5 + check ate)
 * 		Add Check winning or losing probability
 * 		put in kernel calcul
 *
 *
 * 		NOTHING
**/

//	Check if there are free three
/*func CheckFreeThree(board [][]int, max Coord, human, facundo Player) {
	pawned, align, check, dir, x, y, xc, yc int
	oldHu, oldIA bool

	oldHu, oldIA, human.threef, facundo.threef = human.threef, facundo.threef, false, false
	for dir = 0; dir < 4; dir++ {
		for y = 0; y < max.y; y++ {
			for x = 0; x < max.x ; x++ {
				align = 0;
				if !board[x][y] || (board[x][y] ==  human.witch && human.threef) || (board[y][x] == facundo.witch && facundo.threef) {
					continue
				}
				for check = -4; check < 5 && align < 5; check++ {
					if (dir / 2) && (player.pos.y + check < 0) {
						check = -player.pos.y
					} else if (dir / 2) && (player.pos.y + check >= max.y) {
						break
					} else if (!dir) && (player.pos.y - check >= max.y) {
						check = player.pos.y - max.y
					} else if (!dir) && (player.pos.y - check < 0) {
						break
					}
					if ((dir % 2) || (!dir)) && (player.pos.x + check < 0) {
						check = -player.pos.x
					} else if ((dir % 2) || (!dir)) && (player.pos.x + check >= max.x) {
						break
					}
					yc = y + ((dir / 2) * check) - (!dir) * check)
					xc = x + ((dir % 2) * check) + (!dir) * check)
					align++
					if board[yc][xc] == board[y][c] {
						pawned++
					} else if (board[yc][xc] == human.witch && board[y][x] == facundo.witch) || (facundo.witch == board[yc][xc] && human.witch == board[y][x]) {
						pawned = 0
						aling = 0
					}
					if align == 5 && pawned >= 3 {
						if board[x][y] ==  human.witch {
							human.threef = true
						} else if board[y][x] == facundo.witch {
							facundon.threef = true
						}
					}
				}
			}
		}
	}
	if oldHu != human.threef {
		if oldHu {
//			reset unauthorizet space
		} else {
//			set unauthorizet space
		}
	}
	if oldIA != facundo.threef {
		if oldIA {
//			reset unauthorizet space
		} else {
//			set unauthorizet space
		}
	}
}*/
/*
//	Check unauthorized new move (double free-three)
func CheckAuthorizedMove([][]int board, Coords max, pos) {
	pawned, align, check, dir, x, y, xc, yc int
//	If so, unauthorized third place
	return ((align / 5) * player.witch * player.win)
}
*/
//	Check ate pawn (around player.Coord)
func CheckEatPawn(childs GameData, pawn Pawn) {
	var ate bool
	var x, y int
	var otherPlayer int

	ate = false
	otherPlayer = GetOtherTurn(childs)
	for x = -2; x <= 2; x += 2 {
		for y = -2; y <= 2; y += 2 {
			if (x >= 0 || y >= 0 || x < len(childs.board) || y < len(childs.board[0])) && childs.board[pawn.x+x][pawn.y+y] == childs.turn && childs.board[pawn.x+(x%2)][pawn.y+(y%2)] == otherPlayer {
				AddAteNumPlayer(childs, pawn.x+(x%2), pawn.y+(y%2))
				ate = true
			}
		}
	}
	//	Add check free three of ate player's pawn
}

func AddAteNumPlayer(childs GameData, x, y int) {
	childs.board[pawn.x+x][pawn.y+y] = 0
	if childs.turn == childs.facundo.whoiam {
		childs.facundo.ate++
	} else if childs.turn == childs.human.whoiam {
		childs.human.ate
	}
}

//	Check if pawn are aligned winning move (if win, return player ID else 0)
func CheckFiveAligned(childs GameData, pawn Pawn) {
	var x, y, s int
	var xfive, yfive int
	var xmax, ymax int
	var check, align int

	align = 0
	xmax = len(childs.board)
	ymax = len(childs.board[0])
	for s = 0; s < 4 && align < 5; s++ {
		align = 0
		for check = -4; check <= 4 && align < 5; check++ {
			x = pawn.x + check*(s/2+!s)
			y = pawn.y + check*(s%2-!s)
			if x < 0 || y < 0 || x >= xmax || y >= ymax {
				continue
			} else if childs.board[x][y] == childs.turn {
				align++
			} else {
				align = 0
			}
			if align == 5 {
				if childs.turn == childs.facundo.whoiam {
					AddFiveAligned(childs.facundo, x, y, s)
				} else if childs.turn == childs.human.whoiam {
					AddFiveAligned(childs.human, x, y, s)
				}
			}
		}
	}
}

//	Save aligned pawns
func AddFiveAligned(child Player, x, y, s int) {
	x -= 4 * (s/2 + !s)
	y -= 4 * (s%2 - !s)
	child.five_w = append(child.five_w, AlingnW{
		pos: PawnsInit(x, y),
		dir: s,
		win: false,
	})
	child.threef = true
}

//	Check time remaining for calcul
/*func CheckTimeLeft(remain, occur float64) float64 {
	remain -= occu
	if remain < 0.0 {
		remain = 0.0
	}
	return remain
}*/

//	Add pawn and new authorized place on board
func AddPawnOnBoard(childs GameData, pawn Pawn) {
	childs.board[pawn.x][pawn.y] = childs.turn
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if childs.board[pawn.x+x][pawn.y+y] == -4 {
				childs.board[pawn.x+x][pawn.y+y] = 0
			}
		}
	}
	if childs.move.x == -1 || childs.move.y == -1 {
		childs.move = pawn
	}
}

//	MinMax add couche
func MinMax(childs GameData, pawn NextPawns, link chan GameData) {
	AddPawnOnBoard(childs, pawn)
	//CheckVictory(childs)
	CheckEatPawn(childs, pawn)
	CheckFiveAligned(childs, pawn)
	//CheckFreeThree(childs)
	//CheckWinLoseProb(childs)
	childs.turn = GetOtherTurn(childs)
	childs.deep--
	link <- childs
}
