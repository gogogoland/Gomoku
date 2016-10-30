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
 * 		NOTHING
**/

func AlignWInit(x, y, dir int) AlignW {
	return AlignW{
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
