package main

import "fmt"

func main() {

	// array - fixed-length list. Number of elements must be declared!
	var my_array[2]int
	my_array[0], my_array[1] = 4, 5
	fmt.Println("Array my_array", my_array, "length", len(my_array))

	// slice - a flexible list!
	list := []int{11, 22}
	list = append(list, 33, 44)
	fmt.Println("List:", list, "length", len(list))

	//pointer - holds the memory address of the variable
	var p1 *int // declare a pointer
	i := 4	    // assign an integer to a variable
	p1 = &i     // assign address of the integer to the pointer
	fmt.Println(p1)

	p2 := &i   	 // create a pointer via type inference
	fmt.Println(p2)
	fmt.Println(&i)  // print the memory address directly
	fmt.Println(*p2) // use * to 'dereference'  the pointer and get an original value
	fmt.Println(p2)  // pointer remains assigned

	GetX()
	GetY()
}

//pointer use case - using var vs using value of the var

// changing the variable value DOES NOT change the variable in memory
func GetX() {
	x := 5
	ChangeX(x)
	fmt.Println("X is", x)
}

func ChangeX(x int) {
	x = 10
}

func GetY() {
	y := 5
	ChangeY(&y)
	fmt.Println("Y is", y)
}

func ChangeY(Y *int) {
	*Y = 10
}