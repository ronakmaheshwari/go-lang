package main

import (
	"fmt"
	"math"
)

func primeNumber(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Printf("%d,", n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		for i := 3; i < int(math.Sqrt(float64(n)))+1; i++ {
			if n%i == 0 {
				continue
			}
		}
		fmt.Printf("%d,", n)
	}
	fmt.Println()
}

func test(max int) {
	fmt.Printf("Primes up to %v", max)
	fmt.Println()
	primeNumber(max)
}

func main() {
	test(10)
	test(30)
	test(50)
	test(100)
}
