package main
import "fmt"

func sendMessage(threshold float64) int {
	totalCost := 0.0
	for i:= 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > threshold {
			return i
		}
	}
}

func test(threshold float64) {
	fmt.Printf("Threshold: %.2f\n", threshold)
	max := sendMessage(threshold)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(40.00)
}
