package datastruct

import (
	"fmt"
	"strings"
)

// SinglyLinkedList is a singly linked list implementation using generics. This type is not safe for concurrent use
// and thus needs to be guarded using a mutex.
// This empty value of this struct is valid and can be used.
type SinglyLinkedList[V comparable] struct {
	head *singlyLinkedListNode[V]

	// makes the structure slightly more complicated but saves full list traversal when adding entries to the tail
	tail *singlyLinkedListNode[V]

	// Used to record the amount of records, to provide a quick path of bonds checking
	size uint
	// If greater than zero, modifications to the list are not allowed. Incremented while looping over linked lists
	// to prevent modification during forEach loops. Is an int to allow for recursive looping.
	blockMod int
}

// Count returns the amount of elements in the linked list
func (ll *SinglyLinkedList[V]) Count() uint {
	return ll.size
}

// Exists returns true if the element at the given index exists
func (ll *SinglyLinkedList[V]) Exists(index uint) bool {
	return index < ll.size
}

// InsertAt inserts a value at the given index
func (ll *SinglyLinkedList[V]) InsertAt(index uint, value V) {
	if ll.blockMod > 0 {
		panic("inserting and deleting not allowed while iterating over a list")
	}
	// Increment the size counter
	defer func() { ll.size++ }()

	newNode := &singlyLinkedListNode[V]{
		Value: value,
	}

	// Insert at head
	if index == 0 {
		newNode.Next = ll.head
		ll.head = newNode
		if newNode.Next == nil {
			ll.tail = newNode
		}
		return
	}

	// If we are inserting at the tail, we can use the ll.tail pointer to save having to traverse down the list
	var prev *singlyLinkedListNode[V]
	if index == ll.size {
		prev = ll.tail
	} else {
		// Get the node before the target index
		prev = ll.getNode(index - 1)
	}

	// Insert between nodes(if prev.Next == nil this logic still holds)
	newNode.Next = prev.Next
	prev.Next = newNode

	// Only the tail has nil as Next
	if newNode.Next == nil {
		ll.tail = newNode
	}
}

// InsertHead inserts the value at the head of the list
func (ll *SinglyLinkedList[V]) InsertHead(value V) {
	ll.InsertAt(0, value)
}

// InsertTail inserts the value at the tail(last element) of the list
func (ll *SinglyLinkedList[V]) InsertTail(value V) {
	ll.InsertAt(ll.size, value)
}

// Set changes the value of at the given index
func (ll *SinglyLinkedList[V]) Set(index uint, value V) {
	ll.getNode(index).Value = value
}

// DeleteAt deletes a value at the given index
func (ll *SinglyLinkedList[V]) DeleteAt(index uint) {
	if ll.blockMod > 0 {
		panic("inserting and deleting not allowed while iterating over a list")
	}

	if index == 0 {
		if ll.head == nil {
			return
		}

		newHead := ll.head.Next
		ll.head.Next = nil
		ll.head = newHead

		ll.size--

		return
	}

	n := ll.getNode(index - 1)
	n.Next = n.Next.Next
	ll.size--
}

// Get returns the value at the given index, if the given index is out of bounds the function will panic
func (ll *SinglyLinkedList[V]) Get(index uint) V {
	return ll.getNode(index).Value
}

// Loops over each value in the linked list. The linked list should not be modified from the `f` function.
func (ll *SinglyLinkedList[V]) ForEach(f func(index uint, value V) (br bool)) {
	// Disable modification to the list while running a forEach
	ll.blockMod++
	defer func() { ll.blockMod-- }()

	cur := ll.head
	i := uint(0)
	for cur != nil {
		// If the function returns true, break out of the loop
		if f(i, cur.Value) {
			return
		}

		cur = cur.Next
		i++
	}
}

// ToSlice returns a slice of all values in the list in the same order as they appear in the list
func (ll *SinglyLinkedList[V]) ToSlice() []V {
	slice := make([]V, ll.size)
	ll.ForEach(func(index uint, value V) (br bool) {
		slice[index] = value
		return false
	})
	return slice
}

// String returns a string representation of the list
func (ll SinglyLinkedList[V]) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	ll.ForEach(func(index uint, value V) (br bool) {
		if index != 0 {
			sb.WriteString(", ")
		}

		fmt.Fprint(&sb, value)
		return false
	})
	sb.WriteRune(']')
	return sb.String()
}

// Search returns the index of the first instance of the value, or -1 if no instance of the value was found
func (ll *SinglyLinkedList[V]) Search(value V) int {
	index := -1
	ll.ForEach(func(i uint, v V) (br bool) {
		if v == value {
			index = int(i)
			return true
		}
		return false
	})
	return index
}

// TODO Add Sort(f func(a, b V) int) to sort the linked list https://www.chiark.greenend.org.uk/~sgtatham/algorithms/listsort.html

func (ll *SinglyLinkedList[V]) getNode(index uint) *singlyLinkedListNode[V] {
	cur := ll.head
	for i := uint(0); i < index; i++ {
		if cur.Next == nil {
			panic("index out of bounds")
		}
		cur = cur.Next
	}
	return cur
}

func (ll *SinglyLinkedList[V]) forEachNode(f func(index uint, node *singlyLinkedListNode[V]) (br bool)) {
	// Disable modification to the list while running a forEach
	ll.blockMod++
	defer func() { ll.blockMod-- }()

	cur := ll.head
	i := uint(0)
	for cur != nil {
		// If the function returns true, break out of the loop
		if f(i, cur) {
			return
		}

		cur = cur.Next
		i++
	}
}

type singlyLinkedListNode[V any] struct {
	Value V
	Next  *singlyLinkedListNode[V]
}
