package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// short variable declaration
	// is available inside functions only
	// outside, every statement must begin with a keyword
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
