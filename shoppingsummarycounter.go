package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	mp := make(map[string]int)
	word := ""
	sl := []string{}
	for _, ch := range str {
		if ch != ' ' {
			word += string(ch)
		} else {
			sl = append(sl, word)
			word = ""
		}
	}
	sl = append(sl, word)
	for _, value := range sl {
		mp[value]++
	}
	return mp
}
