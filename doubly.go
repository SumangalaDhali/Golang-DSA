package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type dlinked struct {
	head *node
	tail *node
}

func (d *dlinked) len() int {
	p := d.head
	count := 0
	for p != nil {
		count++
		p = p.next
	}
	return count

}

func (d *dlinked) addfirst(e int) {
	new_node := &node{e, nil, nil}
	if d.head == nil {
		d.tail = new_node
	} else {
		new_node.next = d.head
		d.head.prev = new_node
	}
	d.head = new_node
}
func (d *dlinked) addlast(e int) {
	new_node := &node{e, nil, nil}
	if d.tail == nil {
		d.head = new_node
	} else {
		d.tail.next = new_node
		new_node.prev = d.tail
	}
	d.tail = new_node
}
func (d *dlinked) addany(e, pos int) {
	new_node := &node{e, nil, nil}
	p := d.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	new_node.next = p.next
	p.next.prev = new_node
	p.next = new_node
	new_node.prev = p
}

func (d *dlinked) removefirst() int {
	if d.head == nil {
		fmt.Println("list is empty")
		return 0
	} else {
		e := d.head.data
		d.head = d.head.next
		d.head.prev = nil
		if d.head == nil {
			d.tail = nil
		}
		return e
	}
}
func (d *dlinked) removelast() int {
	if d.tail == nil {
		fmt.Println("list is empty")
		return 0
	} else {
		e := d.tail.data
		d.tail = d.tail.prev
		d.tail.next = nil
		return e
	}
}

func (d *dlinked) display() {
	p := d.head
	for p != nil {
		fmt.Printf("%d->", p.data)
		p = p.next
	}
}

func main() {
	d1 := dlinked{}
	d1.addfirst(10)
	d1.addfirst(20)
	d1.addfirst(30)
	d1.addlast(40)
	d1.addlast(50)
	d1.addany(90, 3)
	fmt.Println(d1.removefirst())
	fmt.Println(d1.removelast())
	d1.display()
	fmt.Println()
	n := d1.len()
	fmt.Println(n)

}
