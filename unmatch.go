package piscine

func Unmatch(a []int) int {
	count := make(map[int]int)
	for _, ch := range a {
		count[ch]++
	}
	for _, value := range a {
		if count[value]%2 == 1 {
			return value
		}
	}
	return -1
}
