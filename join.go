package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i, ch := range strs {
		if i != len(strs)-1 {
			str += string(ch) + sep
		} else {
			str += string(ch)
		}
	}
	return str
}
