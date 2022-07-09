package main

import "fmt"

type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	lenght int
}

func (l *linkedList) prepend(n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.lenght++
}

func (l *linkedList) deletewithvalue (value int){
	if l.lenght == 0{
		return
	}

	if l.head.data == value{
		l.head = l.head.next
		l.lenght--
		return
	}
	previoustoDelete := l.head
	for previoustoDelete.next.data != value{
		if previoustoDelete.next.next == nil{
			return
		}
		previoustoDelete = previoustoDelete.next
	}
	previoustoDelete.next = previoustoDelete.next.next
	l.lenght--
}

func (l linkedList) printlist(){
	tyPrint := l.head
	for l.lenght != 0{
		fmt.Printf("%d ", tyPrint.data)
		tyPrint = tyPrint.next
		l.lenght--
	}
	fmt.Printf("\n")
}

func main(){
	mylist := linkedList{}
	node1  := &node{data:48}
	node2  := &node{data:20}
	node3  := &node{data:10}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	fmt.Println(mylist)
	mylist.printlist()
	mylist.deletewithvalue(200)
	mylist.printlist()
}