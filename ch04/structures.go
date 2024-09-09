package main

import "fmt"

// Implementation of a linked list, using generics

type node[T any] struct {
	Data T
	next *node[T]
}

type list[T any] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		next: nil,
	}
	if l.start == nil {
		l.start = &n
		return
	}
	if l.start.next == nil {
		l.start.next = &n
		return
	}
	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
}

func main() {
	var myList list[int]
	fmt.Println(myList)

	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)

	cur := myList.start
	for {
		fmt.Println("*", cur)
		if cur == nil {
			break
		}
		cur = cur.next
	}
}