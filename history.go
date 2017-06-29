package lrucache

type history struct {
	head *node
	tail *node
}

type node struct {
	key  string
	next *node
}

func (h *history) find(key string) (*node, *node) {
	if h.head == nil {
		return nil, nil
	}

	var prev *node
	curr := h.head
	for curr != nil {
		if curr.key == key {
			return curr, prev
		}

		prev = curr
		curr = curr.next
	}

	return nil, nil
}

func (h *history) add(key string) {
	if h.head == nil {
		h.head = &node{
			key: key,
		}
		h.tail = h.head
		return
	}

	val, prev := h.find(key)
	if val != nil {
		// move to front...
		if val == h.tail {
			h.tail = prev
		}
		prev.next = val.next

		val.next = h.head.next
		h.head = val
		return
	}

	newHead := &node{
		key:  key,
		next: h.head,
	}

	h.head = newHead
}

func (h *history) remove(key string) {
	val, prev := h.find(key)

	if val == h.tail {
		h.tail = prev
	}

	if prev != nil {
		prev.next = val.next
	} else {
		h.head = val.next
	}
}
