package main

import "fmt"

func main() {
	goku := &Saiyan{"Goku", 9001, ""}
	goku.Super()
	fmt.Println(goku.Power)
	fmt.Println(goku.Greeting)
	goku.Hello()
	fmt.Println(goku.Greeting)
}

type Saiyan struct {
	Name     string
	Power    int
	Greeting string
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func (s *Saiyan) Hello() {
	s.Greeting = "Hello"
}
