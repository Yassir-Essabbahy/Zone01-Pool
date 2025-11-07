package piscine

func Abort(a, b, c, d, e int) int {
	sl := [5]int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if sl[i] > sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
	return sl[2]
}
