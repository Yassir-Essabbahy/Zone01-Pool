package piscine

func DescendAppendRange(max, min int) []int {
	sl := []int{}
	if max > min {
		for i := max; i > min; i-- {
			sl = append(sl, i)
		}
	} else {
		return sl
	}
	return sl
}
