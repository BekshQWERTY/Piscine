package piscine

// ListMerge merges two linked lists, appending the nodes of l2 to l1.
func ListMerge(l1, l2 *List) {
	// If either list is nil, return early
	if l1 == nil || l2 == nil {
		return
	}

	// If l1 is empty, set l1's head and tail to l2's head and tail
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// Otherwise, append l2's nodes to the end of l1
	current := l1.Head
	// Traverse to the last node of l1
	for current.Next != nil {
		current = current.Next
	}

	// Link the last node of l1 to the first node of l2
	current.Next = l2.Head

	// After merging, update the tail of l1 to the tail of l2
	l1.Tail = l2.Tail
}
