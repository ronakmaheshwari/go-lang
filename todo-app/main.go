package main

import "fmt"

func main() {
	userTodo := []todo{}
	userTodo = add("Learn Go", userTodo)
	userTodo = add("Build CLI", userTodo)
	// userTodo, err := delete(0, userTodo)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	userTodo, err := toggle(0,userTodo)
	if err != nil {
		fmt.Print(err)
	}
	print(userTodo)
}