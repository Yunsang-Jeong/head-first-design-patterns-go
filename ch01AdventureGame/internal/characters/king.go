package characters

import (
	"head-first-design-patterns-go/ch01AdventureGame/internal/weapons"
)

func NewKing() *character {
	return &character{
		characterType:    "king",
		weaponeBehaviour: weapons.NewSword(),
	}
}
