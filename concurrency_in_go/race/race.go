package main

import (
	"fmt"
	"time"
)

/*
A Race condition is when the outcome of a program depends
on the execution sequence of itself. In this case, a race condition
happens because both go race coroutines access the same shared memory (count)
and thus, the execution time will impact on the results of the program.
*/

var count int

func race(count *int, thing string) {
	*count++
	fmt.Println(thing, *count)
}

func main() {
	go race(&count, "First")
	go race(&count, "Second")
	time.Sleep(1 * time.Second)
}
