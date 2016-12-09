/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   PathFinding.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/03/17 19:24:03 by tbalea            #+#    #+#             */
/*   Updated: 2016/03/17 19:24:03 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

import (
	"container/heap"
	"fmt"
	"runtime"
	"time"
)

/**
 * Function of pathfinding
 *
 * TODO:
 * 		Maybe stop the A* when best move detect
 * 		Seek Goroutine
 *		Add timer
 * 		NOTHING
**/

//	Functions A*
//	*	Implementation of A* of MinMax
func Pathfinding(Mom GameData, deepness, IA int) NextPawns {
	var childNum int
	var childs GameData
	var link chan GameData
	var open *PrioQueue
	var childPawn []NextPawns

	//	Start timer
	timeStart := time.Now()

	//	Set number of proc
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	//	Init Heap
	open = InitHeapList(Mom, deepness)
	heap.Init(open)

	//	Make link between goroutine and original tree
	link = make(chan GameData)

	//	Run A* for MinMax
	for len(*open) > 0 {
		childs = heap.Pop(open).(GameData)
		childPawn = GetPossiblePlace(childs)
		childNum = len(childPawn)

		for i := 0; i < childNum; i++ {
			/*In MinMax : check value + rules*/
			go MinMax(childs, childPawn[i], link)
		}

		for i := 0; i < childNum; i++ {
			childs = <-link
			/*
			 *
			 */
			childs.prob = GetProbabilityByDeepness(childs.deep, childNum, 1)
			AddGameDataToHeapList(open, childs)
			if childs.human.winpot >= 1.0 || childs.facundo.winpot >= 1.0 || childs.deep == 0 {
				return GetOptimalPath(open, GetPossiblePlace(Mom), IA, timeStart)
			}
		}
	}
	return GetOptimalPath(open, GetPossiblePlace(Mom), IA, timeStart)
}

/*
 *
 */

/*
			//ADD close heaplist
			if childs.human.winpot >= 1.0 || childs.facundo.winpot >= 1.0 || childs.deep == 0 {
				childs.prob = GetProbabilityByDeepness(childs.deep, childNum, 1)
				if close != nil {
					AddGameDataToHeapList(close, childs)
				} else {
					close = InitGameDataToHeapList(close, childs, childs.deep)
					heap.Init(close)
				}
			} else {
				AddGameDataToHeapList(open, childs)
			}
			i++
		}
	}
	return GetOptimalPath(close, GetPossiblePlace(Mom), IA)
}
*/

//	Init priorityqueue
func InitGameDataToHeapList(list *PrioQueue, gameData GameData, deep int) *PrioQueue {
	list = InitHeapList(gameData, deep)
	heap.Init(list)
	return list
}

//	Add to priorityqueue
func AddGameDataToHeapList(list *PrioQueue, gameData GameData) {
	heap.Push(list, gameData)
	heap.Fix(list, len(*list)-1)
}

//	Get greater probability by low deepness
func GetProbabilityByDeepness(deep, probmax, probmin int) int {
	if deep != 0 {
		return probmax
	}
	return probmin
}

//	Get best move for current close list
func GetOptimalPath(close *PrioQueue, childPawn []NextPawns, IA int, timeStart Time) NextPawns {
	var FacundoMove NextPawns
	var childNum int
	var childs GameData

	FacundoMove = NextPawnsInit(-1, -1, 0.0, 0)
	childNum = len(childPawn)
	for len(*close) > 0 {
		childs = heap.Pop(close).(GameData)
		i := 0
		for childs.move.x != childPawn[i].pawn_p.x || childPawn[i].pawn_p.y != childs.move.x {
			i++
		}
		childPawn[i].winpot *= (float32)(childPawn[i].test_n)
		childPawn[i].test_n += childs.prob
		childPawn[i].winpot += (float32)(UseOfIA(childs, IA))
		childPawn[i].winpot = childPawn[i].winpot / (float32)(childPawn[i].test_n)
	}

	for i := 0; i < childNum; i++ {
		if FacundoMove.winpot < childPawn[i].winpot {
			FacundoMove = childPawn[i]
		}
	}

	timeSince := time.Since(timeStart)
	fmt.Println("Time required: ", timeSince)

	return FacundoMove
}

//	Use of different IA
func UseOfIA(childs GameData, IA int) float32 {
	if IA == 1 {
		return childs.facundo.winpot * (float32)(childs.prob)
	} else if IA == 0 {
		return childs.human.winpot * (float32)(childs.prob) * -1.0
	}
	return 0.0
}

//	Get all avaible places for current player
func GetPossiblePlace(gd GameData) []NextPawns {
	var np []NextPawns
	var curx, cury, i int
	var xmax, ymax int
	var np_size int
	var authPlayer int

	i = 0
	np_size = 0
	xmax = len(gd.board)
	ymax = len(gd.board[0])
	authPlayer = GetOtherTurn(gd) * -1
	for curx = 0; curx < xmax; curx++ {
		for cury = 0; cury < ymax; cury++ {
			if gd.board[curx][cury] != 0 || gd.board[curx][cury] == authPlayer {
				np_size++
			}
		}
	}

	gd.prob = np_size
	np = make([]NextPawns, np_size)
	for curx = 0; curx < xmax; curx++ {
		for cury = 0; cury < ymax; cury++ {
			if gd.board[curx][cury] != 0 || gd.board[curx][cury] == authPlayer {
				np[i] = NextPawnsInit(curx, cury, np_size, 0.0)
				i++
			}
		}
	}
	return np
}
