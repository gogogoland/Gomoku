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

import "container/heap"

/**
 * Function intern of GameData heaplist
 *
 * TODO:
 * 		Test comparaison
 * 		NOTHING
**/

// Intern functions for HeapList
//	*	Add
func (heapList *PrioQueue) Push(toAdd interface{}) {
	*heapList = append(*heapList, GameData{
		facundo: toAdd.(GameData).facundo,
		human:   toAdd.(GameData).human,
		board:   toAdd.(GameData).board,
		maxx:    toAdd.(GameData).maxx,
		maxy:    toAdd.(GameData).maxy,
		deep:    toAdd.(GameData).deep,
		move:    toAdd.(GameData).move,
		turn:    toAdd.(GameData).turn,
		whowin:  toAdd.(GameData).whowin,
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
func (heapList PrioQueue) Swap(i, j int) {
	heapList[i], heapList[j] = heapList[j], heapList[i]
}

//	*	Check order by deepness and scoring value
func (heapList PrioQueue) Less(i, j int) bool {
	if heapList[i].UseOfIA(heapList[i].turn) == heapList[j].UseOfIA(heapList[j].turn) {
		return heapList[i].deep > heapList[j].deep
	}
	return heapList[i].UseOfIA(heapList[i].turn) > heapList[j].UseOfIA(heapList[j].turn)
}

//	*	Get len
func (heapList PrioQueue) Len() int {
	return len(heapList)
}

// GameData dependent functions
//	*	Init HeapList from GameData
func (gameData GameData) InitHeapList(deepness int) *PrioQueue {
	return &PrioQueue{gameData.Deep(deepness)}
}

//	*	Add to Heap List
func (gameData GameData) AddGameDataToHeapList(list *PrioQueue) {
	heap.Push(list, gameData)
	heap.Fix(list, len(*list)-1)
}
