package main
import "fmt"

type messageToSend struct {
	message string
	phoneNumber int64
}

type rect struct {
	height int
	weigth int
}

func calc(r rect) (a int) {
	return r.height * r.weigth
}

func test(m messageToSend) {
	fmt.Printf("Sending message %s to  %d\n", m.message, m.phoneNumber)
	a := calc(rect{
		weigth: 5,
		height: 10,
	})
	fmt.Println(a)
	fmt.Println("============================")
}

func main() {
	test(messageToSend{
		message: "Ronak Maheshwari",
		phoneNumber: 9545293102,
	})
	test(messageToSend{
		message: "Danial HDFC",
		phoneNumber: 3210220045,
	})
}