package main

import (
	"fmt"
	"glade/WoWSim/core"
)

type Thread struct {
	Sim       *core.Simulator
	Player    *core.Player
	Target    *core.Target
	Abilities []core.Ability
	Hittable  *core.HitTable
}

func time_until_execute(combat_time float64, target *core.Target) float64 {
	if target.Health < 20 {
		return 0
	}
	if Settings.ExecuteRange == 0 {
		return 0
	}
	return Settings.Duration*(1-Settings.ExecuteRange/100) - combat_time
}

type AbilityLog struct {
	name     string
	casts    int
	hitTypes map[string]int
	damage   float64
	min      float64
	max      float64
}

type SimInstance struct {
	name            string
	overcapped_rage int
	combat_time     float64
	CombatLog       map[string]*AbilityLog

	player    *core.Player
	target    *core.Target
	abilities []core.Ability
	hittable  *core.HitTable
}

func (sim *SimInstance) CreateLogs() {
	sim.CombatLog = make(map[string]*AbilityLog)

	for _, ability := range sim.abilities {
		sim.CombatLog[ability.Name] = &AbilityLog{
			name:     ability.Name,
			hitTypes: make(map[string]int),
			damage:   0,
			casts:    0,
			min:      100000,
			max:      0,
		}
	}

	// for misc attacks
	weapons := []string{"mainhand", "offhand"}
	for _, name := range weapons {
		sim.CombatLog[name] = &AbilityLog{
			name:     name,
			hitTypes: make(map[string]int),
			damage:   0,
			casts:    0,
			min:      100000,
			max:      0,
		}
	}
}

func (sim *SimInstance) Log(name string, dmg float64, hitType string) {
	if hitType == "" {
		return
	}

	log := sim.CombatLog[name]
	// fmt.Printf("Address of %s:\t%p\n", name, &log)

	log.hitTypes[hitType] += 1
	log.damage += dmg
	log.casts += 1
	if dmg > log.max {
		log.max = dmg
	}
	if dmg > 0 && dmg < log.min {
		log.min = dmg
	}
}

func (sim *SimInstance) get_ability_cooldown(abilityname string) float64 {
	for k, v := range sim.abilities {
		if v.Name == abilityname {
			return sim.abilities[k].RemainingCooldown(sim.combat_time)
		}
	}

	return -1
}

func (sim *SimInstance) swing(hand string) {
	dmg, hitType := sim.player.Swing(hand, sim.combat_time, false)

	sim.Log(hand, dmg, hitType)
}

func (sim *SimInstance) cast(abilityname string) {
	// search first by string
	for k, v := range sim.abilities {
		if v.Name == abilityname {
			dmg, hitType := sim.player.Cast(&sim.abilities[k], sim.combat_time)

			sim.Log(abilityname, dmg, hitType)
			return
		}
	}

	fmt.Println("ERROR: tried to cast ability that doesn't exist ", abilityname)
}
