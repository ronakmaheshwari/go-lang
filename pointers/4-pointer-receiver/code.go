package main
import "fmt"

type email struct {
	fromAddress string
	toAddress string
	message string
}

func (e *email) setMessage(newMsg string) {
	e.message = newMsg
}

func test(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage("this is my second draft")
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)

}

func main() {
	test(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	test(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")
}