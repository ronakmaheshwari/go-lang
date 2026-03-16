/* 
Fix a Bug!
Textio users are reporting that we're billing them for wildly inaccurate amounts. They're supposed to be billed .02 dollars for each text message sent.

Fix the math bug on line 17.
*/

package main
import "fmt"

func main() {
	messagesFromDoris := []string {
		"You doing anything later??",
		"Did you get my last message?",
		"Don't leave me hanging...",
		"Please respond I'm lonely!",
	}
	numMessage := float64(len(messagesFromDoris))
	per_msg_cost := .02

	totalCost := numMessage * per_msg_cost

	fmt.Printf("Doris spent %.02f on each text message today", totalCost)
}