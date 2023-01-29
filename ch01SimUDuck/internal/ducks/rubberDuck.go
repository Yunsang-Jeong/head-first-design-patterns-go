package ducks

import (
	"head-first-design-patterns-go/ch01SimUDuck/internal/flyingBehaviours"
	"head-first-design-patterns-go/ch01SimUDuck/internal/quackBehaviours"
)

func NewRubberDuck() *duck {
	return &duck{
		duckType:        "RubberDuck",
		quackBehaviour:  quackBehaviours.NewJustQuack("rubber"),
		flyingBehaviour: flyingBehaviours.NewflyNoWay(),
	}
}
