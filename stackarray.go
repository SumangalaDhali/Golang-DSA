package main

import "fmt"

type stackarray struct {
	array []int
}

func (s *stackarray) len() int {
	length := len(s.array)
	return length

}

func (s *stackarray) isempty() bool {
	length := len(s.array)
	return length == 0
}

func (s *stackarray) display() {
	length := len(s.array)
	for i := 0; i < length; i++ {
		fmt.Println(s.array[i])
	}
}

func (s *stackarray) push(e int) {
	s.array = append(s.array, e)
}

func (s *stackarray) pop() int {
	length := len(s.array)
	res := s.array[length-1]
	s.array = s.array[:length-1]
	return res

}

func (s *stackarray) top() {
	length := len(s.array)
	fmt.Println(s.array[length-1])
}

func main() {
	c := stackarray{}
	c.push(2)
	c.push(3)
	c.push(4)
	c.push(5)
	e := c.pop()
	fmt.Println(e)
	c.display()
	fmt.Println(c.array)
}
