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
	"math/rand"
	"runtime"
	"time"
)

/**
 * Function of pathfinding
 *
 * TODO:
 *		Use of address for GameData
 * 		Stop the A* when best move detect
 * 		Seek Goroutine
 *		Check Human data inside Pathfinding
 *		Check USE OF IA to defend and attack
 * 		NOTHING
**/

//	Functions A*
//	*	Implementation of A* of MinMax
func (Mom *GameData) Pathfinding(deepness, IA int) /*, Pawns, [][]int*/ {
	var childNum int
	var childs GameData
	//var link chan GameData//GOLD
	var open, clos *PrioQueue
	var childPawn []NextPawns

	//	Start timer
	timeStart := time.Now()

	//	Set number of proc
	//runtime.GOMAXPROCS(runtime.NumCPU())//GOLD
	fmt.Println(runtime.NumGoroutine())

	//	Reinit Mom Data for current pathfinding
	Mom.move = PawnsInit(-1, -1)

	//	Init Heap
	open = Mom.InitHeapList(deepness)
	heap.Init(open)
	clos = Mom.InitHeapList(-1)
	heap.Init(clos)

	//	Make link between goroutine and original tree
	//link = make(chan GameData)//GOLD

	//	Run A* for MinMax
	for len(*open) > 0 {
		childs = heap.Pop(open).(GameData)
		childPawn = childs.GetPossiblePlace()
		childNum = len(childPawn)

		for i := 0; i < childNum; i++ {
			/*In MinMax : check value + rules*/
			//go MinMax(childs, childPawn[i], link)//GOLD
			childo := MinMax(childs, childPawn[i])
			//}//GOLD
			//print("END TURN.\n")//GOLDBUG

			//for i := 0; i < childNum; i++ {//GOLD
			//	Get childs from goroutine minmax
			//childs = <-link//GOLD
			//print("END LINK.\n")//GOLDBUG
			/*
			 * 	!Seek Goroutine!
			 */

			//	Get probability of outcome
			//childs.prob = GetProbabilityByDeepness(childs.deep, childNum, 1)//GOLD
			childo.prob = GetProbabilityByDeepness(childo.deep, childNum, 1)
			//	Add to open's queu current child
			//AddGameDataToHeapList(open, childs)//GOLD
			if childo.deep != 0 && childo.facundo.winpot < 1.0 && childo.human.winpot < 1.0 {
				childo.AddGameDataToHeapList(open)
			} else {
				childo.AddGameDataToHeapList(clos)
			}
			//	In case of victory, go after it
			/*
			 * 	!Get Best Path!
			 * 	!OR ADD close queu!
			 */
			if /* childs.human.winpot >= 1.0 || */ childo.facundo.winpot >= 1.0 /* || childs.deep == 0 */ {
				/*return*/ Mom.GetOptimalPath(clos, IA, timeStart)
				return
			}
		}
	}
	/*return*/ Mom.GetOptimalPath(clos, IA, timeStart)
}

//	Get greater probability by low deepness
func GetProbabilityByDeepness(deep, probmax, probmin int) int {
	if deep != 0 {
		return probmax
	}
	return probmin
}

//	Get best move from current queu
func (Mom *GameData) GetOptimalPath(queu *PrioQueue, IA int, timeStart time.Time) /*, Pawns, [][]int*/ {
	var FacundoMove NextPawns
	var childPawn []NextPawns
	var childs GameData
	var childPawnLen int
	var i, randval int

	childPawn = Mom.GetPossiblePlace()
	FacundoMove = NextPawnsInit(-1, -1, 0.0, 0)
	childPawnLen = len(childPawn)
	//	Get all childs from queu
	for len(*queu) > 0 {
		childs = heap.Pop(queu).(GameData)
		i = 0
		//	Add win potential for nextpawns from current child
		for i < childPawnLen && (childs.move.x != childPawn[i].pawn_p.x || childPawn[i].pawn_p.y != childs.move.x) {
			i++
		}
		if i < childPawnLen {
			childPawn[i].winpot *= (float32)(childPawn[i].test_n)
			childPawn[i].test_n += childs.prob
			childPawn[i].winpot += (float32)(childs.UseOfIA(IA))
			childPawn[i].winpot = childPawn[i].winpot / (float32)(childPawn[i].test_n)
		}
	}

	//	Get best winpot for Facundo
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i = 0; i < childPawnLen; i++ {
		fmt.Println("FacundoMove :", FacundoMove, "ChildPawn :", childPawn[i])
		if FacundoMove.winpot < childPawn[i].winpot {
			FacundoMove = childPawn[i].Copy()
			randval = 1
		} else if FacundoMove.winpot == childPawn[i].winpot {
			randval += 1
			if random.Intn(randval) == randval-1 {
				FacundoMove = childPawn[i].Copy()
			}
		}
	}

	fmt.Println("FINAL TURN PROCESS")
	//	Effectivly process Facundo turn
	Mom.TurnProcess(FacundoMove.pawn_p)

	// Print duration from beginning of AI
	timeSince := time.Since(timeStart)
	fmt.Println("Time required: ", timeSince)
	fmt.Println("Facundo Data :", Mom.facundo)
	fmt.Println("Human Data :", Mom.human)

	/*return FacundoMove.pawn_p, Mom.Board*/
}

//	Use of different IA
func (childs *GameData) UseOfIA(IA int) float32 {
	return childs.facundo.winpot*(float32)(childs.prob) - childs.human.winpot*(float32)(childs.prob)
	if IA == 1 {
		return childs.facundo.winpot * (float32)(childs.prob)
	} else if IA == 0 {
		return childs.human.winpot * (float32)(childs.prob) * -1.0
	}
	return 0.0
}

//	Get all avaible places for current player
func (gd *GameData) GetPossiblePlace() []NextPawns {
	var np []NextPawns
	var curx, cury, i int
	var np_size int

	i = 0
	np_size = 0
	//	Get number of childrens possible move
	for curx = 0; curx < gd.maxx; curx++ {
		for cury = 0; cury < gd.maxy; cury++ {
			if gd.board[curx][cury] == 0 || gd.board[curx][cury] == -1*gd.turn {
				np_size++
			}
		}
	}

	gd.prob = np_size
	//	Set potential next pawns
	np = make([]NextPawns, np_size)
	for curx = 0; curx < gd.maxx; curx++ {
		for cury = 0; cury < gd.maxy; cury++ {
			if gd.board[curx][cury] == 0 || gd.board[curx][cury] == -1*gd.turn {
				np[i] = NextPawnsInit(curx, cury, np_size, 0.0)
				i++
			}
		}
	}
	return np
}
