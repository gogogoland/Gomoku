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
	//link//"runtime"
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
 * 	(each comments beginning with "//link" keyword are for channels (see goroutine))
 * 	(each comments beginning with "//direct" keyword are for stopping A*
 * 		in first seen Win or Lose)
**/
/*func (Mom *GameData) Pathfinding(deepness, IA int) {
	var maxMove, childNum int
	var childs GameData
	//link//var link chan GameData
	var open, clos *PrioQueue
	var childPawn []NextPawns

	timeStart := time.Now()
	Mom.move = PawnsInit(-1, -1)
	Mom.SetIA(IA)
	//link////Set number of proc
	//link//runtime.GOMAXPROCS(runtime.NumCPU())
	//link//fmt.Println(runtime.NumGoroutine())
	//link////Make link between goroutine and original tree
	//link//link = make(chan GameData)

	open = Mom.InitHeapList(deepness)
	heap.Init(open)
	clos = Mom.InitHeapList(-1)
	heap.Init(clos)

	for len(*open) > 0 && time.Since(timeStart).Nanoseconds()/1000000 < 500 {
		childs = heap.Pop(open).(GameData)
		childs.turn = childs.GetOtherTurn()
		childPawn = childs.GetPossiblePlace()
		childNum = len(childPawn)
		maxMove = childNum
		//link//for i := 0; i < childNum; i++ {
		//link//	go MinMax(childs, childPawn[i], link)
		//link//}
		for i := 0; i < childNum; {
			//link//childo, lessMove := <-link
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
			//direct//if childs.human.winpot >= 1.0 || childo.facundo.winpot >= 1.0 {
			//direct//	childo.AddGameDataToHeapList(clos)
			//direct//	Mom.GetOptimalPath(clos, IA, timeStart)
			//direct//	return
			//direct//} else if childo.deep != 0 {
			if childo.deep != 0 && childo.facundo.winpot < 1.0 && childo.human.winpot < 1.0 {
				childo.AddGameDataToHeapList(open)
			} else {
				childo.AddGameDataToHeapList(clos)
			}
		}
	}

	Mom.GetOptimalPath(clos, open, timeStart)
}*/

/**
 * 	"Pathfinding"
 *
 *	Start timer and reset Mome first move
 * 	Implementation of A* using MinMax method for euristic
 * 	Check even number of layer
 * 	(each comments beginning with "//link" keyword are for channels (see goroutine))
 * 	(each comments beginning with "//direct" keyword are for stopping A*
 * 		in first seen Win or Lose)
**/
func (Mom *GameData) Pathfinding(deepness, IA int, vs bool) {
	var childNum, childoNum, less int
	var childs, childo, childt GameData
	//link//var link chan GameData
	var open, clos, hope *PrioQueue
	var childPawn []NextPawns
	var childoPawn []NextPawns

	timeStart := time.Now()
	Mom.move = PawnsInit(-1, -1)
	Mom.SetIA(IA)
	//link////Set number of proc
	//link//runtime.GOMAXPROCS(runtime.NumCPU())
	//link//fmt.Println(runtime.NumGoroutine())
	//link////Make link between goroutine and original tree
	//link//link = make(chan GameData)

	open = Mom.InitHeapList(10)
	heap.Init(open)
	clos = Mom.InitHeapList(-1)
	heap.Init(clos)

	for len(*open) > 0 && time.Since(timeStart).Nanoseconds()/1000000 < 500 {
		childs = heap.Pop(open).(GameData)
		childs.turn = childs.GetOtherTurn()
		childPawn = childs.GetPossiblePlace()
		childNum = len(childPawn)
		//link//for i := 0; i < childNum; i++ {
		//link//	go MinMax(childs, childPawn[i], link)
		//link//}
		//LINK TO US FOR EACH LAYER
		for i := 0; i < childNum; {
			//link//childo, lessMove := <-link
			childo, less = MinMax(childs, childPawn[i])
			/*if childPawn.RemoveInvalidMove(less, i) {
				childNum -= less
				continue
			}*/
			if less == 1 {
				childPawn[i] = childPawn[childNum-1]

				childPawn = childPawn[:childNum-1]
				childNum -= 1
				continue
			}
			// <- HERE THE BEGINNING OF THE LINK (new function)
			if childo.deep != 0 && childo.facundo.winpot < 1.0 && childo.human.winpot < 1.0 {
				hope = childo.InitHeapList(-1)
				heap.Init(hope)
				childo.turn = childo.GetOtherTurn()
				childoPawn = childs.GetPossiblePlace()
				childoNum = len(childoPawn)
				for j := 0; j < childoNum; {
					childt, less = MinMax(childo, childoPawn[j])
					/*if childoPawn.RemoveInvalidMove(less, i) {
					childoNum -= less
					continue*/
					if less == 1 {
						childoPawn[j] = childoPawn[childoNum-1]
						childoPawn = childoPawn[:childoNum-1]
						childoNum -= 1
						continue
					} else if len(childt.GetHuman().GetFive_W()) > 0 { //.GetWhoAmI() == childt.Gain() {
						childoNum = 0
						//THIS PATH IS NOT GOOD
						break
					} else {
						childo.AddGameDataToHeapList(hope)
					}
					j++
				}
				for j := 0; j < childoNum; j++ {
					childt = heap.Pop(hope).(GameData)
					if childo.deep != 0 && childo.facundo.winpot < 1.0 {
						childt.AddGameDataToHeapList(open)
					} else {
						childt.AddGameDataToHeapList(clos)
					}
				}
				hope = nil
			} else {
				childo.AddGameDataToHeapList(clos)
			}
			// <- END HERE
			i++
		}
	}
	Mom.GetOptimalPath(clos, open, timeStart, vs)
}

/*func (child *NextPawns) RemoveInvalidMove(less, cur int) int {
	var childLen int

	if less == 0 {
		return 0
	}
	childLen = len(*child)
	child[cur] = child[childLen]
	child = child[:childLen]
	return 1
}*/

/**
 * 	"Get Optimal Path"
 *
 * 	Check each tested mpath for each first possible move
 * 	Take the best path (if same, take one randomly) and process the first move
 * 	Print AI duration
**/
func (Mom *GameData) GetOptimalPath(clos, open *PrioQueue, timeStart time.Time, vs bool) {
	var FacundoMove NextPawns
	var childPawn []NextPawns
	var childPawnLen int
	var i, randval int
	var turn int

	childPawn = Mom.GetPossiblePlace()
	FacundoMove = NextPawnsInit(-1, -1, 0, -1.0, false)
	childPawnLen = len(childPawn)
	turn = Mom.GetFacundo().GetWhoAmI()

	childPawn = GetHeapPath(open, childPawn, turn)
	childPawn = GetHeapPath(clos, childPawn, turn)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i = 0; i < childPawnLen; i++ {
		fmt.Println("FacundoMove :", FacundoMove, "ChildPawn :", childPawn[i])
		if childPawn[i].tested == false {
			continue
		}
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

	if vs == false {
		Mom.turn = Mom.GetOtherTurn()
		Mom.TurnProcess(FacundoMove.pawn_p)
		Mom.turn = Mom.GetOtherTurn()
		Mom.round++
		fmt.Println("Round turn: ", Mom.GetRound())
	} else {
		fmt.Println("BEST MOVE :", FacundoMove.pawn_p)
	}

	timeSince := time.Since(timeStart)
	fmt.Println("Time required: ", timeSince)
}

/**
 * 	"Get Heap Path"
 *
 * 	Get all move from Priority Queue
**/
func GetHeapPath(queue *PrioQueue, childPawn []NextPawns, turn int) []NextPawns {
	var childs GameData
	var childPawnLen int
	var i int

	childPawnLen = len(childPawn)
	for len(*queue) > 0 {
		childs = heap.Pop(queue).(GameData)
		i = 0
		for i < childPawnLen && (childs.move.x != childPawn[i].pawn_p.x || childPawn[i].pawn_p.y != childs.move.y) {
			i++
		}
		if i < childPawnLen {
			childPawn[i].winpot *= (float32)(childPawn[i].test_n)
			childPawn[i].test_n += childs.prob
			childPawn[i].winpot += (float32)(childs.UseOfIA(turn))
			childPawn[i].winpot = childPawn[i].winpot / (float32)(childPawn[i].test_n)
			childPawn[i].tested = true
		}
	}
	return childPawn
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

/**
 * 	"Use Of AI"
 *
 * 	Use different AI :
 * 	King of the Hill (defense move)
 * 	Spear of Brokens (attack move)
 * 	Facundo's mighty (mixt between defense and attack)
**/
func (childs *GameData) UseOfIA(turn int) float32 {
	//return 0 - childs.human.winpot*(float32)(childs.prob)
	//if childs.GetPlayer(childs.turn)
	/**
	 * 	ADD TURN CONDITION
	 * 	OR RETURN CURRENT TURN PROB
	 * 	AND CHECK USEFULNESS OF .prob
	**/
	return childs.GetOtherPlayer(turn).winpot * (float32)(childs.prob) * -1.0
	return childs.GetPlayer(turn).winpot*(float32)(childs.prob) - childs.GetOtherPlayer(turn).winpot*(float32)(childs.prob)
	//return childs.human.winpot * (float32)(childs.prob) * -1.0
	//return childs.facundo.winpot*(float32)(childs.prob) - childs.human.winpot*(float32)(childs.prob)
	if childs.ai == 1 {
		return childs.GetPlayer(turn).winpot * (float32)(childs.prob)
		//return childs.facundo.winpot * (float32)(childs.prob)
	} else if childs.ai == 0 {
		return childs.GetOtherPlayer(turn).winpot * (float32)(childs.prob) * -1.0
		//return childs.human.winpot * (float32)(childs.prob) * -1.0
	}
	return 0.0
}
