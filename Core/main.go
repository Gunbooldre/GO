package main

import "fmt"

func multiply(a int, b int) int {
	return (a + b)
}

func main() {
	msg := 19
	if msg%2 == 0 {
		fmt.Println("Here we are it is even number")
	} else {
		fmt.Println("Here we are it is odd number")
	}
	fmt.Println(multiply(1, 2))
}
