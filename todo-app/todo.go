package main

import (
	"fmt"
	"time"
)

type todo struct {
	title       string
	completed   bool
	createdAt   time.Time
	completedAt *time.Time
}

func add(title string, userTodo []todo) []todo {
	addTodo := todo{
		title:       title,
		completed:   false,
		createdAt:   time.Now(),
		completedAt: nil,
	}
	userTodo = append(userTodo, addTodo)
	return userTodo
}

func toggle(index int, userTodo []todo) ([]todo, error) {
	if index < 0 {
		return userTodo, fmt.Errorf("Invalid index was provided")
	} else if index >= len(userTodo) {
		return userTodo, fmt.Errorf("Invalid index was provided")
	} else {
		userTodo[index].completed = !userTodo[index].completed
		completionTime := time.Now()
		userTodo[index].completedAt = &completionTime
	}
	return userTodo, nil
}

func update(index int, title string, userTodo []todo) ([]todo, error) {
	if index < 0 {
		return userTodo, fmt.Errorf("Invalid index was provided")
	} else if index >= len(userTodo) {
		return userTodo, fmt.Errorf("Invalid index was provided")
	} else {
		userTodo[index].title= title
	}
	return userTodo, nil
}

func delete(index int, userTodo []todo) ([]todo, error) {
	if index < 0 {
		return userTodo, fmt.Errorf("Invalid index was provided")
	} else if index >= len(userTodo) {
		return userTodo, fmt.Errorf("Invalid index was provided")
	} else {
		userTodo = append(userTodo[:index],userTodo[index+1:]...)
	}
	return userTodo, nil
}