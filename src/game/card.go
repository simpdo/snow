package main

import (
	"sort"
)

const (
	Spade   = iota + 1 //黑桃
	Heart              //红桃
	Club               //梅花
	Diamond            //方块

	CardBase   = 13
	BlackJoker = 53
	RedJoker   = 54

	//牌型
	CARD_ILLEGAL = iota

	CARD_SINGLE
	CARD_SINGLE_STRAIGHT

	CARD_PAIR
	CARD_PAIR_STRAIGHT

	CARD_THREE
	CARD_THREE_WITH_SINGLE
	CARD_THREE_WITH_PAIR

	CARD_PLANE
	CARD_PLANE_WITH_SINGLE
	CARD_PLANE_WITH_PAIR

	CARD_BOMB
	CARD_BOMB_WITH_SINGLE
	CARD_BOMB_WITH_PAIR

	CARD_ROCKET
)

type GameCard struct {
	Point int8
	Shape int8
}

func NewCard(value int8) *GameCard {
	if value == BlackJoker || value == RedJoker {
		return &GameCard{
			Point: value,
			Shape: 0,
		}
	}

	point := value % CardBase
	shape := value/CardBase + 1
	if point == 0 {
		point = CardBase
		shape -= 1
	}

	return &GameCard{
		Point: point,
		Shape: shape,
	}
}

func (this *GameCard) Value() int8 {
	if this.Point == BlackJoker || this.Point == RedJoker {
		return this.Point
	}

	return int8(CardBase*this.Shape + CardBase)
}

func (this *GameCard) Equal(other GameCard) bool {
	return this.Point == other.Point
}

/**********************************************************************
=======================================================================
**********************************************************************/
type GameCardSet []GameCard

func (cards GameCardSet) Len() int {
	return len(cards)
}

func (cards GameCardSet) Swap(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}

func (cards GameCardSet) Less(i, j int) bool {
	first := cards[i].Point
	if cards[i].Point == 1 || cards[i].Point == 2 {
		first = first + CardBase
	}

	second := cards[j].Point
	if cards[j].Point == 1 || cards[j].Point == 2 {
		second = second + CardBase
	}

	if cards[i].Point == cards[j].Point {
		return cards[i].Shape < cards[j].Shape
	}

	return first < second
}

func (cards GameCardSet) Sort() {
	sort.Sort(cards)
}

func (cards GameCardSet) Type() int {
	cards.Sort()

	switch cards.Len() {
	case 1:
		return CARD_SINGLE
	case 2:
		if cards.isPair() {
			return CARD_PAIR
		} else if cards.isRocket() {
			return CARD_ROCKET
		}
	case 3:
		if cards.isThree() {
			return CARD_THREE
		}
	case 4:
		if cards.isBomb() {
			return CARD_BOMB
		}
	}

	type2cards := cards.Separate()
	if type2cards[]

	return CARD_ILLEGAL
}

////////////////////////////////////////////////////////
//以下函数调用前调用sort
func (cards GameCardSet) isPair() bool {
	if cards.Len() != 2 {
		return false
	}

	arr := []GameCard(cards)
	if arr[0].Point != arr[1].Point {
		return false
	}

	return true
}

func (cards GameCardSet) isRocket() bool {
	if cards.Len() != 2 {
		return false
	}

	arr := []GameCard(cards)
	if arr[0].Point == BlackJoker && arr[1].Point == RedJoker {
		return true
	}

	return false
}

func (cards GameCardSet) isThree() bool {
	if cards.Len() != 3 {
		return false
	}

	arr := []GameCard(cards)
	if arr[0].Point == arr[2].Point {
		return true
	}

	return false
}

func (cards GameCardSet) isBomb() bool {
	if cards.Len() != 4 {
		return false
	}

	arr := []GameCard(cards)
	if arr[0].Point == arr[3].Point {
		return true
	}

	return false
}

func (cards GameCardSet) Separate() map[int][]GameCardSet {
	type2cards := make(map[int][]GameCardSet)

	arr := []GameCard(cards)
	tmp := arr[0:1]
	for i := 1; i < cards.Len(); i++ {
		if arr[i].Point == arr[i-1].Point {
			tmp = append(tmp, arr[i])
			continue
		}

		switch len(tmp) {
		case 1:
			type2cards[CARD_SINGLE] = append(type2cards[1], tmp)
		case 2:
			type2cards[CARD_PAIR] = append(type2cards[2], tmp)
		case 3:
			type2cards[CARD_THREE] = append(type2cards[3], tmp)
		case 4:
			type2cards[CARD_BOMB] = append(type2cards[4], tmp)
		}
		tmp = tmp[0:0]
	}

	return type2cards
}
