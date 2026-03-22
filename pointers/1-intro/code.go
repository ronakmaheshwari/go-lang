package main
import "fmt"

type Message struct {
	recipient string
	text string
}

func sendMessage(m Message) {
	fmt.Printf("To: %s\n", m.recipient)
	fmt.Printf("Messages: %s\n", m.text)
}

func test(recipient, text string) {
	defer fmt.Println("=====================================")
	sendMessage(Message{
		recipient: recipient,
		text: text,
	})
}

func main() {
	test("Lane", "Textio is getting better everyday!")
	test("Allan", "This pointer stuff is weird...")
	test("Tiffany", "What time will you be home for dinner?")
}