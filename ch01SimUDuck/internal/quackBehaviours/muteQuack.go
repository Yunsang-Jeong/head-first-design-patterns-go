package quackBehaviours

import "fmt"

type muteQuack struct{}

func (q *muteQuack) Quack() {
	fmt.Println("Try to mute quack")
}

func NewMuteQuack() *muteQuack {
	return &muteQuack{}
}
