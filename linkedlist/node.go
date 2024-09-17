package linkedlist

type Node struct {
	id   string
	next *Node
	data any
}

func NewNode(data any, id ...string) *Node {
	node := &Node{
		id:   "",
		next: nil,
		data: data,
	}

	if len(id) > 0 {
		node.id = id[0]
	}

	return node
}
