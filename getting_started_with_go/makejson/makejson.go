package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	// Getting the name and address.
	fmt.Print("Please write your name: ")
	scanner.Scan()
	data["name"] = scanner.Text()

	fmt.Print("Please write your address: ")
	scanner.Scan()
	data["address"] = scanner.Text()

	json_data, _ := json.MarshalIndent(data, "", "    ")

	fmt.Println("Json:")
	fmt.Println(string(json_data))
	fmt.Println("")
	fmt.Println("Json byte array:")
	fmt.Println(json_data)

}
