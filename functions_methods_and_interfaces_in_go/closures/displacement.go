package main

import (
	"fmt"
	"math"
	"os"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}

}

func main() {

	// Displacement parameters.
	var a, v0, s0, t float64

	// Populating the parameters.
	fmt.Print("Please enter the acceleration: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("[ERROR] ", err)
		os.Exit(1)
	}
	fmt.Print("Please enter the initial velocity: ")
	if _, err := fmt.Scan(&v0); err != nil {
		fmt.Println("[ERROR] ", err)
		os.Exit(1)
	}
	fmt.Print("Please enter the initial position: ")
	if _, err := fmt.Scan(&s0); err != nil {
		fmt.Println("[ERROR] ", err)
		os.Exit(1)
	}
	fmt.Print("Please enter the time: ")
	if _, err := fmt.Scan(&t); err != nil {
		fmt.Println("[ERROR] ", err)
		os.Exit(1)
	}

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("The final position will be:", fn(t))

}
