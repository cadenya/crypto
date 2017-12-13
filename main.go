package main

import "fmt"

func main() {
	var input string
	fmt.Printf("Enter Ticker Symbol:")
	fmt.Scanln(&input)
	fmt.Printf("You entered: %s!", input)
}
