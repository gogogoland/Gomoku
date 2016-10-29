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

/**
 * TODO:
 * 		replace calcul by kernel(GPU)
 * 		GetPossiblePlace to do (return list pawns and his size)
 *
 * 		NOTHING
**/

//	Functions A*
//	*	Implementation of A* of MinMax
func Pathfinding(Mom GameData, deepness, IA int) NextPawns {
	var childNum int
	var childs GameData
	var link chan GameData
	var open, close PrioQueue
	var childPawn *NextPawns

	//	Init Heap
	open = InitAddGameDataToHeapList(open, Mom, deepness)
	close = nil
	//	Make link between goroutine and original tree
	link = make(chan GameData)

	//	Run A* for MinMax
	for len(*open) > 0 {

		childs = heap.Pop(open)
		childPawn = GetPossiblePlace(cur)
		childNum = len(childPawn)

		for i := 0; i < childNum; i++ {
			/*In MinMax : check value + rules + add first move if no move save*/
			go MinMax(childs, childPawn[i], link)
		}

		for i := 0; i < childNum; i++ {
			childs = <-link
			if childs.human.winpot >= 1.0 || child.facundo.winpot >= 1.0 || !childs.deep {
				child.prob = GetProbalityByDeepness(childs.deep, childNum, 1)
				if close {
					AddGameDataToHeapList(close, childs)
				} else {
					close = InitGameDataToHeapList(close, childs, childs.deep)
				}
			} else {
				AddGameDataToHeapList(open, childs)
			}
			i++
		}
	}
	return GetOptimalPath(close, GetPossiblePlace(Mom), IA)
}

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
func GetProbabiltyByDeepness(deep, probmax, probmin int) int {
	if deep {
		return probmax
	}
	return probmin
}

//	Get best move for current close list
func GetOptimalPath(close PrioQueue, childPawn *NextPawns, IA int) NextPawns {
	var FacundoMove NextPawns
	var childNum int
	var childs GameData

	FacundoMove = nil
	childNum = len(childPawn)
	for len(*close) {
		childs = heap.Pop(close)
		i := 0
		for childs.move.x != childPawn[i].x || childPawn[i].y != childs.move.x {
			i++
		}
		childPawn[i].winpot *= childPawn[i].test_n
		childPawn[i].test_n += childs.prob
		childPawn[i].winpot += UseOfIA(childs, IA)
		childPawn[i].winpot = childPawn[i].winpot / childPawn[i].test_n
	}

	for i := 0; i < childNum; i++ {
		if FacundoMove.winpot < childPawn[i].winpot {
			FacundoMove = childPawn[i]
		}
	}

	return FacundoMove
}

//	Use of different IA
func UseOfIA(childs GameData, IA int) int {
	if IA {
		return childs.facundo.winpot * childs.prob
	} else if !IA {
		return childs.human.winpot * childs.prob * -1
	}
	return 0.0
}

//	*	MinMax
/*func MinMax(cur GameData, pawns Pawns, link chan GameData) {
	<-link
}*/
