package ducks

import (
	"head-first-design-patterns-go/ch01SimUDuck/internal/flyingBehaviours"
	"head-first-design-patterns-go/ch01SimUDuck/internal/quackBehaviours"
)

func NewRedHeadDuck() *duck {
	return &duck{
		duckType:        "RedHeadDuck",
		quackBehaviour:  quackBehaviours.NewJustQuack("redhead"),
		flyingBehaviour: flyingBehaviours.NewFlyWithWings(),
	}
}
