package main

import "fmt"

type node struct {
	data int
	next *node
}

type stlinked struct {
	top *node
}

func (s *stlinked) len() int {
	count := 0
	p := s.top
	for p != nil {
		p = p.next
		count++
	}
	return count
}

func (s *stlinked) isempty() bool {
	return s.len() == 0
}

func (s *stlinked) push(e int) {
	newest := &node{e, nil}
	if s.isempty() {
		s.top = newest
	} else {
		newest.next = s.top
		s.top = newest
	}
}

func (s *stlinked) pop() int {
	if s.isempty() {
		fmt.Println("stack is empty")
		return 0
	} else {
		e := s.top.data
		s.top = s.top.next
		return e
	}
}

func (s *stlinked) peek() int {
	if s.isempty() {
		fmt.Println("stack is empty")
		return 0
	}
	return s.top.data
}

func (s *stlinked) display() {
	p := s.top
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}

}

func main() {
	l := stlinked{}
	l.push(10)
	l.push(20)
	l.push(30)
	l.display()
	fmt.Println("Popped elements")
	ele := l.pop()
	fmt.Println(ele)
	length := l.len()
	fmt.Println(length)
	l.display()
	peek := l.peek()
	fmt.Println(peek)

}
