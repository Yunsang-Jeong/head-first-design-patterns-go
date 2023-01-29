package characters

import (
	"fmt"
	"head-first-design-patterns-go/ch01AdventureGame/internal/weapons"
)

type character struct {
	characterType    string
	weaponeBehaviour weapons.WeaponBehaviour
}

func (c *character) Fight() {
	fmt.Printf(
		"[%s] Try to fight (%s)\n",
		c.characterType,
		c.weaponeBehaviour.GetWeaponName(),
	)
	c.weaponeBehaviour.UseWeapon()
}

func (c *character) SetWeapon(wb weapons.WeaponBehaviour) {
	fmt.Printf(
		"[%s] Change weapon (%s to %s)\n",
		c.characterType,
		c.weaponeBehaviour.GetWeaponName(),
		wb.GetWeaponName(),
	)
	c.weaponeBehaviour = wb
}
