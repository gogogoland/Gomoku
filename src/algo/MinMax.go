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
func CheckEatPawn(board [][]int, max Pawns, player Player) int {
	var eat int
	var check int
	var dir int
	var x, y int
	var xe, ye int

	for eat, dir = 0, 0; dir < 4 && align < 5; dir++ {
		for check = -1; check < 2; check += 2 {
			if (dir / 2) && (player.pos.y+check < 0) {
				check = 1
			} else if (dir / 2) && (player.pos.y+check >= max.y) {
				break
			} else if (!dir) && (player.pos.y-check >= max.y) {
				check = 1
			} else if (!dir) && (player.pos.y-check < 0) {
				break
			}
			if ((dir % 2) || (!dir)) && (player.pos.x+check < 0) {
				check = 1
			} else if ((dir % 2) || (!dir)) && (player.pos.x+check >= max.x) {
				break
			}
			y = player.pos.y + ((dir / 2) * 2 * check) - ((!dir) * 2 * check)
			x = player.pos.x + ((dir % 2) * 2 * check) + ((!dir) * 2 * check)
			ye = player.pos.y + ((dir / 2) * check) - ((!dir) * check)
			xe = player.pos.x + ((dir % 2) * check) + ((!dir) * check)
			if board[y][x] == player.witch && board[ye][xe] == -player.witch {
				player.eat++
				board[ye][xe] = 0
			}
		}
	}
	return ((player.eat / 5) * player.witch)
}

//	Check if pawn are aligned winning move (if win, return player ID else 0)
func CheckFiveAligned(board [][]int, max, pos Pawns, player Player) int {
	var align, dir int
	var x, y int
	var check int

	for aling, dir = 0, 0; dir < 4 && align < 5; dir++ {
		for check = -4; check < 5 && align < 5; check++ {
			if (dir / 2) && (player.pos.y+check < 0) {
				check = -player.pos.y
			} else if (dir / 2) && (player.pos.y+check >= max.y) {
				break
			} else if (!dir) && (player.pos.y-check >= max.y) {
				check = player.pos.y - max.y
			} else if (!dir) && (player.pos.y-check < 0) {
				break
			}
			if ((dir % 2) || (!dir)) && (player.pos.x+check < 0) {
				check = -player.pos.x
			} else if ((dir % 2) || (!dir)) && (player.pos.x+check >= max.x) {
				break
			}
			y = player.pos.y + ((dir / 2) * check) - ((!dir) * check)
			x = player.pos.x + ((dir % 2) * check) + ((!dir) * check)
			align++
			if board[y][x] != player.witch {
				align = 0
			}
		}
	}
	return ((align / 5) * player.witch * player.win)
}

//	Check time remaining for calcul
func CheckTimeLeft(remain, occur float64) float64 {
	remain -= occu
	if remain < 0.0 {
		remain = 0.0
	}
	return remain
}

//	check changement of this turn
func EndTurn(board [][]int, max Pawns, human, facundo Player) int {
	var victory int

	victory = CheckFiveAligned(board, max, human)
	if !victory {
		victory = CheckFiveAligned(board, max, facundo)
	}
	if !victory {
		victory = CheckEatPawn(board, max, human, facundo)
	}
	CheckAuthorizedMove(board, max, pos)
	return victory
}

//	MinMax add couche
func MinMax(childs GameData, pawn NextPawns, link chan GameData) {
	childs.board[pawn.pawn_p.x][pawn.pawn_p.y] = childs.deep % 2
	childs.deep--
	CheckEatPawn(childs, pawn)
	CheckFiveAligned(childs)
	CheckFreeThree(childs)
	link <- childs
}
