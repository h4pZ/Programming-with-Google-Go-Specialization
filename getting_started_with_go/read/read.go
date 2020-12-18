package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAX_LENGTH int = 20

type Name struct {
	fname string
	lname string
}

func (name *Name) trim(s string, length int) string {
	return string([]rune(s)[:length])
}

func (name *Name) Set(fname string, lname string) {
	name.fname = name.trim(fname, MAX_LENGTH)
	name.lname = name.trim(lname, MAX_LENGTH)
}

func main() {
	var file_path string
	var name Name
	names := make([]Name, 0)

	fmt.Print("Please enter the file path: ")
	fmt.Scan(&file_path)

	// Opening the file and reading it.
	file, err := os.Open(file_path)
	defer file.Close()

	if err != nil {
		fmt.Printf("There was a problem reading %s\n", file_path)
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line_split := strings.Split(scanner.Text(), " ")
		name.Set(line_split[0], line_split[1])
		names = append(names, name)
	}

	fmt.Println("\nNames on the file:")
	for _, v := range names {
		fmt.Println(v.fname, v.lname)
	}

}
