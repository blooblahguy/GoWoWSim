package core

type Target struct {
	name   string
	level  byte
	armor  int
	Health float64
}

func NewTarget() Target {
	target := Target{
		name:   "boss",
		level:  73,
		armor:  7700,
		Health: 100,
	}

	target.remove_armor(2600) // Sunder Armor
	target.remove_armor(610)  // Faerie Fire
	target.remove_armor(800)  // Curse of Recklessness

	return target
}

func (target *Target) remove_armor(armor int) {
	target.armor -= armor
}

func (target *Target) add_armor(armor int) {
	target.armor += armor
}
