package linkedlist

type LinkedList interface {
	Insert(*Node)
	InsertAt(*Node, int)

	//Reverse()
	//Search()
	Length() int
	//Delete(key any)
	//DeleteAt(int)
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

func (sll *SinglyLinkedList) Length() int {
	return sll.size
}
