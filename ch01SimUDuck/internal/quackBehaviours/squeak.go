package quackBehaviours

import (
	"fmt"
)

type squeak struct {
	sound string
}

func (q *squeak) Quack() {
	fmt.Printf("Try to just quack: %s\n", q.sound)
}

func NewSqueak(sound string) *squeak {
	return &squeak{
		sound: sound,
	}
}
