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
 * Functions for NextPawns structures
 *
 * TODO:
 * 		copy of GameData to Test
 * 		NOTHING
**/

func AlignPInit(x, y, dir int) AlignP {
	return AlignP{
		pos: PawnsInit(x, y),
		dir: dir,
	}
}

func NextPawnsInit(x, y, test_n int, winpot float32) NextPawns {
	return NextPawns{
		pawn_p: PawnsInit(x, y),
		winpot: winpot,
		test_n: test_n,
	}
}

func PawnsInit(x, y int) Pawns {
	return Pawns{
		x: x,
		y: y,
	}
}

func GetOtherTurn(gd GameData) int {
	var newTurn int

	if gd.turn == gd.human.whoiam {
		newTurn = gd.facundo.whoiam
	} else if gd.turn == gd.facundo.whoiam {
		newTurn = gd.human.whoiam
	}
	return newTurn
}

func CompareAlignP(ap1, ap2 AlignP) bool {
	return (ComparePawn(ap1.pos, ap2.pos) && ap1.dir == ap2.dir)
}

func ComparePawn(p1, p2 Pawn) bool {
	return (p1.x == p2.x && p1.y == p2.y)
}

func AddAlignP(lst []AlignP, new AlignP) {
	var cur, len int

	len = len(lst)
	for cur = 0; cur < len; cur++ {
		if CompareAlignP(lst[cur], new) {
			return
		}
	}
	lst = append(lst, new)
}

func CopyGameData(tocopy, theone GameData) {
	*theone = *tocopy
}
