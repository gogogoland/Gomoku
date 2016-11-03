/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   HeapFunctions.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/03/22 18:42:35 by tbalea            #+#    #+#             */
/*   Updated: 2016/03/22 18:42:35 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

/**
 * Function intern of GameData heaplist
 *
 * TODO:
 * 		NOTHING
**/

// Intern functions for HeapList
//	*	Add
func (heapList *PrioQueue) Push(toAdd interface{}) {
	*heapList = append(*heapList, GamesData{
		facundo: toAdd.(GamesData).facundo,
		human:   toAdd.(GamesData).human,
		board:   toAdd.(GamesData).board,
		deep:    toAdd.(GamesData).deep,
		move:    toAdd.(GamesData).move,
		prob:    toAdd.(GamesData).prob,
	})
}

//	*	Get and Delete
func (heapList *PrioQueue) Pop() interface{} {
	old := *heapList
	n := len(old)
	poped := old[n-1]
	*heapList = old[0 : n-1]
	return poped
}

//	*	Swap data
func (heapList *PrioQueue) Swap(i, j int) {
	heapList[i], heapList[j] = heapList[j], heapList[i]
}

//	*	Check order by deepness and scoring value
func (heapList *PrioQueue) Less(i, j int) bool {
	return heapList[i].prob*heapList[i].deep > heapList[j].prob*heapList[j].deep
}

//	*	Get len
func (heapList *PrioQueue) Len() int {
	return len(heapList)
}

// Extern functions
//	*	Init HeapList
func InitHeapList(gameData GameData, deepness int) *PrioQueue {
	var five_f, five_h []AlignW
	var newBoard [][]int
	var max int

	five_f = make([]AlignW, len(gameData.facundo.five_w))
	copy(five_f, gameData.facundo.five_w)
	five_h = make([]AlignW, len(gameData.human.five_w))
	copy(five_h, gameData.human.five_w)
	max = len(gameData.board)
	newBoard = make([][]int, max)
	for i := 0; i < max; i++ {
		newBoard[i] = make([]int, len(gameData.board[i]))
		copy(newBoard[i], gameData.board[i])
	}

	return &PrioQueue{
		GamesData{
			facundo: Player{
				atenum: gameData.facundo.atenum,
				whoiam: gameData.facundo.whoiam,
				pawn_p: PawnsInit(gameData.facundo.pawn_p.x, gameData.facundo.pawn_p.y),
				five_w: five_f,
				threef: gameData.facundo.threef,
				haswin: gameData.facundo.haswin,
				winpot: gameData.facundo.winpot,
			},
			human: Player{
				atenum: gameData.human.atenum,
				whoiam: gameData.human.whoiam,
				pawn_p: PawnsInit(gameData.human.pawn_p.x, gameData.human.pawn_p.y),
				five_w: five_h,
				threef: gameData.human.threef,
				haswin: gameData.human.haswin,
				winpot: gameData.human.winpot,
			},
			board: newBoard,
			deep:  deepness,
			move:  PawnsInit(gameData.move.x, gameData.move.y),
			prob:  gameData.prob,
		},
	}
	return nil
}
