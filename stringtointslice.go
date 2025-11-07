package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	sl := []int{}
	for _, ch := range str {
		sl = append(sl, int(ch))
	}
	return sl
}
