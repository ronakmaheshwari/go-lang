// Constants
// Constants are declared like variables but use the const keyword. Constants can't use the := short declaration syntax.

// Constants can be character, string, boolean, or numeric values. They can not be more complex types like slices, maps and structs, which are types we will explain later.

// As the name implies, the value of a constant can't be changed after it has been declared.

// Use two separate constants
// Something weird is happening in this code.

// What should be happening is that we create 2 separate constants: premiumPlanName and basicPlanName. Right now it looks like we're trying to overwrite one of them.
package main
import "fmt"

func main() {
	const premiumPlan = "Premium Plan"
	const basicPlan = "Basic Plan"
	const number = 10
	fmt.Println("The Plan:", premiumPlan)
	fmt.Println("The Plan:",basicPlan)
	fmt.Println("the number:", number)
}