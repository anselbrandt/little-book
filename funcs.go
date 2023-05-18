package main

import (
	"fmt"
)

func main() {
	log("hello")
	total := add(1, 2)
	fmt.Printf("%d\n", total)

}

func log(message string) {
	fmt.Printf("%s\n", message)
}

func add(a, b int) int {
	return a + b
}
