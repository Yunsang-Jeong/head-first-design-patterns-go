package main

import (
	"head-first-design-patterns-go/ch01SimUDuck/internal/ducks"
	"head-first-design-patterns-go/ch01SimUDuck/internal/flyingBehaviours"
	"head-first-design-patterns-go/ch01SimUDuck/internal/quackBehaviours"
)

func main() {
	d := ducks.NewMallardDuck()
	d.PerformFly()
	d.PerformQuack()

	d.SetQuackBehaviour(quackBehaviours.NewMuteQuack())
	d.PerformQuack()

	d.SetFlyingBehaviour(flyingBehaviours.NewflyNoWay())
	d.PerformFly()
}
