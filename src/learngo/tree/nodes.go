package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{3, nil, nil}

	// 不论地址还是结构一律使用.访问成员
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3, left: nil, right: nil},
		{value: 5},
		{8, &root, nil},
	}
	fmt.Println(nodes)
}
