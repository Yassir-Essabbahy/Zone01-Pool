package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[1:]
	for _, ch := range s {
		for _, c := range ch {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
