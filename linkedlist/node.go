package linkedlist

type Node struct {
	next *Node
	data any
}

func NewNode(data any) *Node {
	return &Node{
		next: nil,
		data: data,
	}
}
