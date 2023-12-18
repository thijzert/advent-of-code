package pq

type priorityNode[a any] struct {
	Value    a
	Priority int
	Length   int
	Left     *priorityNode[a]
	Right    *priorityNode[a]
}

func (nd *priorityNode[a]) Pop() (a, int, *priorityNode[a]) {
	rv, rvp := nd.Value, nd.Priority
	if nd.Left == nil && nd.Right == nil {
		return rv, rvp, nil
	} else if nd.Left == nil {
		return rv, rvp, nd.Right
	} else if nd.Right == nil {
		return rv, rvp, nd.Left
	}

	nd.Length--
	if nd.Left.Priority < nd.Right.Priority {
		nd.Value, nd.Priority, nd.Left = nd.Left.Pop()
	} else {
		nd.Value, nd.Priority, nd.Right = nd.Right.Pop()
	}
	return rv, rvp, nd
}

func (nd *priorityNode[a]) Push(v a, priority int) *priorityNode[a] {
	if priority < nd.Priority {
		nv, nvp := nd.Value, nd.Priority
		nd.Value, nd.Priority = v, priority
		return nd.Push(nv, nvp)
	}

	nd.Length++
	if nd.Left == nil {
		nd.Left = &priorityNode[a]{Value: v, Priority: priority, Length: 1}
	} else if nd.Right == nil {
		nd.Right = &priorityNode[a]{Value: v, Priority: priority, Length: 1}
	} else if nd.Right.Length < nd.Left.Length {
		nd.Right = nd.Right.Push(v, priority)
	} else {
		nd.Left = nd.Left.Push(v, priority)
	}
	return nd
}

type PriorityQueue[a any] struct {
	root *priorityNode[a]
}

func (pq *PriorityQueue[a]) Push(v a, priority int) {
	if pq.root == nil {
		pq.root = &priorityNode[a]{Value: v, Priority: priority, Length: 1}
	} else {
		pq.root = pq.root.Push(v, priority)
	}
}

func (pq *PriorityQueue[a]) Pop() (a, int, bool) {
	if pq.root == nil {
		var zero a
		return zero, 0, false
	}
	rv, rvp, nnd := pq.root.Pop()
	pq.root = nnd
	return rv, rvp, true
}

func (pq *PriorityQueue[a]) Len() int {
	if pq.root == nil {
		return 0
	}
	return pq.root.Length
}
