package main
import "fmt"

type car struct {
	make string
	model string
}

type truck struct {
	car
	bedSize int
}

func truckStruct() {
	cartype := truck{
		bedSize: 15,
		car: car{
			make: "Toyota",
			model: "Sevron",
		},
	}
	fmt.Println("The molel of truck is ",cartype.model);
	fmt.Printf("The truck has bedsize: %d", cartype.bedSize);
}
