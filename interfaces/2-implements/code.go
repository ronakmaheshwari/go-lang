package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

type fullTime struct {
	name string
	salary int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func (t fullTime) getName() string {
	return t.name
}

func (t fullTime) getSalary() int {
	return t.salary
}

func print(c employee) {
	fmt.Println(c.getName())
	fmt.Println(c.getSalary())
}

func main() {
	print(contractor{
		name: "Ronak",
		hourlyPay: 5,
		hoursPerYear: 10,
	})
	print(fullTime{
		name: "Jene",
		salary: 500000,
	})
}
