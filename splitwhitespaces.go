package piscine

func SplitWhiteSpaces(s string) []string {
	sl := []string{}
	str := ""
	for _, ch := range s {
		if ch != ' ' {
			str += string(ch)
		} else if str != "" {
			sl = append(sl, str)
			str = ""
		}
	}
	if str != "" {
		sl = append(sl, str)
		str = ""
	}
	return sl
}
