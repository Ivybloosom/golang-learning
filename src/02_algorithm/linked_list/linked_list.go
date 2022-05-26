package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

func printLinkedList(head *Node) {
	temp := head
	for temp.Next != nil {
		fmt.Printf("%v -> ", temp.Value)
		temp = temp.Next
	}

	fmt.Printf("%v -> nil", temp.Value)
}

func main() {
	var a Node
	var b Node
	var c Node
	var d Node

	a = Node{"A", &b}
	b = Node{"B", &c}
	c = Node{"C", &d}
	d = Node{"1", nil}

	printLinkedList(&a)

	//fmt.Println(*a.Next)

}
