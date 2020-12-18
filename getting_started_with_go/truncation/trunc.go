package main

import "fmt"

func main() {
	var input_num float32
	fmt.Printf("Enter a number (float): ")
	fmt.Scan(&input_num)
	fmt.Printf("The truncated number is: %d", int32(input_num))
}
