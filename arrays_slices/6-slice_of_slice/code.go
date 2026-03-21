package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	arr := make([][]int, rows)
	for i:=0;i<rows;i++{
		arr[i] = make([]int, cols)
		for j:=0;j<cols;j++{
			arr[i][j] = i*j
		}
	}
	return arr
}

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	fmt.Println("The length of matrix:", len(matrix))
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}