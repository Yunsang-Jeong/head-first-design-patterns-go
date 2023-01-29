package characters

import (
	"head-first-design-patterns-go/ch01AdventureGame/internal/weapons"
)

func NewTroll() *character {
	return &character{
		characterType:    "troll",
		weaponeBehaviour: weapons.NewAxe(),
	}
}
