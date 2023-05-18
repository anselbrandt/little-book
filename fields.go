package main

import "fmt"

func main() {
	gohan := &Saiyan{
		Name:   "Gohan",
		Power:  1000,
		Father: &Saiyan{Name: "Goku", Power: 9001, Father: nil}}
	fmt.Printf("%v\n", gohan)
}

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}
