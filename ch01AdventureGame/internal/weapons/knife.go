package weapons

import "fmt"

type knife struct{}

func (w *knife) UseWeapon() {
	fmt.Println("cutting with knife")
}

func (w *knife) GetWeaponName() string {
	return "knife"
}

func NewKnife() *knife {
	return &knife{}
}
