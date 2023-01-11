package main
import "fmt"

//queue 
type queue []*node

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) enQueue(l *node) {
	*q = append(*q, l)
}

func (q *queue) deQueue() (*node) {
	if q.isEmpty() {
		return nil
	}
	element := (*q)[0]

	*q = (*q)[1:]
	return element
}
func (n *node) levelOrder() {
	if n == nil {
		return
	}
	var q1 queue
	q1.enQueue(n)
	q1.enQueue(nil)
	for len(q1) > 1 {
		curr := q1.deQueue()

		if(curr == nil) {
			q1.enQueue(nil)
			fmt.Println()
		} else{
			if(curr.left!= nil) {
				q1.enQueue(curr.left)
			}
			if(curr.right != nil) {
				q1.enQueue(curr.right)
			}
			fmt.Print(curr.key," ")
		}
	}
}
//tree node
type node struct {
	key int
	left *node
	right *node
}


func main() {
	tree := &node{key:1}
	tree.left = &node{key:2}
	tree.right = &node{key:3}
	tree.left.left = &node{key:4}
	tree.left.right = &node{key:5}
	tree.right.right = &node{key:6}
	tree.levelOrder()
	// fmt.Println(tree)
}