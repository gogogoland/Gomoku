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
 * 		Add direct win detect (use only the winning move by this path)
 *
 * 		Stop the A* when best move detect (die hard at 5 deepness)
 * 		Take the first best move from open list
 * 		Seek Goroutine
 *
 *		Check USE OF IA to defend and attack
 * 		NOTHING
**/

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
	Mom.GetHuman().pawn_p = PawnsInit(-1, -1)
	Mom.GetFacundo().pawn_p = PawnsInit(-1, -1)
	Mom.SetIA(IA)
	//link////Set number of proc
	//link//runtime.GOMAXPROCS(runtime.NumCPU())
	//link//fmt.Println(runtime.NumGoroutine())
	//link////Make link between goroutine and original tree
	//link//link = make(chan GameData)

	open = Mom.InitHeapList(2 * 10)
	heap.Init(open)
	clos = nil
	hope = nil

	for len(*open) > 0 && time.Since(timeStart).Nanoseconds()/1000000 < 450 {
		childs = heap.Pop(open).(GameData)
		childs.turn = childs.GetOtherTurn()
		childPawn = childs.GetPossiblePlace()
		childNum = len(childPawn)
		//link//for i := 0; i < childNum; i++ {
		//link//	go MinMax(childs, childPawn[i], link)
		//link//}
		//LINK TO US FOR EACH LAYER
		for i := 0; i < childNum && time.Since(timeStart).Nanoseconds()/1000000 < 450; {
			//link//childo, lessMove := <-link
			childo, less = MinMax(childs, childPawn[i])
			if less == 1 {
				childPawn[i] = childPawn[childNum-1]

				childPawn = childPawn[:childNum-1]
				childNum -= 1
				continue
			}
			// <- HERE THE BEGINNING OF THE LINK (new function)
			if childo.deep != 0 && childo.facundo.winpot < 1.0 && childo.human.winpot < 1.0 {
				childo.turn = childo.GetOtherTurn()
				childoPawn = childs.GetPossiblePlace()
				childoNum = len(childoPawn)
				for j := 0; j < childoNum; {
					childt, less = MinMax(childo, childoPawn[j])
					if less == 1 {
						childoPawn[j] = childoPawn[childoNum-1]
						childoPawn = childoPawn[:childoNum-1]
						childoNum -= 1
						continue
					} else if childt.GetHuman().GetWhoAmI() == childt.Gain() {
						if clos == nil {
							clos = childt.InitHeapList(childt.deep)
							heap.Init(clos)
						} else {
							childt.AddGameDataToHeapList(clos)
						}
						childoNum = 0
						break
					} else if hope == nil {
						hope = childt.InitHeapList(childt.deep)
						heap.Init(hope)
					} else {
						childt.AddGameDataToHeapList(hope)
					}
					j++
				}
				for j := 0; j < childoNum && time.Since(timeStart).Nanoseconds()/1000000 < 450; j++ {
					childt = heap.Pop(hope).(GameData)
					if childo.deep != 0 && childo.facundo.winpot < 1.0 {
						childt.AddGameDataToHeapList(open)
					} else if clos == nil {
						clos = childt.InitHeapList(childt.deep)
						heap.Init(clos)
					} else {
						childt.AddGameDataToHeapList(clos)
					}
				}
				hope = nil
			} else if clos == nil {
				clos = childo.InitHeapList(childo.deep)
				heap.Init(clos)
			} else {
				childo.AddGameDataToHeapList(clos)
			}
			// <- END HERE
			i++
		}
	}
	Mom.GetOptimalPath(clos, open, timeStart, vs)
}

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
	var cur float32

	childPawnLen = len(childPawn)
	for queue != nil && len(*queue) > 0 {
		childs = heap.Pop(queue).(GameData)
		i = 0
		for i < childPawnLen && (childs.move.x != childPawn[i].pawn_p.x || childPawn[i].pawn_p.y != childs.move.y) {
			i++
		}
		if i < childPawnLen {
			cur = childs.UseOfIA(turn)
			if childPawn[i].winpot > cur {
				childPawn[i].winpot = cur
			}
			childPawn[i].tested = true
		}
	}
	return childPawn
}

/**
 * 	"Get Available Places"
 *
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

	np = make([]NextPawns, np_size)
	for curx = 0; curx < gd.maxx; curx++ {
		for cury = 0; cury < gd.maxy; cury++ {
			if gd.board[curx][cury] == 0 || gd.board[curx][cury] == -1*gd.turn {
				np[i] = NextPawnsInit(curx, cury, np_size, 1.0, false)
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
	return BoolToFloat32(childs.GetPlayer(turn).atenum == 5) + BoolToFloat32(len(childs.GetPlayer(turn).five_w) > 0) - childs.GetOtherPlayer(turn).winpot
	if childs.ai == 0 {
		return childs.GetPlayer(turn).winpot
	} else if childs.ai == 1 {
		return childs.GetOtherPlayer(turn).winpot * -1.0
	} else if childs.ai == 2 {
		return BoolToFloat32(childs.GetPlayer(turn).atenum == 5) + BoolToFloat32(len(childs.GetPlayer(turn).five_w) > 0) - childs.GetOtherPlayer(turn).winpot
	}
	return 0.0
}
