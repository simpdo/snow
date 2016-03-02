package main

import ()

const (
	Spade   = iota //黑桃
	Heart          //红桃
	Club           //梅花
	Diamond        //方块
)

type GameCard struct {
	Value int8
	Shape int
}

type GameCardSet []GameCard

func (arr GameCardSet) Len() int {
	return len(arr)
}

func (arr GameCardSet) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr GameCardSet) Less(i, j int) bool {
	if arr[i].Value == arr[j].Value {
		return arr[i].Shape < arr[j].Shape
	}
	return arr[i].Value < arr[j].Value
}
