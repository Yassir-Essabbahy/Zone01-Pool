package piscine

func ReverseMenuIndex(menu []string) []string {
	sl := make([]string, len(menu))
	i := len(menu) - 1
	for j := 0; j < len(menu); j++ {
		sl[j] = menu[i]
		i--
	}
	return sl
}
