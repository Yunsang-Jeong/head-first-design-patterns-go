package weapons

import "fmt"

type sword struct{}

func (w *sword) UseWeapon() {
	fmt.Println("cutting with sword")
}

func (w *sword) GetWeaponName() string {
	return "sword"
}

func NewSword() *sword {
	return &sword{}
}
