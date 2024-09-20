package main

import (
	"fmt"
	"glade/WoWSim/core"
	"glade/WoWSim/items"
	"math"
	"math/rand"
	"sync"
	"time"
)

func simulate(wg *sync.WaitGroup, sim chan SimInstance, i int) {
	defer wg.Done()
	wg.Add(1)

	// init main objects
	player := core.NewPlayer()
	target := core.NewTarget()
	hittable := core.NewHitTable(&player, &target)

	// change any player or target variables here

	// now the sim object
	Sim := SimInstance{
		name:      fmt.Sprintf("%s%d", "thread - ", i),
		player:    &player,
		target:    &target,
		abilities: core.NewAbilities(&player),
		hittable:  &hittable,
	}

	// attack pointers to objects to reduce paramters
	player.Target = &target
	player.Hittable = &hittable
	hittable.Player = &player
	hittable.Target = &target

	// let's equip some items
	player.Equip("Destroyer Battle-Helm", "Helm")
	player.Equip("Pendant of the Perilous", "Neck")
	player.Equip("Blood-stained Pauldrons", "Shoulders")
	player.Equip("Shadowmoon Destroyer's Drape", "Back")
	player.Equip("Heartshatter Breastplate", "Chest")
	// player.Equip("Destroyer Breastplate", "Chest")
	player.Equip("Bracers of Eradication", "Bracers")
	player.Equip("Pillager's Gauntlets", "Gloves")
	// player.Equip("Red Belt of Battle", "Belt")
	player.Equip("Vindicator's Plate Belt", "Belt")
	player.Equip("Destroyer Greaves", "Legs")
	// player.Equip("Gladiator's Plate Legguards", "Legs")
	player.Equip("Dreadboots of the Legion", "Boots")
	player.Equip("Band of Devastation", "Ring1")
	player.Equip("Ring of a Thousand Marks", "Ring2")
	player.Equip("Bloodlust Brooch", "Trinket1")
	player.Equip("Empty Mug of Direbrew", "Trinket2")
	player.Equip("Dragonstrike", "Mainhand")
	player.Equip("Talon of Azshara", "Offhand")
	player.Equip("Sunfury Bow of the Phoenix", "Ranged")

	player.Enchant("Glyph of Ferocity", "Helm")
	player.Enchant("Greater Inscription of Vengeance", "Shoulders")
	player.Enchant("Enchant Cloak - Greater Agility", "Back")
	player.Enchant("Chest - Exceptional Stats", "Chest")
	player.Enchant("Bracer - Brawn", "Bracers")
	player.Enchant("Gloves - Major Strength", "Gloves")
	player.Enchant("Nethercobra Leg Armor", "Legs")
	player.Enchant("Enchant Boots - Cat's Swiftness", "Boots")
	player.Enchant("Mongoose", "Mainhand")
	player.Enchant("Mongoose", "Offhand")

	player.Stats.MeleeCrit += 4 // hardcoding epic gems in boots right now

	// now lets calculate player stats
	player.CalculateStats()
	Sim.CreateLogs()

	// start loop
	ticks := Settings.Duration * (1 / Settings.PerSecond)
	for i := float64(0); i <= ticks; i++ {
		combat_time := i * Settings.PerSecond
		combat_time = float64(math.Round(float64(combat_time*100))) / 100
		target.Health = (100 - combat_time/Settings.Duration*100) // set target hp
		Sim.combat_time = combat_time                             // store combat time

		// ttexecute := time_until_execute(combat_time, &target)

		// check if we can swing weapons and get rage
		if player.CanSwingMH(combat_time) && player.QueueHeroicStrike {
			player.QueueHeroicStrike = false
			player.NextSwing[0] = combat_time + player.Mainhand.SwingSpeed // reset swing timer
			player.LastSwing = combat_time
			player.RemoveFlurry()
			Sim.cast("heroic_strike")
		} else {
			Sim.swing("mainhand")
		}
		Sim.swing("offhand")

		// bloodthirst always #1 priority
		Sim.cast("bloodthirst")

		// EXECUTE ROTATION HERE
		if target.Health <= Settings.ExecuteRange {
			if Sim.get_ability_cooldown("bloodthirst") < 2 && player.Rage < 20 {
				continue // we'll wait for bloodthirst and make sure we have rage
			}
			Sim.cast("execute")
			continue // we're in execute range, we've casted execute. nice, move on
		}

		Sim.cast("whirlwind") // whirlwind

		// should we heroic strike?
		if (Sim.get_ability_cooldown("bloodthirst") > 1 && player.Rage >= 60) || player.Rage > 60 {
			player.QueueHeroicStrike = true
		}
	}

	fmt.Println(Sim.CombatLog["heroic_strike"])
	fmt.Println(Sim.CombatLog["bloodthirst"])
	fmt.Println(Sim.CombatLog["whirlwind"])
	fmt.Println(Sim.CombatLog["execute"])

	sim <- Sim
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	// misc
	items.Init()

	// Dispatch the threads for sim
	start := time.Now()
	sims := make(chan SimInstance)
	for i := 0; i < Settings.Iterations; i++ {
		go simulate(&wg, sims, i)
	}

	// Read from the sims now and tally things up
	all_sims := make([]SimInstance, Settings.Iterations)
	for i := 0; i < Settings.Iterations; i++ {
		all_sims[i] = <-sims
	}

	// Close everything out once goroutines finish up
	close(sims)
	wg.Wait()

	// Display results
	total_dmg := float64(0)
	for _, sim := range all_sims {
		for _, log := range sim.CombatLog {
			total_dmg += log.damage
		}
		// fmt.Println(sim.CombatLog["bloodthirst"])
	}

	fmt.Println(total_dmg / float64(Settings.Iterations) / Settings.Duration)

	// All done!
	duration := time.Since(start)
	fmt.Println("")
	fmt.Printf("All routines finished. Sims took %s", duration)

	// fmt.Println("")
	// fmt.Println("... Press the Enter Key to terminate the console screen")
	// fmt.Scanln() // wait for Enter Key
}
