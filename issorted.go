package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	cp := 0
	cn := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			cp++
		} else if f(a[i], a[i+1]) < 0 {
			cn++
		}
	}
	if cp == len(a)-1 || cn == len(a)-1 || cp == 0 && cn == 0 {
		return true
	} else {
		return false
	}
}
