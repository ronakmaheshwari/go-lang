// Formatting Strings in Go
// Go follows the printf tradition from the C language. In my opinion, string formatting/interpolation in Go is currently less elegant than JavaScript and Python.

// fmt.Printf - Prints a formatted string to standard out.
// fmt.Sprintf() - Returns the formatted string
// Examples
// %v - Interpolate the default representation
// The %v variant prints the Go syntax representation of a value. You can usually use this if you're unsure what else to use. That said, it's better to use the type-specific variant if you can.

// fmt.Printf("I am %v years old", 10)
// // I am 10 years old

// fmt.Printf("I am %v years old", "way too many")
// // I am way too many years old
// %s - Interpolate a string
// fmt.Printf("I am %s years old", "way too many")
// // I am way too many years old
// %d - Interpolate an integer in decimal form
// fmt.Printf("I am %d years old", 10)
// // I am 10 years old
// %f - Interpolate a decimal
// fmt.Printf("I am %f years old", 10.523)
// // I am 10.523000 years old

// // The ".2" rounds the number to 2 decimal places
// fmt.Printf("I am %.2f years old", 10.523)
// // I am 10.53 years old
// If you're interested in all the formatting options, feel free to take a look at the fmt package's docs here.

// Assignment
// Create a new variable called msg on line 9. It's a string that contains the following:

// Hi NAME, your open rate is OPENRATE percent
// Where NAME is the given name, and OPENRATE is the openRate rounded to the nearest "tenths" place.

package main
import "fmt"

func main() {
	//%v
	fmt.Printf("I am %v years old\n", 21)
	fmt.Printf("I am %v years old","way too much\n")
	//%d
	fmt.Printf("I have %d choclates\n", 30)
	//%s
	fmt.Printf("I am %s very much\n","happy")
	//%f
	fmt.Printf("Value of PI: %f\n", 3.14)
	fmt.Printf("The number is %.2f\n", 19.54321)
	//%T
	fmt.Printf("The type of the variable is %T\n", true)

	const name = "Ronak Maheshwari"
	const openRate = 30.5

	var msg string = fmt.Sprintf("Hii %s, Your open rate is %.1f percent",name,openRate)
	fmt.Println(msg)
}