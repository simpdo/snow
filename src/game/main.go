// game project main.go
package main

import (
	"fmt"
)

func main() {

	arr := []GameCard{
		{4, Heart},
		{3, Club},
		{3, Spade},
		{4, Heart},
		{3, Diamond},
		{4, Heart},
		{4, Heart},
		{5, Heart},
	}

	card_set := GameCardSet(arr)
	card_set.Sort()
	fmt.Println(card_set)
	fmt.Println(card_set.Type())
}
