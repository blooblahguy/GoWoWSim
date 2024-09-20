package structs

type Stats struct {
	Strength  float64
	Agility   float64
	Stamina   int
	Intellect int
	Armor     int

	MeleeHaste       float64
	MeleeCrit        float64
	AttackPower      float64
	MeleeHit         float64
	ArmorPenetration float64
	Expertise        float64

	Haste           float64
	HitChance       float64
	ExpertiseChance float64
	CritChance      float64
}
