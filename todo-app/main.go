package main

import "fmt"

func main() {
	userTodo := []todo{}
	userTodo = add("Learn Go", userTodo)
	userTodo = add("Build CLI", userTodo)
	a, err := delete(0, userTodo)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(a)
}