package main

import "fmt"

func getMessagesWithRetries() [3]string {
	a := [3]string{"click here to sign up", "pretty please click here", "we beg you to sign up"}
	return a
}

func testSend(name string, doneAt int) {
	fmt.Printf("Sending to %v ...", name)
	fmt.Println()

	messages := getMessagesWithRetries()
	for i:=0;i<len(messages);i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages) - 1{
			fmt.Println("complete failure")
		} 
	}
}

func main() {
	testSend("Bob", 0)
	testSend("Alice", 1)
	testSend("Mangalam", 2)
	testSend("Ozgur", 3)
}