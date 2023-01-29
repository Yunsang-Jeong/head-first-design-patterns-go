package quackBehaviours

import (
	"fmt"
)

type justQuack struct {
	sound string
}

func (q *justQuack) Quack() {
	fmt.Printf("Try to just quack: %s\n", q.sound)
}

func NewJustQuack(sound string) *justQuack {
	return &justQuack{
		sound: sound,
	}
}
