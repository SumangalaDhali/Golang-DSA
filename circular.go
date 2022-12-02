package main

import "fmt"

type node struct {
	data int
	next *node
}

type circular struct {
	head *node
	tail *node
}

func (l *circular) len() int {
	p := l.head
	count := 0
	for p.next != l.head {
		count++
		p = p.next
	}
	count++
	return count
}

func (l *circular) addlast(e int) {
	newest := &node{e, nil}
	if l.head == nil {
		newest.next = newest
		l.head = newest
	} else {
		newest.next = l.tail.next
		l.tail.next = newest
	}
	l.tail = newest
}

func (l *circular) addfirst(e int) {
	newest := &node{e, nil}
	if l.head == nil {
		newest.next = newest
		l.tail = newest
	} else {
		l.tail.next = newest
		newest.next = l.head
	}
	l.head = newest

}

func (l *circular) addany(e int, pos int) {
	newest := &node{e, nil}
	p := l.head
	i := 1
	for i < pos-1 {
		p = p.next
		i++
	}
	newest.next = p.next
	p.next = newest
}

func (l *circular) deletefirst() int {
	e := l.head.data
	if l.head == nil {
		fmt.Println("list is empty")
	} else {
		l.tail.next = l.head.next
		l.head = l.head.next
	}
	if l.head == nil {
		l.tail = nil
	}
	return e
}

func (l *circular) deletelast() int {
	e := 0
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		p := l.head
		for p.next != l.tail {
			p = p.next
		}
		l.tail = p
		p = p.next
		l.tail.next = l.head
		e = p.data
		p.next = nil
	}
	return e
}

func (l *circular) deleteany(pos int) int {
	p := l.head
	i := 1
	for i < pos-1 {
		p = p.next
		i++
	}
	e := p.next.data
	p.next = p.next.next
	return e
}

func (l *circular) display() {
	p := l.head
	for p.next != l.head {
		fmt.Printf("%v->", p.data)
		p = p.next
	}
	fmt.Printf("%v->", p.data)
}

func main() {
	c := circular{}
	c.addlast(10)
	c.addlast(20)
	c.addlast(30)
	c.addfirst(9)
	c.addany(90, 3)
	el := c.deletefirst()
	fmt.Println(el)
	ll := c.deletelast()
	fmt.Println(ll)
	al := c.deleteany(2)
	fmt.Println(al)
	c.display()
	fmt.Println()
	L := c.len()
	fmt.Println(L)
}
