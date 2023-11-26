package main

func SecondMax(list []int) int {
	return secondMax(list, 0, 0, false, false)
}

func secondMax(list []int, m1, m2 int, found1, found2 bool) int {
	if len(list) == 0 {
		if !found1 && !found2 {
			return 0
		}

		if !found2 {
			return m1
		}

		return m2
	}

	if !found1 {
		m1 = list[0]
		found1 = true
		return secondMax(list[1:], m1, m2, found1, found2)
	}

	if !found2 {
		m2 = list[0]
		found2 = true
		if m2 > m1 {
			m1, m2 = m2, m1
		}
		return secondMax(list[1:], m1, m2, found1, found2)
	}

	if list[0] > m1 {
		m2 = m1
		m1 = list[0]
	} else if list[0] > m2 {
		m2 = list[0]
	}

	return secondMax(list[1:], m1, m2, found1, found2)
}

func main() {
}
