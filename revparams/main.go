package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[1:]
	for i := len(s) - 1; i >= 0; i-- {
		for _, ch := range s[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
