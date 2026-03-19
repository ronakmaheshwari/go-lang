package main
import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouce string) (float64, error) {
	costForCustomer, err := sendSms(msgToCustomer)
	if err != nil {
		return 0.0,err
	}
	costForSpouce, err := sendSms(msgToSpouce)
	if err != nil {
		return 0.0,err
	}
	return costForCustomer + costForSpouce, nil
}

func sendSms(message string) (float64, error) {
	const maxTextLength = 25
	const costPerChar = .0002
	if int(len(message)) > maxTextLength {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLength)
	}
	return float64(len(message)) * costPerChar, nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {
	test(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	test(
		"Thank you.",
		"Enjoy!",
	)
	test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)
}