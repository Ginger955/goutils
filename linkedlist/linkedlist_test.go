package linkedlist

import "testing"

func TestLinkedList_Insert(t *testing.T) {
	t.Run("INSERT", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(NewNode(1))
		ll.Insert(NewNode(2))
		ll.Insert(NewNode(3))
	})

	//t.Run("INSERT AT", func(t *testing.T) {
	//
	//})
}
