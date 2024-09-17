package linkedlist

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestLinkedList_Insert(t *testing.T) {
	t.Run("INSERT", func(t *testing.T) {
		ll := NewLinkedList[int]()
		ll.Insert(NewNode(1))
		ll.Insert(NewNode(2, "1"))
		ll.Insert(NewNode(3))

		assert.Assert(t, ll.Length() == 3)
	})

	t.Run("INSERT AT", func(t *testing.T) {
		ll := NewLinkedList[int]()
		ll.InsertAt(NewNode(1), 0)
		ll.InsertAt(NewNode(2, "1"), 1)

		assert.Assert(t, ll.Length() == 2)
	})
}

func TestLinkedList_Search(t *testing.T) {
	t.Run("SEARCH", func(t *testing.T) {})
	t.Run("SEARCH AT", func(t *testing.T) {})
}

func TestLinkedList_Reverse(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Insert(NewNode(1))
	ll.Insert(NewNode(2))
	ll.Insert(NewNode(3))

	ll.Reverse()
	n1 := ll.SearchAt(0)
	n2 := ll.SearchAt(ll.Length())
	fmt.Println(n1, n2)
}

func TestLinkedList_Delete(t *testing.T) {
	t.Run("DELETE", func(t *testing.T) {})
	t.Run("DELETE AT", func(t *testing.T) {})
}
