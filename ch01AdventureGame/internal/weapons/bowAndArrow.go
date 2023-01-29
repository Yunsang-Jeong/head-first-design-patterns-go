package weapons

import "fmt"

type bowAndArrow struct{}

func (w *bowAndArrow) UseWeapon() {
	fmt.Println("cutting with bowAndArrow")
}

func (w *bowAndArrow) GetWeaponName() string {
	return "bowAndArrow"
}

func NewBowAndArrow() *bowAndArrow {
	return &bowAndArrow{}
}
