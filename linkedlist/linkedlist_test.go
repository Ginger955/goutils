package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList_Insert(t *testing.T) {
	t.Run("INSERT", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(NewNode(1))
		ll.Insert(NewNode(2, "1"))
		ll.Insert(NewNode(3))

		n := ll.Search("2")
		fmt.Println(n)
	})

	//t.Run("INSERT AT", func(t *testing.T) {
	//
	//})
}
