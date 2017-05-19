package main

import "fmt"

func main() {
	a := 42           // int
	b := 3.142        // float64
	c := 0.867 + 0.5i // complex128
	d := "foo"        // string
	e := true	  // boolean
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	fmt.Printf("d is of type %T\n", d)
	fmt.Printf("e is of type %T\n", e)
}

