package piscine

func Map(f func(int) bool, a []int) []bool {
	sl := []bool{}
	for _, nb := range a {
		sl = append(sl, f(nb))
	}
	return sl
}
