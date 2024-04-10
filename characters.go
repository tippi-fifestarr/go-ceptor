package main

import (
	"fmt"
	"strings"
)

type Character struct {
	Name               string
	ClassAllocation    map[string]int
	Background         string
	Abilities          map[string]int
	EffectiveAbilities map[string]int
	Skills             []string
	Features           map[string][]string
	Equipment          []string
	Spells             map[string][]string
	Debuffs            map[string]int
}

func (c *Character) CalculateEffectiveAbilities() {
	// Initialize the EffectiveAbilities map if it's not already
	if c.EffectiveAbilities == nil {
		c.EffectiveAbilities = make(map[string]int)
	}

	// Start with copying base abilities to effective abilities
	for ability, score := range c.Abilities {
		c.EffectiveAbilities[ability] = score
	}

	// Apply Artificer debuffs and buffs
	artificerLevels := c.ClassAllocation["Artificer"]
	if artificerLevels > 0 {
		c.EffectiveAbilities["Strength"] -= artificerLevels
		dexReduction := int(float64(artificerLevels)/2 + float64(0.5)) // Round up
		c.EffectiveAbilities["Dexterity"] -= dexReduction
		c.EffectiveAbilities["Intelligence"] += artificerLevels // Adjusted for the Artificer's Toll buff
	}

	// Apply Barbarian debuffs and buffs for levels above 1
	barbarianLevels := c.ClassAllocation["Barbarian"]
	if barbarianLevels > 1 {
		extraLevels := barbarianLevels - 1
		c.EffectiveAbilities["Charisma"] -= 2 * extraLevels
		c.EffectiveAbilities["Wisdom"] += extraLevels
		c.EffectiveAbilities["Strength"] += extraLevels // Adjusted for Barbarian's Rage buff
	}
}

func (c *Character) Display() {
	fmt.Println("Name:", c.Name)
	fmt.Println("Class Allocation:")
	for class, level := range c.ClassAllocation {
		fmt.Printf("- %s: %d\n", class, level)
	}
	fmt.Println("Background:", c.Background)
	fmt.Println("Base Abilities:")
	for ability, score := range c.Abilities {
		fmt.Printf("- %s: %d\n", ability, score)
	}
	fmt.Println("Effective Abilities After Debuffs:")
	for ability, score := range c.EffectiveAbilities {
		fmt.Printf("- %s: %d\n", ability, score)
	}
	fmt.Println("Skills:")
	for _, skill := range c.Skills {
		fmt.Println("-", skill)
	}
	fmt.Println("Features:")
	for class, features := range c.Features {
		fmt.Printf("- %s: %v\n", class, features)
	}
	fmt.Println("Equipment:")
	for _, item := range c.Equipment {
		fmt.Println("-", item)
	}
	fmt.Println("Spells:")
	for level, spells := range c.Spells {
		fmt.Printf("- %s: %v\n", level, spells)
	}
	fmt.Println("Debuffs:")
	for debuff, value := range c.Debuffs {
		fmt.Printf("- %s: %d\n", debuff, value)
	}
}

func NewSavageGuardian() *Character {
	character := &Character{
		Name: "Tippi - The Savage Guardian",
		ClassAllocation: map[string]int{
			"Barbarian": 3,
			"Artificer": 1,
			"Druid":     1,
		},
		Background: "Tippi, once a defender of his village, endured trials of strength and spirit. Captured and transformed by alien artificers, he rebelled against his captors, awakening to the druidic magic of the cosmos under the tutelage of Fish Naturally.",
		Abilities: map[string]int{
			"Strength":     14,
			"Dexterity":    12,
			"Constitution": 16,
			"Intelligence": 13,
			"Wisdom":       15,
			"Charisma":     10,
		},
		EffectiveAbilities: map[string]int{},
		Skills:             []string{"Athletics", "Survival", "Nature", "Perception"},
		Features: map[string][]string{
			"Barbarian": {"Rage", "Unarmored Defense", "Reckless Attack"},
			"Artificer": {"Magical Tinkering", "Spellcasting"},
			"Druid":     {"Druidic", "Wild Shape", "Starry Form (Archer), Spellcasting"},
		},
		Equipment: []string{"Greataxe", "Explorer's Pack", "Artisan's Tools (Tinker's Tools)", "Druidic Focus"},
		Spells: map[string][]string{
			"Cantrips":  {"Mending", "Produce Flame", "Guidance", "Druidcraft"},
			"1st Level": {"Cure Wounds", "Faerie Fire"},
		},
		Debuffs: map[string]int{
			"ArtificersToll": -1,
			"BarbariansWild": -2,
		},
	}

	character.CalculateEffectiveAbilities()
	return character
}

func NewTechnomageUprising() *Character {
	character := &Character{
		Name: "Tippi - The Technomage Rebel",
		ClassAllocation: map[string]int{
			"Barbarian": 1,
			"Artificer": 3,
			"Druid":     1,
		},
		Background: "Adapting the artificer's tools against them, Tippi engineers a rebellion, fueled by rage and newfound magical prowess. Under Fish Naturally's guidance, he uncovers a druidic connection that empowers his crusade.",
		Abilities: map[string]int{
			"Strength":     14,
			"Dexterity":    12,
			"Constitution": 16,
			"Intelligence": 13,
			"Wisdom":       15,
			"Charisma":     10,
		},
		Skills: []string{"Athletics", "Investigation", "Nature", "Perception"},
		Features: map[string][]string{
			"Barbarian": {"Rage", "Unarmored Defense"},
			"Artificer": {"Magical Tinkering", "Infuse Item", "The Right Tool for the Job", "Spellcasting"},
			"Druid":     {"Druidic", "Wild Shape", "Spellcasting"},
		},
		Equipment: []string{"Greataxe", "Explorer's Pack", "Artisan's Tools (Tinker's Tools)", "Druidic Focus"},
		Spells: map[string][]string{
			"Cantrips":  {"Mending", "Produce Flame", "Guidance", "Druidcraft"},
			"1st Level": {"Cure Wounds", "Faerie Fire", "Detect Magic", "Shield"},
		},
		Debuffs: map[string]int{
			"ArtificersToll": -3,
			"BarbariansWild": 0,
		},
	}

	character.CalculateEffectiveAbilities()
	return character
}

func NewCosmicProtector() *Character {
	character := &Character{
		Name: "Tippi - The Cosmic Protector",
		ClassAllocation: map[string]int{
			"Barbarian": 1,
			"Artificer": 1,
			"Druid":     3,
		},
		Background: "With a heart heavy from the destruction he's witnessed, Tippi delves into the mysteries of the cosmos under the tutelage of Fish Naturally. He learns to channel the fury of the storm, the resilience of the earth, and the warmth of the sun to protect those who cannot protect themselves.",
		Abilities: map[string]int{
			"Strength":     14,
			"Dexterity":    12,
			"Constitution": 16,
			"Intelligence": 13,
			"Wisdom":       15,
			"Charisma":     10,
		},
		Skills: []string{"Athletics", "Survival", "Nature", "Perception"},
		Features: map[string][]string{
			"Barbarian": {"Rage", "Unarmored Defense"},
			"Artificer": {"Magical Tinkering", "Spellcasting"},
			"Druid":     {"Druidic", "Wild Shape", "Circle of Stars", "Starry Form (Archer)", "Starry Form (Dragon)", "Starry Form (Chalice)", "Spellcasting"},
		},
		Equipment: []string{"Greataxe", "Explorer's Pack", "Artisan's Tools (Tinker's Tools)", "Druidic Focus"},
		Spells: map[string][]string{
			"Cantrips":  {"Mending", "Produce Flame", "Guidance", "Druidcraft"},
			"1st Level": {"Cure Wounds", "Faerie Fire", "Entangle", "Goodberry"},
			"2nd Level": {"Moonbeam", "Flaming Sphere"},
		},
		Debuffs: map[string]int{
			"ArtificersToll": -1, // Slight debuff due short but awful Artificer experience
			"BarbariansWild": 0,
		},
	}

	character.CalculateEffectiveAbilities()
	return character
}

func NewElementalWarden() *Character {
	character := &Character{
		Name: "Tippi - The Elemental Warden",
		ClassAllocation: map[string]int{
			"Barbarian": 2,
			"Artificer": 2,
			"Druid":     1,
		},
		Background: "Embracing his role as a guardian, Tippi harmonizes the raw energy of his barbaric roots with the refined craft of artifice. His awakening to druidic magic reinforces his resolve to be the shield against those who dare threaten the natural equilibrium.",
		Abilities: map[string]int{
			"Strength":     14,
			"Dexterity":    12,
			"Constitution": 16,
			"Intelligence": 13,
			"Wisdom":       15,
			"Charisma":     10,
		},
		Skills: []string{"Athletics", "Investigation", "Nature", "Perception"},
		Features: map[string][]string{
			"Barbarian": {"Rage", "Unarmored Defense", "Reckless Attack", "Danger Sense"},
			"Artificer": {"Magical Tinkering", "Infuse Item", "Spellcasting"},
			"Druid":     {"Druidic", "Wild Shape (Archer)", "Wild Shape (Dragon)", "Wild Shape (Chalice)", "Spellcasting"},
		},
		Equipment: []string{"Greataxe", "Explorer's Pack", "Artisan's Tools (Tinker's Tools)", "Druidic Focus"},
		Spells: map[string][]string{
			"Cantrips":  {"Mending", "Produce Flame", "Guidance", "Druidcraft"},
			"1st Level": {"Cure Wounds", "Faerie Fire", "Entangle"},
		},
		Debuffs: map[string]int{
			"ArtificersToll": -2, // Balancing technological manipulation with physical prowess not easy
			"BarbariansWild": -1, // Maintaining mental clarity amidst barbaric rage?
		},
	}

	character.CalculateEffectiveAbilities()
	return character
}

func NewArcaneReclaimer() *Character {
	character := &Character{
		Name: "Tippi - The Arcane Reclaimer",
		ClassAllocation: map[string]int{
			"Barbarian": 1,
			"Artificer": 2,
			"Druid":     2,
		},
		Background: "In a world teetering on the brink of ecological collapse, Tippi dedicates himself to the restoration of corrupted lands. Combining the analytical mind of an artificer with the natural intuition of a druid, he devises innovative solutions to heal the land and fight against those who would see it despoiled.",
		Abilities: map[string]int{
			"Strength":     14,
			"Dexterity":    12,
			"Constitution": 16,
			"Intelligence": 13,
			"Wisdom":       15,
			"Charisma":     10,
		},
		Skills: []string{"Athletics", "Investigation", "Nature", "Perception"},
		Features: map[string][]string{
			"Barbarian": {"Rage", "Unarmored Defense"},
			"Artificer": {"Magical Tinkering", "Infuse Item", "The Right Tool for the Job", "Spellcasting"},
			"Druid":     {"Druidic", "Wild Shape", "Circle Spells", "Spellcasting"},
		},
		Equipment: []string{"Greataxe", "Explorer's Pack", "Artisan's Tools (Tinker's Tools)", "Druidic Focus"},
		Spells: map[string][]string{
			"Cantrips":  {"Mending", "Produce Flame", "Guidance", "Druidcraft"},
			"1st Level": {"Cure Wounds", "Faerie Fire", "Entangle", "Goodberry"},
			"2nd Level": {"Moonbeam", "Flaming Sphere", "Lesser Restoration"},
			"3rd Level": {"Purify Food and Drink", "Protection from Energy"},
		},
		Debuffs: map[string]int{
			"ArtificersToll": -2,
			"BarbariansWild": 0,
		},
	}

	character.CalculateEffectiveAbilities()
	return character
}

func NewNaturesVanguard() *Character {
	character := &Character{
		Name: "Tippi - The Nature's Vanguard",
		ClassAllocation: map[string]int{
			"Barbarian": 1,
			"Artificer": 2,
			"Druid":     2,
		},
		Background: "With the wilds under threat, Tippi channels his barbarian rage into a fierce determination to protect nature. Harnessing both the inventive potential of artifice and the empowering magic of druidry, he stands as a beacon of resistance against the unnatural.",
		Abilities: map[string]int{
			"Strength":     14,
			"Dexterity":    12,
			"Constitution": 16,
			"Intelligence": 13,
			"Wisdom":       15,
			"Charisma":     10,
		},
		Skills: []string{"Athletics", "Investigation", "Nature", "Perception"},
		Features: map[string][]string{
			"Barbarian": {"Rage", "Unarmored Defense"},
			"Artificer": {"Magical Tinkering", "Infuse Item", "Spellcasting"},
			"Druid":     {"Druidic", "Wild Shape", "Circle Spells", "Spellcasting"},
		},
		Equipment: []string{"Greataxe", "Explorer's Pack", "Artisan's Tools (Tinker's Tools)", "Druidic Focus"},
		Spells: map[string][]string{
			"Cantrips":  {"Mending", "Produce Flame", "Guidance", "Druidcraft"},
			"1st Level": {"Cure Wounds", "Faerie Fire", "Entangle", "Goodberry"},
			"2nd Level": {"Moonbeam", "Flaming Sphere", "Barkskin"},
		},
		Debuffs: map[string]int{
			"ArtificersToll": -2,
			"BarbariansWild": -1,
		},
	}

	character.CalculateEffectiveAbilities()
	return character
}

func main() {
	fmt.Println("Welcome to Tippi's choices. Please enter a number or the first word of the name to choose:")
	fmt.Println("1. Savage Guardian")
	fmt.Println("2. Technomage Rebel")
	fmt.Println("3. Cosmic Protector")
	fmt.Println("4. Elemental Warden")
	fmt.Println("5. Arcane Reclaimer")
	fmt.Println("6. Natures Vanguard")

	var choice string
	fmt.Scanln(&choice)

	switch strings.ToLower(choice) {
	case "1", "savage":
		savageGuardian := NewSavageGuardian()
		savageGuardian.Display()
	case "2", "technomage":
		technomageUprising := NewTechnomageUprising()
		technomageUprising.Display()
	case "3", "cosmic":
		cosmicProtector := NewCosmicProtector()
		cosmicProtector.Display()
	case "4", "elemental":
		elementalWarden := NewElementalWarden()
		elementalWarden.Display()
	case "5", "arcane":
		arcaneReclaimer := NewArcaneReclaimer()
		arcaneReclaimer.Display()
	case "6", "natures":
		naturesVanguard := NewNaturesVanguard()
		naturesVanguard.Display()
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
