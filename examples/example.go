package main

import (
	"fmt"

	parser "github.com/ppalone/duration-parser"
)

func main() {
	s := "01:30"
	d, _ := parser.Parse(s)
	fmt.Println("Duration in seconds:", d.Seconds())
}
