package main

import (
	"fmt"
	"strings"
)

func main() {
	banner("Go", 6)
	banner("G", 6)
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", width))
}

// banner("Go", 6)
//   Go
// ------
