package main

import "fmt"

func main() {
	scoresImpl := []int{1, 4, 293, 4, 9}
	fmt.Printf("%v\n", scoresImpl)
	scoresExpl := make([]int, 0, 10)
	scoresExpl = append(scoresExpl, 5)
	fmt.Println(scoresExpl)
	scores := []int{1, 2, 3, 4, 5}
	fmt.Println(scores)
	slice := scores[2:4]
	slice[0] = 999
	fmt.Println(scores)
}
