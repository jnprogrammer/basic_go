package coolpack

func Sum(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
