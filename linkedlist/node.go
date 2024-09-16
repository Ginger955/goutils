package linkedlist

type Node struct {
	id   string
	next *Node
	data any
}

func NewNode(data any) *Node {
	return &Node{
		id:   "",
		next: nil,
		data: data,
	}
}
