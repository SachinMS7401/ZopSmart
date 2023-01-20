package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	head := Node{}

	head.insert(2)
	head.insert(3)
	head.insert(4)
	head.insert(5)
	head.insert(7)
	head.insert(5)
	head.display()
	head.delete_(2)
	head.display()
	head.delete_val(5)
	head.display()

}
func (head *Node) insert(val int) {
	t := Node{}
	t.data = val
	if head == nil {
		head = &t

	} else {
		cur := head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = &t
	}
}

func (head *Node) display() {
	cur := head.next
	fmt.Printf("Linked list elements are:")
	if cur != nil {
		for cur != nil {
			fmt.Print(cur.data)
			cur = cur.next
		}
		fmt.Println()
	} else {
		fmt.Println("empty list")
	}
}

func (head *Node) delete_(pos int) {
	cur := head
	c := 0
	for c < pos-1 {
		cur = cur.next
		c += 1
	}
	cur.next = cur.next.next
}

func (head *Node) delete_val(val int) {
	cur := head
	for cur.next.data != val {
		cur = cur.next
	}
	cur.next = cur.next.next
}
