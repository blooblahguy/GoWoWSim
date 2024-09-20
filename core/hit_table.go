package core

type Hit struct {
	damage   float64
	hit_type string
}

type HitTable struct {
	miss      float64    // dual wielding
	dodge     float64    // from the back
	glance    float64    // flat chance for 73
	glance_dr float64    // flat reduction
	crit      float64    // filler stat
	modifier  [2]float64 // calculate what a diminished or amplified attack should look like
	Player    *Player    // set these as pointers to reduce func params
	Target    *Target    // set these as pointers to reduce func params
}

func NewHitTable(player *Player, target *Target) HitTable {
	hittable := HitTable{
		miss:      28,
		dodge:     6.5,
		glance:    24,
		glance_dr: 25,
		crit:      0,
		Player:    player,
		Target:    target,
	}

	return hittable
}

func (hittable *HitTable) CalcWhiteHit(damage float64) (float64, string) {
	hittable.glance = 24
	hittable.miss = 28

	hittable.modifier[0] = (100 - hittable.glance_dr) / 100
	hittable.modifier[1] = 1 + (2*(1+hittable.Player.CritMult) - 1)

	// fmt.Println("white crit mult:", hittable.modifier[1])

	return hittable._calc_hit(damage, false)
}

func (hittable *HitTable) CalcSpecialHit(damage float64) (float64, string) {
	hittable.glance = 0
	hittable.miss = 9

	hittable.modifier[0] = 0
	hittable.modifier[1] = 1 + (2*(1+hittable.Player.CritMult)-1)*(1+0.1*float64(hittable.Player.talents.impale))
	// fmt.Println("yellow crit mult:", hittable.modifier[1])

	return hittable._calc_hit(damage, true)
}

func (hittable *HitTable) _calc_hit(damage float64, special bool) (float64, string) {
	// fmt.Println(damage)
	damage *= hittable.Player.DamageMult
	// fmt.Println(damage)

	hittable.miss = hittable.miss - hittable.Player.Stats.HitChance
	hittable.dodge = 6.5 - hittable.Player.Stats.ExpertiseChance
	hittable.crit = hittable.Player.Stats.CritChance

	roll := float64(randint(0, 100))
	calc_miss := hittable.miss
	calc_dodge := hittable.miss + hittable.dodge
	calc_glance := hittable.miss + hittable.dodge + hittable.glance
	calc_crit := hittable.miss + hittable.dodge + hittable.glance + hittable.crit

	// fmt.Println("roll", roll)
	// fmt.Println("miss", calc_miss)
	// fmt.Println("dodge", calc_dodge)
	// fmt.Println("glance", calc_glance)
	// fmt.Println("crit", calc_crit)
	// # if (special):
	// # 	print("miss", calc_miss)
	// # 	print("dodge", calc_dodge, "(", self.dodge, ")")
	// # 	print("glance", calc_glance, "(", self.glance, ")")
	// # 	print("crit", calc_crit, "(", self.crit, ")")
	// # 	print("remainder", 100 - calc_crit)

	// reduce from armor now
	armor := float64(hittable.Target.armor)
	reduce := 1 - armor/(armor-22167.5+467.5*float64(hittable.Player.level))
	damage *= reduce

	// roll
	if roll < calc_miss {
		return 0, "miss"
	}
	if roll < calc_dodge {
		return 0, "dodge"
	}
	if roll < calc_glance {
		return damage * hittable.modifier[0], "glance"
	}
	if roll < calc_crit {
		hittable.Player.ApplyFlurry()
		return damage * hittable.modifier[1], "crit"
	}

	return damage, "hit"
}
