package main

import "errors"

func Sum(l1 LinkedList, l2 LinkedList) (LinkedList, error) {
	if l1.Count() != l2.Count() {
		return LinkedList{}, errors.New("list sizes are not equal")
	}

	res := LinkedList{}
	n1, n2 := l1.head, l2.head

	for n1 != nil && n2 != nil {
		res.AddInTail(Node{value: n1.value + n2.value})

		n1 = n1.next
		n2 = n2.next
	}

	return res, nil
}
