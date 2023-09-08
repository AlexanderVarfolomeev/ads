package main

import "errors"

func Sum(l1 LinkedList, l2 LinkedList) (LinkedList, error) {
	res := LinkedList{}
	n1, n2 := l1.head, l2.head

	for n1 != nil || n2 != nil {
		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
			return LinkedList{}, errors.New("list sizes are not equal")
		}

		res.AddInTail(Node{value: n1.value + n2.value})

		n1 = n1.next
		n2 = n2.next
	}

	return res, nil
}
