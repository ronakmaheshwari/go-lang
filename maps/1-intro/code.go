package main
import "fmt"

type user struct {
	name string
	phoneNumber int
}

func getUserMap (name []string, phoneNumber []int) (map[string]user , error){
	userMap := make(map[string]user)
	if len(name) != len(phoneNumber) {
		return map[string]user{}, fmt.Errorf("Invalid Slices")
	}

	for i:=0;i<len(name);i++ {
		a := name[i]
		b := phoneNumber[i]
		userMap[a] = user{
				name: a,
				phoneNumber: b,
		}
		
	}
	return userMap, nil
}

func test(name []string,phoneNumber []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, error := getUserMap(name,phoneNumber)
	if error != nil {
		fmt.Println(error)
		return
	}
	for _, name := range name {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Mary", "Alpha"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}