package linkedlist

type Node[T any] struct {
	id   string
	next *Node[T]
	data T
}

func NewNode[T any](data T, id ...string) *Node[T] {
	node := &Node[T]{
		id:   "",
		next: nil,
		data: data,
	}

	if len(id) > 0 {
		node.id = id[0]
	}

	return node
}
