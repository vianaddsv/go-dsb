package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateEmptyLinkedList(t *testing.T) {

	ll := NewLinkedList[string]()
	assert.NotNil(t, ll)
	assert.Empty(t, ll.Head())
}

func TestShouldAddNewNode(t *testing.T) {
	ll := NewLinkedList[string]().
		Add("Daniel").
		Add("Larissa").
		Add("Rodrigo").
		Add("Laura")

	expected := map[int]string{
		0: "Daniel",
		1: "Larissa",
		2: "Rodrigo",
		3: "Laura",
	}

	next := ll.Head()
	for i := 0; i <= 2; i++ {
		assert.Equal(t, expected[i], next.Data(), "itens aren't in same order")
		next = next.Next()
	}

}

func TestShouldGetHead(t *testing.T) {
	ll := NewLinkedList[string]().
		Add("Daniel").
		Add("Larissa").
		Add("Rodrigo").
		Add("Laura")

	expected := "Daniel"
	assert.Equal(t, expected, ll.Head().Data(), "Daniel isn't the first item")
}

func TestShouldGetLast(t *testing.T) {
	ll := NewLinkedList[string]().
		Add("Daniel").
		Add("Larissa").
		Add("Rodrigo").
		Add("Laura")

	expected := "Laura"
	assert.Equal(t, expected, ll.Last().Data(), "Laura isn't the last item")
}

func TestShouldGetSameHeadAdded(t *testing.T) {
	ll := NewLinkedList[string]().
		Add("Larissa").
		Add("Rodrigo").
		Add("Laura")

	expected := "Daniel"
	ll.AddHead(expected)

	assert.Equal(t, expected, ll.Head().Data(), "Daniel isn't the First item")
}
