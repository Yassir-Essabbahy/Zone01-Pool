package piscine

func ActiveBits(n int) int {
	str := ""
	for n >= 1 {
		str += string(rune((n % 2) + '0'))
		n = n / 2
	}
	count := 0
	for _, ch := range str {
		if ch == '1' {
			count++
		}
	}
	return count
}
