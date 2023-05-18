package main

func main() {

}

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

type Saiyan struct {
	Name  string
	Power int
}

/*

goku := new(Saiyan)

same as

goku := &Saiyan{}

*/
