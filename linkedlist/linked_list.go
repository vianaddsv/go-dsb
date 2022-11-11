package linkedlist

type LinkedList[N any] struct {
	head *node[N]
	last *node[N]
}

// Create pointer to linkedlist with reference head and last nil
func NewLinkedList[N any]() *LinkedList[N] {
	return &LinkedList[N]{
		head: nil,
		last: nil,
	}
}

// Get pointer to node head
func (ll *LinkedList[N]) Head() *node[N] {
	return ll.head
}

// Get pointer to node last
func (ll *LinkedList[N]) Last() *node[N] {
	return ll.last
}

// Add new node to linked list. This implementation consider
// new node are inserted at the final of list
func (ll *LinkedList[N]) Add(n N) *LinkedList[N] {

	newNode := &node[N]{
		data: n,
	}

	if ll.head == nil {
		ll.head = newNode
		ll.last = newNode
		return ll
	}

	currentLast := ll.last
	currentLast.next = newNode
	ll.last = newNode

	return ll
}

// Add new node at the beginning of list
func (ll *LinkedList[N]) AddHead(n N) *LinkedList[N] {

	newNode := &node[N]{
		data: n,
	}

	if ll.head == nil {
		ll.head = newNode
		ll.last = newNode
		return ll
	}

	currentHead := ll.head
	newNode.next = currentHead
	ll.head = newNode

	return ll
}
