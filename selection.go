package main

import "fmt"

func selectionsort(List []int) []int {
	n := len(List)
	for i := 0; i < n-1; i++ {
		position := i
		for j := i + 1; j < n; j++ {
			if List[j] < List[position] {
				position = j
			}
			temp := List[i]
			List[i] = List[j]
			List[j] = temp

		}
	}
	return List

}

func main() {
	slice := []int{7, 5, 3, 4, 2, 1}
	s := selectionsort(slice)
	fmt.Println(s)
}
