package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	//创建对象
	var root treeNode
	root.value = 0
	root.left = &treeNode{}
	root.right = &treeNode{100, nil, nil}

	//new  创建
	root.right.left = new(treeNode)
	root.right.right = createNode(2)

	root.print()

	fmt.Println()

	root.setValue(3)
	root.print()

	fmt.Println()

	fmt.Println(root)
}
