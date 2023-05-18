package main

import (
	"fmt"
)

func main() {
	var power int
	power = 9000
	fmt.Printf("It's over %d\n", power)

	// declaration operator infers type
	morePower := 9001
	fmt.Printf("It's %d\n", morePower)

	// power can be reassigned, but type can't be changed
	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
}
