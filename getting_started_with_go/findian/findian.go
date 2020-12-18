package main

import "os"
import "bufio"
import "fmt"
import "strings"

func main() {
	var text string
	scanner := bufio.NewScanner(os.Stdin)

	// Getting the user input.
	fmt.Printf("Please input your text: ")
	scanner.Scan()
	text = strings.ToLower(scanner.Text())

	// Check for ian.
	if strings.HasPrefix(text, "i") &&
		strings.Contains(text, "a") &&
		strings.HasSuffix(text, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
