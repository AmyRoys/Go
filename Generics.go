package main

import "fmt"

func Index[T comparable] (s []T, x T) int {
	for i, v := range s{
		if v == x{
			return i 
		}
	}
	return -1
}

type List[T any] struct{
	next *List[T]
	val T
}

func(l *List[T]) Push(val T){
	node := &List[T]{val:val}
	node.next = l.next
	l.next = node
}

func main() {
	si := []int{10,20,15,-10}
	fmt.Println(Index(si, 15))

	ss:= []string{"hello", "bar", "baz"}
	fmt.Println(Index(ss,"hello"))

	list := &List[int]{}
	list.Push(1)
	fmt.Println(list.val)

}