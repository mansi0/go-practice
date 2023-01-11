package main

import "fmt"

type node struct {
	key int
	left *node
	right *node
}

func(r *node) inorder() {
	if(r == nil) {
		return
	}
	r.left.inorder()
	fmt.Printf("%d \t",r.key)
	r.right.inorder()
}

func(r *node) preorder() {
	if(r == nil) {
		return
	}
	fmt.Printf("%d\t",r.key)
	r.left.preorder()	
	r.right.preorder()
}

func(r *node) postorder() {
	if(r == nil) {
		return
	}
	r.left.postorder()	
	r.right.postorder()
	fmt.Printf("%d\t",r.key)
}
func main() {
	tree := &node{key:1}
	tree.left = &node{key:2}
	tree.right = &node{key:3}
	tree.left.left = &node{key:4}
	tree.left.right = &node{key:5}
	tree.inorder()
	fmt.Print(",")
	tree.preorder()
	fmt.Print(",")
	tree.postorder()
}