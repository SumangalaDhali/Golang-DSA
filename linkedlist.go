package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head *node
	tail *node
}

func (l *linkedlist) pushlast(e int) {
	newest := &node{data: e}
	if l.head == nil {
		l.head = newest
	} else {
		l.tail.next = newest
	}
	l.tail = newest
}

func (l *linkedlist) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, "->")
		p = p.next
	}

}

func (l *linkedlist) len() int {
	count := 0
	p := l.head
	for p != nil {
		count++
		p = p.next
	}
	return count

}
func (l *linkedlist) pushfirst(e int) {
	newest := node{data: e}
	if l.head == nil {
		l.head = &newest
		l.tail = &newest
	} else {
		newest.next = l.head
		l.head = &newest
	}
}

func (l *linkedlist) deletefirst() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		e := l.head.data
		l.head = l.head.next
		fmt.Println(e)
	}
	if l.head == nil {
		l.tail = nil
	}
}
func (l *linkedlist) Pushany(e, Pos int) {
	newest := node{data: e}
	p := l.head
	i := 1
	for i < Pos-1 {
		p = p.next
		i++
	}
	newest.next = p.next
	p.next = &newest

}

func (l *linkedlist) deletelast() {
	if l.head == nil {
		fmt.Println("the list is empty")
	} else {
		p := l.head
		i := 0
		for i < l.len()-1 {
			p = p.next
			i++
		}
		l.tail = p
		p = p.next
		e := p.data
		l.tail.next = nil
		fmt.Println(e)
	}
}

func main() {
	l := linkedlist{}
	l.pushlast(10)
	l.pushlast(20)
	l.pushfirst(9)
	l.pushfirst(8)
	l.Pushany(22, 3)
	l.display()
	fmt.Println()
	l.deletefirst()
	fmt.Println()
	n := l.len()
	fmt.Println(n)
	l.deletelast()
	l.display()

}
