package characters

import (
	"head-first-design-patterns-go/ch01AdventureGame/internal/weapons"
)

func NewQueen() *character {
	return &character{
		characterType:    "queen",
		weaponeBehaviour: weapons.NewBowAndArrow(),
	}
}
