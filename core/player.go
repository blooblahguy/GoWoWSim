package core

import (
	"fmt"
	"glade/WoWSim/items"
	"glade/WoWSim/structs"
	"math"
)

type Weapons struct {
	Mainhand items.Weapon
	Offhand  items.Weapon
}

type Talents struct {
	weapon_mastery            int
	improved_heroic_strike    int
	impale                    int
	cruelty                   int
	commanding_presence       int
	dual_wield_specialization int
	improved_berserker_stance int
	improved_whirlwind        int
}

type Player struct {
	level        int
	Equipment    map[string]items.Item
	Enchantments map[string]items.Enchant
	Mainhand     items.Item
	Offhand      items.Item
	Race         string

	NextSwing []float64
	LastSwing float64
	// next_oh_swing []float64

	Target    *Target
	Hittable  *HitTable
	Abilities *[]Ability

	// # add in all static sources of stats
	CritMult        float64       // kings base
	DamageMult      float64       // kings base
	strength_mult   float64       // kings base
	agility_mult    float64       // kings base
	character_stats structs.Stats // stats when naked

	gear_stats structs.Stats // stats from gear
	buffs      structs.Stats // tracking active buffs and stats
	cooldowns  structs.Stats // tracking procs (active or not) and stats
	Stats      structs.Stats // final calculated stats

	// formula stats
	Rage        int
	rage_factor float64

	// quick talents
	talents Talents

	// combat variables
	QueueHeroicStrike bool
	last_anger_rage   float64
	last_bloodrage    float64
	blood_rage_total  int
	blood_rage_active bool
	flurry            int
	last_wf           float64

	// gcd
	base_gcd      float64
	gcd           float64
	off_gcd       bool
	next_gcd      float64
	last_gcd      float64
	reported_gcd1 bool
	reported_gcd2 bool
}

// add item to the player
func (player *Player) Enchant(enchantname string, slot string) {
	enchant := items.EnchantsByName[enchantname]

	player.Enchantments[slot] = enchant
	// Bonus
}

func (player *Player) Equip(itemname string, slot string) {
	item := items.ByName[itemname]

	// Lets equip some gems first
	// harcoding preference for now
	meta := items.GemsByName["Relentless Earthstorm Diamond"]
	// red := items.GemsByName["Inscribed Pyrestone"]
	// yellow := items.GemsByName["Smooth Lionseye"]
	// blue := items.GemsByName["Jagged Seaspray Emerald"]
	red := items.GemsByName["Inscribed Noble Topaz"]
	yellow := items.GemsByName["Smooth Dawnstone"]
	blue := items.GemsByName["Jagged Talasite"]

	// force array size
	item.Gems = make([]items.Gem, len(item.GemSockets))
	// add gems into item sockets
	for key, id := range item.GemSockets {
		if id == items.GemColorMeta {
			item.Gems[key] = meta
		} else if id == items.GemColorRed {
			item.Gems[key] = red
		} else if id == items.GemColorBlue {
			item.Gems[key] = blue
		} else if id == items.GemColorYellow {
			item.Gems[key] = yellow
		}

	}

	// If these are weapons we're storing them separately
	if slot == "Mainhand" {
		player.Mainhand = item
		player.Mainhand.Hand = "Mainhand"
		return
	} else if slot == "Offhand" {
		player.Offhand = item
		player.Offhand.Hand = "Offhand"
		return
	}

	// Otherwise dump in our equipment string map
	player.Equipment[slot] = item
}

func (player *Player) add_stats(stats structs.Stats) {
	// merge values together here
	player.Stats.Strength += stats.Strength
	player.Stats.Agility += stats.Agility
	player.Stats.MeleeHaste += stats.MeleeHaste
	player.Stats.MeleeCrit += stats.MeleeCrit
	player.Stats.CritChance += stats.CritChance
	player.Stats.AttackPower += stats.AttackPower
	player.Stats.MeleeHit += stats.MeleeHit
	player.Stats.ArmorPenetration += stats.ArmorPenetration
}

func (player *Player) print_stats() {
	fmt.Println("Strength:", math.Round(player.Stats.Strength*100)/100)
	fmt.Println("Agility:", math.Round(player.Stats.Agility*100)/100)
	fmt.Println("MeleeCrit:", math.Round(player.Stats.MeleeCrit*100)/100)
	fmt.Println("	CritChance:", math.Round(player.Stats.CritChance*100)/100)
	fmt.Println("AttackPower:", math.Round(player.Stats.AttackPower*100)/100)
	fmt.Println("MeleeHit:", math.Round(player.Stats.MeleeHit*100)/100)
	fmt.Println("	HitChance:", math.Round(player.Stats.HitChance*100)/100)
	fmt.Println("ArmorPenetration:", math.Round(player.Stats.ArmorPenetration*100)/100)
	fmt.Println("MeleeHaste:", math.Round(player.Stats.MeleeHaste*100)/100)
	fmt.Println("	Haste:", math.Round(player.Stats.Haste*10000)/10000)
	fmt.Println("Expertise:", math.Round(player.Stats.Expertise*100)/100)
	fmt.Println("	ExpertiseChance:", math.Round(player.Stats.ExpertiseChance*100)/100)
}

// start to calculate all stats together
func (player *Player) CalculateStats() {
	// naked stats
	player.add_stats(player.character_stats)

	// gear stats
	for _, item := range player.Equipment {
		player.add_stats(item.Stats)       // item base stats
		player.add_stats(item.SocketBonus) // item socket bonus stats
		for _, gem := range item.Gems {
			player.add_stats(gem.Stats) // socketed gem stats
		}
	}

	// weapon stats
	player.add_stats(player.Mainhand.Stats)       // item base stats
	player.add_stats(player.Mainhand.SocketBonus) // item socket bonus stats
	for _, gem := range player.Mainhand.Gems {
		player.add_stats(gem.Stats) // socketed gem stats
	}
	player.add_stats(player.Offhand.Stats)        // item base stats
	player.add_stats(player.Mainhand.SocketBonus) // item socket bonus stats
	for _, gem := range player.Offhand.Gems {
		player.add_stats(gem.Stats) // socketed gem stats
	}

	// enchants
	for _, enchant := range player.Enchantments {
		player.add_stats(enchant.Bonus) // item socket bonus stats
	}

	// talents

	// buff stats

	// modifers

	// covert values

	// calculate hit
	player.Stats.HitChance = player.Stats.MeleeHit / 15.77
	if player.Stats.HitChance < 0 {
		player.Stats.HitChance = 0
	}

	// calculate expertise
	if player.Race == "human" {
		player.Stats.ExpertiseChance += 1.25
	}
	player.Stats.ExpertiseChance += (player.Stats.Expertise / 3.9423) * 0.25
	player.Stats.ExpertiseChance += float64(player.talents.weapon_mastery)

	if player.Stats.ExpertiseChance > 6.5 {
		player.Stats.ExpertiseChance = 6.5
	}

	// apply kings
	// player.Stats.Strength *= 1.1 //player.strength_mult
	// player.Stats.Agility *= 1.1  //player.agility_mult

	// calculate crit chance
	player.Stats.CritChance += player.Stats.MeleeCrit / 22.08
	player.Stats.CritChance += player.Stats.Agility / 33
	player.Stats.CritChance += float64(player.talents.cruelty) // cruelty talent
	player.Stats.CritChance += 3                               // berserker stance

	// calculate haste rating
	player.Stats.Haste = (player.Stats.MeleeHaste / 15.77) / 100

	// lastly add strength to AP
	player.Stats.AttackPower += player.Stats.Strength * 2

	// imp berserker stance
	player.Stats.AttackPower *= 1 + (.02 * float64(player.talents.improved_berserker_stance))

	// unleashed rage?
	// player.Stats.AttackPower *= 1.1

	// all done
	// player.print_stats()

	// get baseline haste set
	player.update_haste()
}

func (player *Player) gain_rage(rage int) {
	player.Rage += rage
	if player.Rage > 100 {
		player.Rage = 100
	}
}
func (player *Player) spend_rage(rage int) {
	player.Rage -= rage
	if player.Rage < 0 {
		player.Rage = 0
	}
}

func (player *Player) generate_rage(damage float64, weapon items.Item, crit bool) {
	if damage == 0 {
		return
	}

	// # weapon rage factor from mh/oh crit/nocrit
	wep_factor := float64(0)
	if weapon.Hand == "mainhand" {
		if crit {
			wep_factor = 7
		} else {
			wep_factor = 3.5
		}
	} else if weapon.Hand == "offhand" {
		if crit {
			wep_factor = 3.5
		} else {
			wep_factor = 1.75
		}
	}

	// # normalized weapon speed
	wep_speed_normal := float64(0)
	if weapon.WeaponType == items.HandTypeTwoHand {
		wep_speed_normal = 3.3
	} else if weapon.WeaponType == items.WeaponTypeDagger {
		wep_speed_normal = 1.7
	} else { // is normal 1 hander
		wep_speed_normal = 2.4
	}

	// # final rage formula
	rage := int((((damage / player.rage_factor) * 7.5) + (wep_speed_normal * wep_factor)) / 2)

	player.gain_rage(rage)
}

func (player *Player) normalize_swing(weapon items.Item, bonus int) float64 {
	damage := float64(randint(weapon.MinDamage, weapon.MaxDamage) + bonus)

	wep_speed := 2.4 // normalized for all 1h non-daggers
	if weapon.WeaponType == items.WeaponTypeDagger {
		wep_speed = 1.7 // 1.7 for daggers
	} else if weapon.HandType == items.HandTypeTwoHand {
		wep_speed = 3.3 // 3.3 for two-handed weapons
	} else if weapon.Type == items.ItemTypeRanged {
		wep_speed = 2.8 // 2.8 for ranged weapons
	}

	damage += float64((player.Stats.AttackPower)/14) * wep_speed

	if weapon.Hand == "Offhand" {
		damage *= 1 + (.05 * float64(player.talents.dual_wield_specialization)) // talents
		damage /= 2
	}

	return damage
}
func (player *Player) ApplyFlurry() {
	// # update the haste with new / old flurry
	player.flurry = 3
	player.update_haste()
}

func (player *Player) RemoveFlurry() {
	player.flurry = player.flurry - 1
	if player.flurry < 0 {
		player.flurry = 0
		player.update_haste()
	}
}

func (player *Player) CanSwingMH(combat_time float64) bool {
	slot := 0

	wep_can_swing := combat_time >= player.NextSwing[slot]
	wep_staggered := combat_time-player.LastSwing >= 0.2 // blizzard makes you alternate swings with a gap

	return wep_can_swing && wep_staggered
}

func (player *Player) Swing(hand string, combat_time float64, force bool) (float64, string) {
	slot := 0
	weapon := player.Mainhand
	if hand == "offhand" {
		slot = 1
		weapon = player.Offhand
	}

	if player.CanSwingMH(combat_time) || force { // force used for windfury
		player.NextSwing[slot] = combat_time + weapon.SwingSpeed // reset swing timer
		player.LastSwing = combat_time

		// use up a flurry charge
		player.RemoveFlurry()

		damage := float64(0)
		hitType := "cast"

		// check if we're doing heroic strike
		// if player.QueueHeroicStrike {
		// 	player.QueueHeroicStrike = false
		// 	damage, hitType = Sim.cast("heroic_strike", combat_time)

		// 	// return damage, hitType
		// }

		// white hit damage
		damage = player.normalize_swing(weapon, 0) //weapon.BonusDamage)

		// calculate the swing against the hit table
		hitType = "cast"
		damage, hitType = player.Hittable.CalcWhiteHit(damage)

		// generate some rage
		player.generate_rage(damage, weapon, hitType == "crit" && true || false)

		// log this automatically
		// fmt.Println(combat_time, hand, math.Round(damage), hitType, player.flurry, player.rage)
		// player.Sim.log(weapon, damage, hitType)

		// try to proc windfury (yes even off of heroic strike)
		if combat_time > player.last_wf+3 {
			player.last_wf = combat_time
			roll := randint(0, 100)
			if roll > 20 {
				player.Stats.AttackPower += 445
				player.Swing("mainhand", combat_time, true) // force a new melee, since we procced WF
				player.Stats.AttackPower -= 445
			}
		}

		// return hit data
		return damage, hitType
	}

	return 0, ""
}
func (player *Player) update_gcd() {
	player.gcd = player.base_gcd - ((player.base_gcd / 2) * player.Stats.Haste)
	if player.gcd < 1 {
		player.gcd = 1
	}
}

func (player *Player) update_haste() {
	player.Mainhand.SwingSpeed = player.Mainhand.BaseSwingSpeed
	player.Offhand.SwingSpeed = player.Offhand.BaseSwingSpeed

	// flurry goes first
	if player.flurry > 0 {
		increase := float64(25)
		player.Mainhand.SwingSpeed = player.Mainhand.BaseSwingSpeed / ((100 + increase) / 100)
		player.Offhand.SwingSpeed = player.Offhand.BaseSwingSpeed / ((100 + increase) / 100)
	}

	// then add player haste
	increase := player.Stats.Haste * 100
	player.Mainhand.SwingSpeed = player.Mainhand.SwingSpeed / ((100 + increase) / 100)
	player.Offhand.SwingSpeed = player.Offhand.SwingSpeed / ((100 + increase) / 100)

	// fmt.Println("haste total", player.Mainhand.SwingSpeed)
	player.update_gcd()
}
func (player *Player) OffGCDCheck(combat_time float64) bool {
	if combat_time >= player.next_gcd {
		return true
	}
	return false
}

func (player *Player) Cast(ability *Ability, combat_time float64) (float64, string) {
	// fmt.Println(ability, combat_time)

	if player.OffGCDCheck(combat_time) && ability.CanCast(combat_time) && player.Rage >= ability.cost {
		// ok we're casting it, spend rage
		player.spend_rage(ability.cost)

		// set the last GCD to now
		if ability.triggersGCD {
			player.last_gcd = combat_time
			player.next_gcd = combat_time + player.gcd
		}

		// # raw damage
		damage := ability.Cast(player, player.Target, combat_time)

		// # now reduce/hittable it
		hitType := "cast"
		if damage != 0 {
			damage, hitType = player.Hittable.CalcSpecialHit(damage)
		}

		// print output
		// fmt.Println(combat_time, ability.Name, math.Round(damage), hitType)

		// # log this automatically
		// self.Sim.log(name, damage, hitType)

		return damage, hitType

	}

	return 0, ""
}

func NewPlayer() Player {
	player := Player{
		level: 70,

		Equipment:    make(map[string]items.Item, 17),
		Enchantments: make(map[string]items.Enchant, 13),

		NextSwing:     make([]float64, 2),
		Race:          "human",
		strength_mult: 1,
		agility_mult:  1,

		QueueHeroicStrike: false,
		last_anger_rage:   0,
		last_bloodrage:    0,
		blood_rage_total:  0,
		blood_rage_active: false,
		flurry:            0,
		last_wf:           -3,

		base_gcd:      1.5,
		gcd:           1.5,
		off_gcd:       true,
		next_gcd:      0,
		last_gcd:      10,
		reported_gcd1: false,
		reported_gcd2: false,
	}

	// rage factor for auto attacks
	player.rage_factor = rage_factor(player.level)

	// stats while naked at level 70
	player.character_stats.CritChance = 1.14
	player.character_stats.Strength = 145
	player.character_stats.AttackPower = 190
	player.character_stats.Agility = 96
	player.DamageMult = 1
	player.CritMult = 0.03

	// set default talents
	player.talents.weapon_mastery = 2
	player.talents.improved_whirlwind = 1
	player.talents.improved_heroic_strike = 3
	player.talents.impale = 2
	player.talents.cruelty = 5
	player.talents.commanding_presence = 5
	player.talents.dual_wield_specialization = 5
	player.talents.improved_berserker_stance = 5

	// # add_action("combat_tick_first", self.update)

	return player
}
