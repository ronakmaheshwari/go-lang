package main
import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = increments(sendsSoFar,sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

func increments(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}