package linkedlist

type LinkedList[T any] interface {
	Insert(*Node[T])
	InsertAt(*Node[T], int)

	Reverse()

	Search(id string) *Node[T]
	SearchAt(int) *Node[T]

	Length() int

	Delete(id string)
	DeleteAt(int)
}

type SinglyLinkedList[T any] struct {
	head *Node[T]
	size int
}

func NewLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
		size: 0,
	}
}

// Insert adds the node at the beginning of the list.
// The passed Node becomes the head of the list.
func (sll *SinglyLinkedList[T]) Insert(node *Node[T]) {
	if sll.head == nil {
		sll.head = node
	} else {
		v := sll.head
		sll.head = node
		sll.head.next = v
	}

	sll.size++
}

func (sll *SinglyLinkedList[T]) InsertAt(node *Node[T], index int) {
	if sll.head == nil {
		sll.head = node
	} else {
		v := sll.head
		for i := 0; i <= index; i++ {
			if v.next != nil {
				v = v.next
			}
		}
		v.next = node
	}

	sll.size++
}

func (sll *SinglyLinkedList[T]) Reverse() {
	if sll.head == nil {
		return
	}

	cpy := NewLinkedList[T]()

	v := sll.head

	for v != nil {
		n := NewNode(v.data.(T), v.id)
		cpy.Insert(n)
		v = v.next
	}

	sll.head = cpy.head
}

func (sll *SinglyLinkedList[T]) Search(id string) *Node[T] {
	if sll.head == nil {
		return nil
	}

	v := sll.head
	for v != nil {
		if v.id == id {
			return v
		}

		v = v.next
	}

	return nil
}

func (sll *SinglyLinkedList[T]) SearchAt(index int) *Node[T] {
	if sll.head == nil {
		return nil
	}

	v := sll.head
	for i := 0; i < index; i++ {
		if v.next != nil {
			v = v.next
		}
	}

	return v
}

func (sll *SinglyLinkedList[T]) Length() int {
	return sll.size
}

func (sll *SinglyLinkedList[T]) Delete(id string) {
	if sll.head == nil {
		return
	}

	v := sll.head
	prev := v
	for v != nil {
		if v.id == id {
			if v.next != nil {
				prev.next = v.next
			} else {
				prev.next = nil
			}

			sll.size--
		}

		prev = v
		v = v.next
	}
}

func (sll *SinglyLinkedList[T]) DeleteAt(index int) {
}
