package no1796

//SecondHighest 一次遍历
func SecondHighest(s string) int {
	a, b := -1, -1
	for _, c := range s {
		if c >= '0' && c <= '9' {
			v := int(c - '0')
			if v > a {
				b, a = a, v
			} else if v > b && v < a {
				b = v
			}
		}
	}
	return b
}

// SecondHighest2 位运算
func SecondHighest2(s string) int {
	mask := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			mask |= 1 << int(c-'0')
		}
	}
	for i, cnt := 9, 0; i >= 0; i-- {
		if mask>>i&1 == 1 {
			cnt++
			if cnt == 2 {
				return i
			}
		}
	}
	return -1
}
