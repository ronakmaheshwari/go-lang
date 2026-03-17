// Functions
// Functions in Go can take zero or more arguments.

// To make Go code easier to read, the variable type comes after the variable name.

// For example, the following function:

// func sub(x int, y int) int {
//   return x-y
// }
// Accepts two integer parameters and returns another integer.

// Here, func sub(x int, y int) int is known as the "function signature".

// Assignment
// We often will need to manipulate strings in our messaging app. For example, adding some personalization by using a customer's name within a template. The concat function should take two strings and smash them together.

// hello + world = helloworld
// Fix the function signature of concat to reflect its behavior.

package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	if a <= 0 || b <= 0 {
		return -1
	}
	return a * b
}

func divide(a int, b int) int {
	if b == 0 {
		return -1
	}

	if a <= 0 || b <= 0 {
		return -1
	}

	return a / b
}

func main() {
	fmt.Println(concat("Ronak", "Maheshwari"))
	fmt.Println(add(10, 13))
	fmt.Println(sub(20, 10))
	fmt.Println(multiply(10,5))
	fmt.Println(divide(12, 2))
}
