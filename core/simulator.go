package core

type AbilityLog struct {
	name     string
	casts    int
	hitTypes map[string]int
	damage   float64
	min      float64
	max      float64
}

type Simulator struct {
	Name       string
	Log        AbilityLog
	CombatTime float64
}

func (sim *Simulator) cast(abilityname string) {

}
