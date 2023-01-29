package weapons

import "fmt"

type axe struct{}

func (w *axe) UseWeapon() {
	fmt.Println("cutting with axe")
}

func (w *axe) GetWeaponName() string {
	return "axe"
}

func NewAxe() *axe {
	return &axe{}
}
