package items

import (
	"glade/WoWSim/structs"
)

type Gem struct {
	ID      int
	Name    string
	Stats   structs.Stats // flat stats gem adds
	Color   int           //proto.GemColor
	Unique  bool
	Quality int
	Phase   byte
}

type Enchant struct {
	ID          int32 // ID of the enchant item.
	EffectID    int32 // Used by UI to apply effect to tooltip
	Name        string
	Quality     int
	Bonus       structs.Stats
	ItemType    int // Which slot the enchant goes on.
	EnchantType int // Additional category when ItemType isn't enough.
	Phase       int32
}

type Weapon struct {
	Hand           string
	Type           string
	MinDamage      int
	MaxDamage      int
	BonusDamage    int
	BaseSwingSpeed float64
	SwingSpeed     float64
	Item           Item
}

type Item struct {
	Name  string
	ID    int
	Stats structs.Stats

	Type             int
	ArmorType        int
	Hand             string
	HandType         int
	RangedWeaponType int
	WeaponType       int
	MinDamage        int
	MaxDamage        int
	BaseSwingSpeed   float64
	SwingSpeed       float64

	ClassAllowlist []int

	Phase   byte
	Quality int
	Unique  bool
	Ilvl    int

	GemSockets  []int
	SocketBonus structs.Stats

	// Modified for each instance of the item.
	Gems []Gem
}

var ByName = map[string]Item{}
var GemsByName = map[string]Gem{}
var EnchantsByName = map[string]Enchant{}

func Init() {
	for _, v := range Items {
		ByName[v.Name] = v
	}

	for _, v := range Gems {
		GemsByName[v.Name] = v
	}

	for _, v := range Enchants {
		EnchantsByName[v.Name] = v
	}
}
