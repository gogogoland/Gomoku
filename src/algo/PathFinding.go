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
	//"runtime"//link
	"time"
)

/**
 * Function of pathfinding
 *
 * TODO:
 * 		Stop the A* when best move detect (die hard at 5 deepness)
 * 		Seek Goroutine
 *		Check USE OF IA to defend and attack
 * 		NOTHING
**/

/**
 * 	"Pathfinding"
 *
 *	Start timer and reset Mome first move
 * 	Implementation of A* using MinMax method for euristic
 * 	(each comment ending with "//link" keyword are for channels (see goroutine))
 * 	(each comment ending with "//direct" keyword are for stopping A*
 * 		in first seen Win or Lose)
**/
func (Mom *GameData) Pathfinding(deepness, IA int) {
	var maxMove, childNum int
	var childs GameData
	//var link chan GameData//link
	var open, clos *PrioQueue
	var childPawn []NextPawns

	timeStart := time.Now()
	Mom.move = PawnsInit(-1, -1)
	////Set number of proc//link
	//runtime.GOMAXPROCS(runtime.NumCPU())//link
	//fmt.Println(runtime.NumGoroutine())//link
	////Make link between goroutine and original tree//link
	//link = make(chan GameData)//link

	open = Mom.InitHeapList(deepness)
	heap.Init(open)
	clos = Mom.InitHeapList(-1)
	heap.Init(clos)

	for len(*open) > 0 {
		childs = heap.Pop(open).(GameData)
		childs.turn = childs.GetOtherTurn()
		childPawn = childs.GetPossiblePlace()
		childNum = len(childPawn)
		maxMove = childNum
		//for i := 0; i < childNum; i++ {//link
		//	go MinMax(childs, childPawn[i], link)//link
		//}//link
		for i := 0; i < childNum; {
			//childo, lessMove := <-link//link
			childo, lessMove := MinMax(childs, childPawn[i])
			if lessMove > 0 {
				childNum -= lessMove
				maxMove -= lessMove
				childPawn[i] = childPawn[childNum]
				childPawn = childPawn[:childNum]
				continue
			}
			i++
			childo.prob = GetProbabilityByDeepness(childo.deep, childNum, 1)
			//if childs.human.winpot >= 1.0 || childo.facundo.winpot >= 1.0 {//direct
			//	childo.AddGameDataToHeapList(clos)//direct
			//	Mom.GetOptimalPath(clos, IA, timeStart)//direct
			//	return//direct
			//} else if childo.deep != 0 {//direct
			if childo.deep != 0 && childo.facundo.winpot < 1.0 && childo.human.winpot < 1.0 {
				childo.AddGameDataToHeapList(open)
			} else {
				childo.AddGameDataToHeapList(clos)
			}
		}
	}

	Mom.GetOptimalPath(clos, IA, timeStart)
}

/**
 * 	"Get Probability By Deepness"
 *
 * 	Get best probability if it's not at maximum deepness
**/
func GetProbabilityByDeepness(deep, probmax, probmin int) int {
	if deep != 0 {
		return probmax
	}
	return probmin
}

/**
 * 	"Get Optimal Path"
 *
 * 	Check each tested mpath for each first possible move
 * 	Take the best path (if same, take one randomly) and process the first move
 * 	Print AI duration
**/
func (Mom *GameData) GetOptimalPath(queu *PrioQueue, IA int, timeStart time.Time) {
	var FacundoMove NextPawns
	var childPawn []NextPawns
	var childs GameData
	var childPawnLen int
	var i, randval int

	childPawn = Mom.GetPossiblePlace()
	FacundoMove = NextPawnsInit(-1, -1, 0, -1.0, false)
	childPawnLen = len(childPawn)
	for len(*queu) > 0 {
		childs = heap.Pop(queu).(GameData)
		i = 0
		for i < childPawnLen && (childs.move.x != childPawn[i].pawn_p.x || childPawn[i].pawn_p.y != childs.move.y) {
			i++
		}
		if i < childPawnLen {
			childPawn[i].winpot *= (float32)(childPawn[i].test_n)
			childPawn[i].test_n += childs.prob
			childPawn[i].winpot += (float32)(childs.UseOfIA(IA))
			childPawn[i].winpot = childPawn[i].winpot / (float32)(childPawn[i].test_n)
			childPawn[i].tested = true
		}
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i = 0; i < childPawnLen; i++ {
		if childPawn[i].tested == false {
			continue
		}
		//fmt.Println("FacundoMove :", FacundoMove, "ChildPawn :", childPawn[i])
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

	Mom.turn = Mom.GetOtherTurn()
	Mom.TurnProcess(FacundoMove.pawn_p)
	Mom.turn = Mom.GetOtherTurn()

	timeSince := time.Since(timeStart)
	fmt.Println("Time required: ", timeSince)

	//fmt.Println("Facundo Data :", Mom.facundo)
	//fmt.Println("Human Data :", Mom.human)
}

/**
 * 	"Use Of AI"
 *
 * 	Use different AI :
 * 	King of the Hill (defense move)
 * 	Spear of Brokens (attack move)
 * 	Facundo's mighty (mixt between defense and attack)
**/
func (childs *GameData) UseOfIA(IA int) float32 {
	return childs.human.winpot * (float32)(childs.prob) * -1.0
	//return 0 - childs.human.winpot*(float32)(childs.prob)
	return childs.facundo.winpot*(float32)(childs.prob) - childs.human.winpot*(float32)(childs.prob)
	if IA == 1 {
		return childs.facundo.winpot * (float32)(childs.prob)
	} else if IA == 0 {
		return childs.human.winpot * (float32)(childs.prob) * -1.0
	}
	return 0.0
}

/**
 * 	"Get Available Places"
 *
 *	Get number of potential move for GamdeData.prob
 * 	Return an array of NextPawns for all box close to a pawn
**/
func (gd *GameData) GetPossiblePlace() []NextPawns {
	var np []NextPawns
	var curx, cury, i int
	var np_size int

	i = 0
	np_size = 0
	for curx = 0; curx < gd.maxx; curx++ {
		for cury = 0; cury < gd.maxy; cury++ {
			if gd.board[curx][cury] == 0 || gd.board[curx][cury] == -1*gd.turn {
				np_size++
			}
		}
	}

	gd.prob = np_size
	np = make([]NextPawns, np_size)
	for curx = 0; curx < gd.maxx; curx++ {
		for cury = 0; cury < gd.maxy; cury++ {
			if gd.board[curx][cury] == 0 || gd.board[curx][cury] == -1*gd.turn {
				np[i] = NextPawnsInit(curx, cury, np_size, 0.0, false)
				i++
			}
		}
	}
	return np
}
