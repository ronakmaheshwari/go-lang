package main
import "fmt"

type user struct {
	name string
	number int64
}

type sender struct {
	user
	ratelimit int
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.ratelimit)
	fmt.Println("====================================")
}

func main() {
	test(sender{
		user: user{
			name: "Ronak Maheshwari",
			number: 9545293102,
		},
		ratelimit: 5,
	})
	test(sender{
		user: user{
			name: "Jaby Koey",
			number: 1213342009,
		},
		ratelimit: 10,
	})

}