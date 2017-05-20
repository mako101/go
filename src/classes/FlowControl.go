package main

import (
	"fmt"
	"runtime"
)

var i int = 1

func main() {
	// defer runs just before the function exits
	defer fmt.Println("Finished!")
	if_loop(i)
	switch_case()

}

// if / else
func if_loop(i int) {
	// curly bracket has to be on the same line
	if i < 0 {
		fmt.Println("Negative")
	// the next condition needs to begin after the closing bracket
	} else if i == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Positive")
	}


}

// case statements

func switch_case() {
	// can initialize variable inside
	// the flow control statement
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
		// this forces te flow to enter (and execute!) the next case
		// even if it will not match
		fallthrough
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Println(os)
	}
}