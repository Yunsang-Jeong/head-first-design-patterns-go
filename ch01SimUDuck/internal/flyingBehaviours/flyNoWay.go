package flyingBehaviours

import "fmt"

type flyNoWay struct{}

func (q *flyNoWay) Fly() {
	fmt.Println("Try to fly without wings")
}

func NewflyNoWay() *flyNoWay {
	return &flyNoWay{}
}
