package main

import "fmt"

type node struct {
	data int
	next *node
}

type qlinked struct {
	front *node
	rear  *node
	size  int
}

func (s *qlinked) len() int {
	count := 0
	p := s.front
	for p != nil {
		p = p.next
		count++
	}
	return count
}

func (s *qlinked) isempty() bool {
	return s.len() == 0
}

func (s *qlinked) enqueue(e int) {
	newest := &node{e, nil}
	if s.isempty() {
		s.front = newest
	} else {
		s.rear.next = newest
	}
	s.rear = newest
}

func (s *qlinked) dequeue() int {
	if s.isempty() {
		fmt.Println("Queue is empty")
		return 0
	} else {
		e := s.front.data
		s.front = s.front.next
		return e
	}

}

func (s *qlinked) first() int {
	if s.isempty() {
		fmt.Println("stack is empty")
		return 0
	}
	return s.front.data
}

func (s *qlinked) display() {
	p := s.front
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}

func main() {
	Q1 := qlinked{}
	Q1.enqueue(10)
	Q1.enqueue(20)
	Q1.enqueue(30)
	Q1.enqueue(40)
	Q1.enqueue(50)
	Q1.display()
	length := Q1.len()
	fmt.Println("the length of the Queue is")
	fmt.Println(length)
	fmt.Println("Printing the queues first element")
	fmt.Println(Q1.first())
	fmt.Println(Q1.dequeue())
}
