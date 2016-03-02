// game project main.go
package main

import (
	"fmt"
	"sort"
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

	sort.Sort(GameCardSet(arr))
	fmt.Println(arr)
}
