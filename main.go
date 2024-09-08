package main

import (
	"fmt"
)

var score = 9.78 // this shouldn't be inside the main the main function either

func main() {

	sayHello("Noah")

	for _, v := range points {
		fmt.Println("value", v)
	}

	showScore()

}

// calling functions within the same package - go run main.go greetings.go
