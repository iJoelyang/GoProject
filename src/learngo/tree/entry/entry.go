package main

import (
	"fmt"
	"learngo/tree"
)

func main() {

	var root tree.Node
	root = tree.Node{3, nil, nil}

	// 不论地址还是结构一律使用.访问成员
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 3, Left: nil, Right: nil},
		{Value: 5},
		{8, &root, nil},
	}
	fmt.Println(nodes)

	// 只要使用指针才能改变结构
	pRoot := &root
	pRoot.SetValue(4)
	pRoot.Print()
	fmt.Println()

	root.SetValue(10)
	root.Print()
	fmt.Println()

	var p *tree.Node
	p.SetValue(100)
	fmt.Println()

	p = &root
	p.SetValue(200)
	p.Print()
	fmt.Println()

	root.Traverse()
}
