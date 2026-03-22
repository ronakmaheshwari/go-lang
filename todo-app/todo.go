package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
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
		userTodo[index].title = title
	}
	return userTodo, nil
}

func delete(index int, userTodo []todo) ([]todo, error) {
	if index < 0 {
		return userTodo, fmt.Errorf("Invalid index was provided")
	} else if index >= len(userTodo) {
		return userTodo, fmt.Errorf("Invalid index was provided")
	} else {
		userTodo = append(userTodo[:index], userTodo[index+1:]...)
	}
	return userTodo, nil
}

func print(userTodo []todo) {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#Id", "Title", "Completed", "Created At", "Completed At")
	for i, x := range userTodo {
		completedAt := ""
		if x.completed {
			if x.completedAt != nil {
				completedAt = x.completedAt.Format(time.RFC1123)
			}
		}
		table.AddRows([]string{strconv.Itoa(i), x.title, fmt.Sprintf("%v", x.completed), x.createdAt.Format(time.RFC1123), completedAt})
	}
	table.Render()
}
