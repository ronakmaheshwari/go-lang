/* 
Concatenating strings
Two strings can be concatenated with the + operator. Because Go is strongly typed, it won't allow you to concatenate a string variable with a numeric variable.

Assignment
We'll be using simple basic authentication for the Textio API. This is how our users will communicate to us who they are and how many features they are paying for with each request to our API.
The code on the right has a type error. Change the type of password to a string (but use the same numeric value) so that it can be concatenated with the username variable.
*/
package main
import "fmt"

func main() {
	var username string = "Ronak"
	var password string = "123456"

	fmt.Println("Authorization: Basic", username+":"+password)
}