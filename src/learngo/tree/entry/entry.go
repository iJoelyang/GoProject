package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOfOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOfOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOfOrder()
	myNode.node.Print()
}

func main() {

	var root tree.Node
	root = tree.Node{Value: 3}

	// 不论地址还是结构一律使用.访问成员
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
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

	//空指针也可以调用方法
	var p *tree.Node
	p.SetValue(100)
	fmt.Println()

	root.Traverse()
	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOfOrder()

}
