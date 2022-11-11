package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetNextNode(t *testing.T) {

	node := &node[string]{
		next: &node[string]{
			data: "oi",
			next: nil,
		},
		data: "hello,",
	}

	assert.NotNil(t, node.Next())
	assert.Nil(t, node.Next().Next())
}

func TestShouldGetData(t *testing.T) {
	firstNodeValue, secondeNodeValue := "data1", "data2"
	node := &node[string]{
		next: &node[string]{
			data: secondeNodeValue,
			next: nil,
		},
		data: firstNodeValue,
	}

	assert.Equal(t, node.Data(), firstNodeValue)
	assert.Equal(t, node.Next().Data(), secondeNodeValue)
}
