package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, ch := range args {
		if ch == "01" || ch == "galaxy" || ch == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
