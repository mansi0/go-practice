package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkdlist struct {
	head   *node
	length int
}

func (l *linkdlist) postpend(value int) {
	newNode := node{data: value}
	if l.length == 0 {
		l.head = &newNode
		l.length++
		return
	}
	isListContain,new := l.listContain(value)
	// fmt.Println(isListContain,new)
	curr := l.head
	if(isListContain) {
		
		for curr.next!=nil {
			curr = curr.next
		}
		curr.next = new
		return
	}
	
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &newNode
	l.length++
}

func(l *linkdlist) listContain( value int)(bool, *node) {
	curr := l.head
	for curr != nil {
		if(curr.data == value) {
			return true,curr
		}
		curr = curr.next
	}
	return false,nil
}

func (l *linkdlist) printList() {
	if l.head == nil {
		return
	}
	cnt :=8
	currNode := l.head
	for cnt>0 {
		fmt.Println(currNode.data)
		currNode = currNode.next
		cnt--
	}
}

func (l *linkdlist) isContainCycle() (bool, int) {
	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if(fast == slow) {
			cnt := 1
			slow = slow.next
			for fast != slow {
				slow = slow.next
				cnt += 1
			}
			return true, cnt
		}
	}
	return false , 0
}
func main() {
	myList := linkdlist{}
	myList.postpend(1)
	myList.postpend(2)
	myList.postpend(3)
	myList.postpend(4)
	myList.postpend(2)
	// myList.printList()
	isContainCycle, len := myList.isContainCycle()
	fmt.Println(isContainCycle,len)
}
