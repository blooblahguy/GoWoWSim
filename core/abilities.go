package core

// struct callback function
type AbilityCallback func(player *Player, target *Target) float64

// main ability struct
type Ability struct {
	Name     string
	callback AbilityCallback
	cooldown float64
	cost     int

	last_used   float64
	active      bool
	triggersGCD bool
	NextCast    float64
}

// maker function with some default values
func NewAbility(name string, callback AbilityCallback, cost int, cooldown float64) Ability {
	ability := Ability{
		Name:        name,
		callback:    callback,
		cost:        cost,
		cooldown:    cooldown,
		last_used:   0,
		active:      false,
		triggersGCD: false,
		NextCast:    0,
	}

	return ability
}

// check if we can cast
func (ability *Ability) CanCast(combat_time float64) bool {
	if combat_time >= ability.NextCast || ability.NextCast == 0 {
		return true
	} else {
		return false
	}
}

// get remaining cooldown
func (ability *Ability) RemainingCooldown(combat_time float64) float64 {
	return ability.NextCast - combat_time
}

// cast the ability and return callback damage
func (ability *Ability) Cast(player *Player, target *Target, combat_time float64) float64 {
	ability.NextCast = float64(combat_time + ability.cooldown)
	return ability.callback(player, target)
}

// #heroic strike
func heroic_strike(player *Player, target *Target) float64 {
	damage := player.normalize_swing(player.Mainhand, 0)
	damage += 176
	return damage
}

// bloodthirst
func bloodthirst(player *Player, target *Target) float64 {
	var dmg = (player.Stats.AttackPower * 45) / 100
	return float64(dmg)
}

// whirlwind
func whirlwind(player *Player, target *Target) float64 {
	mh_damage := float64(randint(player.Mainhand.MinDamage, player.Mainhand.MaxDamage))
	oh_damage := float64((randint(player.Offhand.MinDamage, player.Offhand.MinDamage))) //+ player.stats["oh_damage"])

	oh_damage *= float64(float64(1) + float64(0.05)*float64(player.talents.dual_wield_specialization)) // add talent damage
	oh_damage /= 2                                                                                     // offhand penalty

	mh_damage = mh_damage + (float64(player.Stats.AttackPower) * (player.Mainhand.BaseSwingSpeed / 14))
	oh_damage = oh_damage + (float64(player.Stats.AttackPower) * (player.Offhand.BaseSwingSpeed / 14))

	return mh_damage + oh_damage
}

// execute
func execute(player *Player, target *Target) float64 {
	damage := float64(925 + (player.Rage-15)*21)
	player.spend_rage(player.Rage) // zero out rage
	return damage
}

func NewAbilities(player *Player) []Ability {
	// holder arrays
	Abilities := []Ability{
		NewAbility("bloodthirst", bloodthirst, 30, 6),
		NewAbility("whirlwind", whirlwind, 25, 10-(1*float64(player.talents.improved_whirlwind))),
		NewAbility("execute", bloodthirst, 10, 0),
		NewAbility("heroic_strike", heroic_strike, 15-(1*player.talents.improved_heroic_strike), 0),
	}

	return Abilities
}
