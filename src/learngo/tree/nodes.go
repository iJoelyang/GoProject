package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(v int) {
	if node == nil {
		fmt.Println("ignored")
		return
	}
	node.Value = v
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
