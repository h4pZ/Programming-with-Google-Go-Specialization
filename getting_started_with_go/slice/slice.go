package main

import "fmt"
import "sort"
import "strconv"

func main() {
	// Main variables.
	var input string
	int_store := make([]int, 3)

	for {
		fmt.Print("Please input an integer or X to exit: ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		// Checking whether the input is a number or not.
		num, err := strconv.Atoi(input)

		if err == nil {
			int_store = append(int_store, num)
			sort.Ints(int_store)
			fmt.Println(int_store)
		} else {
			fmt.Printf("[ERROR] The input %s is not a integer", input)
			fmt.Println("")
		}
		fmt.Println("")
	}

}
