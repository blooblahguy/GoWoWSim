package items

import (
	"glade/WoWSim/structs"
)

var Enchants = []Enchant{
	// Head
	{ID: 29192, EffectID: 3003, Name: "Glyph of Ferocity", Quality: ItemQualityUncommon, Bonus: structs.Stats{AttackPower: 34, MeleeHit: 16}, ItemType: ItemTypeHead},

	// Shoulder
	{ID: 28888, EffectID: 2986, Name: "Greater Inscription of Vengeance", Quality: ItemQualityRare, Bonus: structs.Stats{AttackPower: 30, MeleeCrit: 10}, ItemType: ItemTypeShoulder},
	{ID: 28910, EffectID: 2997, Name: "Greater Inscription of the Blade", Quality: ItemQualityRare, Bonus: structs.Stats{AttackPower: 20, MeleeCrit: 15}, ItemType: ItemTypeShoulder},
	{ID: 23548, EffectID: 2717, Name: "Might of the Scourge", Quality: ItemQualityEpic, Bonus: structs.Stats{AttackPower: 26, MeleeCrit: 14}, ItemType: ItemTypeShoulder},

	// Back
	{ID: 33150, EffectID: 2621, Name: "Enchant Cloak - Subtlety", Quality: ItemQualityRare, Bonus: structs.Stats{}, ItemType: ItemTypeBack},
	{ID: 11206, EffectID: 849, Name: "Enchant Cloak - Lesser Agility", Quality: ItemQualityUncommon, Bonus: structs.Stats{Agility: 3}, ItemType: ItemTypeBack},
	{ID: 34004, EffectID: 368, Name: "Enchant Cloak - Greater Agility", Quality: ItemQualityCommon, Bonus: structs.Stats{Agility: 12}, ItemType: ItemTypeBack},
	{ID: 28277, EffectID: 1441, Name: "Enchant Cloak - Greater Shadow Resistance", Quality: ItemQualityRare, Bonus: structs.Stats{}, ItemType: ItemTypeBack},

	// Chest
	{ID: 24003, EffectID: 2661, Name: "Chest - Exceptional Stats", Quality: ItemQualityCommon, Bonus: structs.Stats{Stamina: 6, Intellect: 6, Strength: 6, Agility: 6}, ItemType: ItemTypeChest},

	// Wrist
	{ID: 27899, EffectID: 2647, Name: "Bracer - Brawn", Quality: ItemQualityUncommon, Bonus: structs.Stats{Strength: 12}, ItemType: ItemTypeWrist},
	{ID: 34002, EffectID: 1593, Name: "Bracer - Assault", Quality: ItemQualityUncommon, Bonus: structs.Stats{AttackPower: 24}, ItemType: ItemTypeWrist},

	// Hands
	{ID: 33995, EffectID: 684, Name: "Gloves - Major Strength", Quality: ItemQualityUncommon, Bonus: structs.Stats{Strength: 15}, ItemType: ItemTypeHands},
	{ID: 33152, EffectID: 2564, Name: "Gloves - Major Agility", Quality: ItemQualityRare, Bonus: structs.Stats{Agility: 15}, ItemType: ItemTypeHands},
	{ID: 33153, EffectID: 2613, Name: "Gloves - Threat", Quality: ItemQualityRare, Bonus: structs.Stats{}, ItemType: ItemTypeHands},

	// Legs
	{ID: 29533, EffectID: 3010, Name: "Cobrahide Leg Armor", Quality: ItemQualityRare, Bonus: structs.Stats{AttackPower: 40, MeleeCrit: 10}, ItemType: ItemTypeLegs},
	{ID: 29535, EffectID: 3012, Name: "Nethercobra Leg Armor", Quality: ItemQualityEpic, Bonus: structs.Stats{AttackPower: 50, MeleeCrit: 12}, ItemType: ItemTypeLegs},

	// Feet
	{ID: 35297, EffectID: 2940, Name: "Enchant Boots - Boar's Speed", Quality: ItemQualityRare, Bonus: structs.Stats{Stamina: 9}, ItemType: ItemTypeFeet},
	{ID: 22544, EffectID: 2657, Name: "Enchant Boots - Dexterity", Quality: ItemQualityUncommon, Bonus: structs.Stats{Agility: 12}, ItemType: ItemTypeFeet},
	{ID: 28279, EffectID: 2939, Name: "Enchant Boots - Cat's Swiftness", Quality: ItemQualityRare, Bonus: structs.Stats{Agility: 6}, ItemType: ItemTypeFeet},
	{ID: 22545, EffectID: 2658, Name: "Enchant Boots - Surefooted", Quality: ItemQualityUncommon, Bonus: structs.Stats{MeleeHit: 10}, ItemType: ItemTypeFeet},

	// Weapon
	{ID: 16250, EffectID: 1897, Name: "Weapon - Superior Striking", Quality: ItemQualityUncommon, Bonus: structs.Stats{}, ItemType: ItemTypeWeapon},
	{ID: 16252, EffectID: 1900, Name: "Weapon - Crusader", Quality: ItemQualityUncommon, Bonus: structs.Stats{}, ItemType: ItemTypeWeapon},
	{ID: 22559, EffectID: 2673, Name: "Weapon - Mongoose", Quality: ItemQualityRare, Bonus: structs.Stats{}, ItemType: ItemTypeWeapon},
	{ID: 19445, EffectID: 2564, Name: "Weapon - Agility", Quality: ItemQualityCommon, Bonus: structs.Stats{Agility: 15}, ItemType: ItemTypeWeapon},
	{ID: 33165, EffectID: 3222, Name: "Weapon - Greater Agility", Quality: ItemQualityCommon, Bonus: structs.Stats{Agility: 20}, ItemType: ItemTypeWeapon},
	{ID: 22556, EffectID: 2670, Name: "2H Weapon - Major Agility", Quality: ItemQualityUncommon, Bonus: structs.Stats{Agility: 35}, ItemType: ItemTypeWeapon, EnchantType: EnchantTypeTwoHand},
	{ID: 33307, EffectID: 3225, Name: "Weapon - Executioner", Phase: 4, Quality: ItemQualityRare, Bonus: structs.Stats{}, ItemType: ItemTypeWeapon},

	// Ring
	{ID: 22535, EffectID: 2929, Name: "Ring - Striking", Quality: ItemQualityCommon, Bonus: structs.Stats{}, ItemType: ItemTypeFinger},
}
