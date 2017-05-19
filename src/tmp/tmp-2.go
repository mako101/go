package main

import (
	"fmt"
)


// using function closure - ?

func incrementer() func() int {   // what is this syntax?
	// initialises i
	i := 0
	return func() int {
		// the new value of i is retained
		i++
		return i
	}
}

func foobar() func() int {   // what is this syntax?
	// initialises i
	i := 2
	return func() int {
		// the new value of i is retained
		i = i * i
		return i
	}
}


func main()  {
	// value assigned directly from function
	add := func(a, b int)int {
		return a + b
	}
	bar := add(3,4)

	fmt.Println(add) // why does this work? becomes a memory object, like in Python?
	fmt.Println(add(5, 7))
	fmt.Println(bar)

	inc := incrementer()
	fmt.Println(inc(), inc(), inc(), inc())
	foo := foobar()
	fmt.Println(foo(), foo(), foo(), foo())
}