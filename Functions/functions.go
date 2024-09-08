package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Hello %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Melanie")
	// sayGreeting("Josh")
	// sayBye("Melanie")

	cycleNames([]string{"Clare", "Deny", "Bosh"}, sayGreeting)
	cycleNames([]string{"Clare", "Deny", "Bosh"}, sayBye)

	a1 := circleArea(10.5)
	fmt.Println("Area of circle", a1)
	fmt.Printf("Area of circle %0.2f\n", a1)

	fmt.Println("Area of circle", circleArea(12))
}
