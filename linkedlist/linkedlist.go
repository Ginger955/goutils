package linkedlist

type LinkedList interface {
	Insert(*Node)
	InsertAt(*Node, int)

	Reverse()

	Search(id string) *Node
	SearchAt(int) *Node

	Length() int

	Delete(id string)
	DeleteAt(int)
}

type SinglyLinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
		size: 0,
	}
}

// Insert adds the node at the beginning of the list.
// The passed Node becomes the head of the list.
func (sll *SinglyLinkedList) Insert(node *Node) {
	if sll.head == nil {
		sll.head = node
	} else {
		v := sll.head
		sll.head = node
		sll.head.next = v
	}

	sll.size++
}

func (sll *SinglyLinkedList) InsertAt(node *Node, index int) {
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

func (sll *SinglyLinkedList) Reverse() {

}

func (sll *SinglyLinkedList) Search(id string) *Node {
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

func (sll *SinglyLinkedList) SearchAt(index int) *Node {
	if sll.head == nil {
		return nil
	}

	v := sll.head
	for i := 0; i <= index; i++ {
		if v.next != nil {
			v = v.next
		}
	}

	return v
}

func (sll *SinglyLinkedList) Length() int {
	return sll.size
}

func (sll *SinglyLinkedList) Delete(id string) {
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

func (sll *SinglyLinkedList) DeleteAt(index int) {
}
