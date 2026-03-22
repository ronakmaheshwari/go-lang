package main

import (
	"fmt"
	"strings"
)

func removeProfanity (m *string) {
	fmt.Println(m) //addr
	msgValue := *m //access value
	msgValue = strings.ReplaceAll(msgValue , "dang", "****")
	msgValue = strings.ReplaceAll(msgValue ,"shoot","*****")
	msgValue = strings.ReplaceAll(msgValue, "heck", "****")
	*m = msgValue //adds value to it
}

func test (msg []string) {
	for _, x := range msg {
		removeProfanity(&x)
		fmt.Println(x)
	}
}

func main () {
	msg1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}
	msg2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}
	test(msg1)
	test(msg2)
}