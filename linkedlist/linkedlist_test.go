package linkedlist

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestLinkedList_Insert(t *testing.T) {
	t.Run("INSERT", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(NewNode(1))
		ll.Insert(NewNode(2, "1"))
		ll.Insert(NewNode(3))

		assert.Assert(t, ll.Length() == 3)
	})

	t.Run("INSERT AT", func(t *testing.T) {
		ll := NewLinkedList()
		ll.InsertAt(NewNode(1), 0)
		ll.InsertAt(NewNode(2, "1"), 1)

		assert.Assert(t, ll.Length() == 2)
	})
}

func TestLinkedList_Delete(t *testing.T) {

}

func TestLinkedList_DeleteAt(t *testing.T) {}
