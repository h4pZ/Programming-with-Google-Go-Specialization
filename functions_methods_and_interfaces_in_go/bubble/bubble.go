package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TextToIntSlice(text string) []int {
	sequence := make([]int, 0)
	integers := strings.Fields(text)

	if len(integers) > 10 {
		err := "[ERROR] The max numbers of integers is 10. You inputed %d\n"
		fmt.Printf(err, len(integers))
		os.Exit(1)
	}

	for _, entry := range integers {
		if number, err := strconv.Atoi(entry); err == nil {
			sequence = append(sequence, number)
		} else {
			fmt.Println("There was an error with the input")
			fmt.Println("[ERROR]", err)
			os.Exit(1)
		}
	}
	return sequence
}

func Swap(slice []int, idx int) {
	next_value := slice[idx+1]
	slice[idx+1] = slice[idx]
	slice[idx] = next_value
}

func BubbleSort(slice []int) {
	for n := len(slice); n > 1; n-- {
		for i := 0; i < n-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
			}
		}
	}
}

func main() {
	var user_input string
	scanner := bufio.NewScanner(os.Stdin)

	// Getting the user input.
	fmt.Print("Please input a list of integers (max 10 integers): ")
	scanner.Scan()
	user_input = scanner.Text()

	// Create the squence and sort it.
	sequence := TextToIntSlice(user_input)
	BubbleSort(sequence)
	fmt.Println("The sorted list of integers is: ", sequence)

}
