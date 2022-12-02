/*package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60}
	i := binarysearch(s, 50)
	fmt.Println(i)

}

// returns the index
// return -1 if it does not exist
func binarysearch(arr []int, key int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2
		if key < arr[mid] {
			end = mid - 1
		} else if key > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}*/

package main

import "fmt"

func main() {
	s := []int{-1, -2, 0, 10, 20, 30, 40, 50, 60}
	i := binarysearch(s, -2)
	fmt.Println(i)

}

// returns the index
// return -1 if it does not exist
func binarysearch(arr []int, key int) int {
	start := 0
	end := len(arr) - 1

	// find wheather the array is sorted in acsending or descending order
	isAsc := arr[start] < arr[end]

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return mid
		}
		if isAsc {
			if key < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if key > arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

	}
	return -1
}
