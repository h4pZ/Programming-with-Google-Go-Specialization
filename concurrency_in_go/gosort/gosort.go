package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n_workers int = 4
	sorted_array := make([]int, 0)
	jobs := make(chan []int, n_workers)
	results := make(chan []int, n_workers)

	// Populating the jobs.
	user_array := GetArray()
	var batch_size int = len(user_array) / n_workers

	for i := 0; i < n_workers; i++ {
		if i < n_workers-1 {
			slice := user_array[i*batch_size : (1+i)*batch_size]
			jobs <- slice
		} else {
			slice := user_array[i*batch_size:]
			jobs <- slice
		}
	}
	close(jobs)

	// Fetching the jobs and sending out the workers.
	for i := 0; i < n_workers; i++ {
		job := <-jobs

		go func(job []int, idx int, results chan<- []int) {
			fmt.Printf("Goroutine idx: %d | Array: %v\n", idx, job)
			sort.Ints(job)
			fmt.Printf("Goroutine idx: %d | Sorted array: %v\n", idx, job)
			results <- job
		}(job, i, results)
	}

	// Getting the results and sorting them out.
	for i := 0; i < n_workers; i++ {
		sorted_array = append(sorted_array, <-results...)
	}

	sort.Ints(sorted_array)
	fmt.Printf("The sorted array is: %v\n", sorted_array)
}

func TextToIntSlice(text string) []int {
	sequence := make([]int, 0)
	integers := strings.Fields(text)

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

func GetArray() []int {
	scanner := bufio.NewScanner(os.Stdin)
	var valid_array bool = false
	var sequence []int

	for valid_array == false {
		// Getting the user input.
		fmt.Print("Please input a list of integers (min 4 integers): ")
		scanner.Scan()
		user_input := scanner.Text()

		// Create the sequence.
		sequence = TextToIntSlice(user_input)

		if len(sequence) >= 4 {
			valid_array = true
		} else {
			fmt.Printf("The length of the array should be greater than 4 not %d\n", len(sequence))
		}
	}

	return sequence
}
