package ducks

import (
	"head-first-design-patterns-go/ch01SimUDuck/internal/flyingBehaviours"
	"head-first-design-patterns-go/ch01SimUDuck/internal/quackBehaviours"
)

func NewDecoyDuck() *duck {
	return &duck{
		duckType:        "DecoyDuck",
		quackBehaviour:  quackBehaviours.NewMuteQuack(),
		flyingBehaviour: flyingBehaviours.NewflyNoWay(),
	}
}
