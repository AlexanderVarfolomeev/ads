package main

func circle[T any](q Queue[T], n int) Queue[T] {
	for q.size != 0 && n != 0 {
		val, _ := q.Dequeue()
		q.Enqueue(val)

		n--
	}

	return q
}
