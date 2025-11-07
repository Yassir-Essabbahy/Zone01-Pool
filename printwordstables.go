package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		r := []rune(a[i])
		for a := 0; a < len(r); a++ {
			z01.PrintRune(r[a])
		}
		z01.PrintRune('\n')
	}
}
