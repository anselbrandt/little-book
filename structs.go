package main

import "fmt"

func main() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	fmt.Printf("%d\n", goku.Power)
}

type Saiyan struct {
	Name  string
	Power int
}
