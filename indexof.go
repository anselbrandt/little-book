package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "the spice must flow"
	res := strings.Index(haystack[5:], " ")
	fmt.Printf("%v\n", res)
	scores := []int{1, 2, 3, 4, 5}
	scores = scores[:len(scores)-1]
	fmt.Printf("%v\n", scores)
}
