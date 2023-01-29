package flyingBehaviours

import "fmt"

type flyWithWings struct{}

func (q *flyWithWings) Fly() {
	fmt.Println("Try to just Fly")
}

func NewFlyWithWings() *flyWithWings {
	return &flyWithWings{}
}
