package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	root := getTreeRootWithCopy()
	invertTree(root)
	for {
		fmt.Println(root.Val)
		if root.Right == nil {
			break
		}
		root = root.Right
	}

}
func invertTree(root *Node) {
	if root == nil {
		return
	}
	invertTree(root.Left)
	invertTree(root.Right)
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	return
}

func getTreeRootWithCopy() (root *Node) {
	root = &Node{Val: 9}
	root.Right = &Node{Val: 8}
	root.Left = &Node{Val: 7}
	rootLeft := root.Left
	rootLeft.Right = &Node{Val: 6}
	rootLeft.Left = &Node{Val: 5}
	return
}
func getTreeRootWithPointer() (root *Node) {
	root = new(Node)
	//root.Val = 10
	return
}
