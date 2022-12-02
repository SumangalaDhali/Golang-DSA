// select one element at a time from the left of the collection
//Insert element at a proper position
//After insertion every element to its left will be sorted

package main

import "fmt"

func insertionsort(List []int) []int {
	n := len(List)
	for i := 1; i < n; i++ {
		cvalue := List[i] // used to compare and swap
		position := i
		for position > 0 && List[position] > cvalue {
			List[position] = List[position-1]
			position = position - 1
		}
		List[position] = cvalue // to insert element in given position
	}
	return List
}

func main() {
	L := []int{3, 5, 8, 9, 6, 2}
	I := insertionsort(L)
	fmt.Println(I)

}
