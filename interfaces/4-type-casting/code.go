package main

import (
	"fmt"
)

type expense interface{
	cost() float64
}

type email struct {
	isSubscribed bool
	body string
	toAddress string
}

type sms struct {
	isSubscribed bool
	body string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * 0.01
	}
	return float64(len(e.body)) * 0.05
}

func (s sms) cost() float64 {
	if s.isSubscribed {
		return  float64(len(s.body)) * 0.03
	}
	return  float64(len(s.body)) * 0.1
}

func (i invalid) cost() float64 {
	return 0.0
}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func getExpenseReport(e expense) (string, float64) {
	c, ok := e.(email)
	if ok {
		return c.toAddress,c.cost()
	}
	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber,s.cost()
	} 
	return "",0.0
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})
}