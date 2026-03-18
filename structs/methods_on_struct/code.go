package main

import "fmt"

type authorizationInfo struct {
	username string
	password string
}

func (a authorizationInfo) getAuthDetails() string {
	b := fmt.Sprintf("Authorization Basic %s:%s", a.username, a.password)
	return b
}

func test(a authorizationInfo){
	fmt.Println(a.getAuthDetails())
}

func main() {

}