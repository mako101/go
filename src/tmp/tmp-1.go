package main

import "fmt"

// using naked vars
func maths_a(a, b int) (int, int, int) {
	  return a+b, a-b, a*b
}

// using named vars
func maths_b(a, b int) (x, y, z int) {
	z = a+b
	y = a-b
	z = a*b
	return
}

func main() {
	fmt.Println("Math A")
	fmt.Println(maths_a(3, 5))
	fmt.Println("\nMath B")
	fmt.Println(maths_b(5, 10))
}



