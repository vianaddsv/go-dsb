package linkedlist

type node[N any] struct {
	next *node[N]
	data N
}

// Get next pointer of node
func (n node[N]) Next() *node[N] {
	return n.next
}

// Get value of node
func (n node[N]) Data() N {
	return n.data
}
