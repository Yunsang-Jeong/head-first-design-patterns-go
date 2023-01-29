package main

import (
	"head-first-design-patterns-go/ch01AdventureGame/internal/characters"
	"head-first-design-patterns-go/ch01AdventureGame/internal/weapons"
)

func main() {
	k := characters.NewKing()
	k.Fight()

	k.SetWeapon(weapons.NewAxe())
	k.Fight()

	k.SetWeapon(weapons.NewBowAndArrow())
	k.Fight()

	k.SetWeapon(weapons.NewKnife())
	k.Fight()
}
