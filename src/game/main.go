// game project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	arr := []GameCard{
		{2, Heart},
		{1, Club},
		{6, Spade},
		{3, Diamond},
		{3, Spade},
	}

	card_set := GameCardSet(arr)
	card_set.Sort()
	fmt.Println(arr)
	arr = arr[0:0]
	fmt.Println(arr)

	ha := make(map[int]int)
	ha[1] = 1
	ha[2] = 1
	fmt.Println(len(ha))
}
