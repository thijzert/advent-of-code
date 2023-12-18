package pq

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := PriorityQueue[byte]{}
	pq.Push('!', 12)
	pq.Push(' ', 6)
	pq.Push('l', 3)
	pq.Push('o', 5)
	pq.Push('d', 11)
	pq.Push('e', 2)
	pq.Push('l', 10)
	pq.Push('w', 7)
	pq.Push('r', 9)
	pq.Push('o', 8)
	pq.Push('h', 1)
	pq.Push('l', 4)

	t.Logf("Queue length: %d", pq.Len())

	rv := make([]byte, 0, 12)
	for pq.Len() > 0 {
		c, _, _ := pq.Pop()
		rv = append(rv, c)
	}
	t.Logf("%02x -> %s", rv, rv)
	if string(rv) != "hello, world!" {
		t.Fail()
	}
}
