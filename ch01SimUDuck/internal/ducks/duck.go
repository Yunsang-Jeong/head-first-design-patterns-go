package ducks

import (
	"fmt"
	"head-first-design-patterns-go/ch01SimUDuck/internal/flyingBehaviours"
	"head-first-design-patterns-go/ch01SimUDuck/internal/quackBehaviours"
)

type duck struct {
	duckType        string
	quackBehaviour  quackBehaviours.QuackBehaviour
	flyingBehaviour flyingBehaviours.FlyingBehaviour
}

func (d *duck) Display() {
	fmt.Printf("This is %s\n", d.duckType)
}

func (d *duck) PerformQuack() {
	d.quackBehaviour.Quack()
}

func (d *duck) PerformFly() {
	d.flyingBehaviour.Fly()
}

func (d *duck) SetQuackBehaviour(qb quackBehaviours.QuackBehaviour) {
	d.quackBehaviour = qb
}

func (d *duck) SetFlyingBehaviour(fb flyingBehaviours.FlyingBehaviour) {
	d.flyingBehaviour = fb
}
