package main

import "fmt"

func main() {
	goku := &Saiyan{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power)
}

func Super(s *Saiyan) {
	s.Power += 10000
}

type Saiyan struct {
	Name  string
	Power int
}
