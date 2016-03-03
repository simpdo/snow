package main

import ()

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
	CARD_PAIR
	CARD_ROCKET
	CARD_THREE
	CARD_THREE_WITH_SINGLE
	CARD_THREE_WITH_PAIR
	CARD_BOMB
	CARD_BOMB_WITH_SINGLE
	CARD_BOMB_WITH_PAIR
	CARD_STRAIGHT
	CARD_STRAIGHT_PAIR

	CARD_PLANE
	CARD_PLANE_WITH_SINGLE
	CARD_PLANE_WITH_PAIR
)

type GameCard struct {
	Point int8
	Shape int8
}

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

func NewCard(Value int8) *GameCard {
	if Value < 1 || Value > RedJoker {
		return nil
	}

	if Value == BlackJoker || Value == RedJoker {
		return &GameCard{
			Point: Value,
			Shape: 0,
		}
	}

	return &GameCard{
		Point: int8(Value % CardBase),
		Shape: int8(Value / CardBase),
	}
}

func (this *GameCard) Value() int8 {
	if this.Point == BlackJoker || this.Point == RedJoker {
		return this.Point
	}

	return int8(CardBase*(this.Shape-1) + CardBase)
}

func (this *GameCard) Type() int {
	return 0
}
