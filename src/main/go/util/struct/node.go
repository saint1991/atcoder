package _struct

type Node struct {
	Key   string
	Value interface{}
}

type IntNode struct {
	Key   string
	Value int
}

type BiNode struct {
	Node
	Left  *Node
	Right *Node
}

type BiIntNode struct {
	IntNode
	Left  *Node
	Right *Node
}
