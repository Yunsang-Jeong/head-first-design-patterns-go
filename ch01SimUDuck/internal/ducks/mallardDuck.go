package ducks

import (
	"head-first-design-patterns-go/ch01SimUDuck/internal/flyingBehaviours"
	"head-first-design-patterns-go/ch01SimUDuck/internal/quackBehaviours"
)

func NewMallardDuck() *duck {
	return &duck{
		duckType:        "MallardDuck",
		quackBehaviour:  quackBehaviours.NewJustQuack("mallard"),
		flyingBehaviour: flyingBehaviours.NewFlyWithWings(),
	}
}
