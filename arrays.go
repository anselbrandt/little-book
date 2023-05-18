package main

import "fmt"

func main() {
	var scores [10]int
	scores[0] = 339
	fmt.Printf("%v\n", scores)
	for index, value := range scores {
		fmt.Printf("index %v value %v\n", index, value)
	}
}
