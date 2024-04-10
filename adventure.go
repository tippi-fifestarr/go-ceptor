// Description: This file contains the Adventure and Scenario structs, as well as the MainAdventure variable that holds the main adventure's data.
package main

type Adventure struct {
	Name        string
	Description string
	Scenarios   map[string]Scenario
}

type Scenario struct {
	Location   string
	Challenge  string
}

var MainAdventure = Adventure{
	Name: "Hooty Dooty, everyone! Our worlds are being scanned by the mysterious Digitizers, which are looking to copy and then quantum compute all our D&D 5e realities for some nefarious reason (the Paywall)...",
	Description: "The Digitizers are scanning the multiverse for D&D 5e realities to copy and quantum compute. The players must find a way to disrupt the scanning process and protect their worlds.",
	Scenarios: map[string]Scenario{
		"Neon Forest1": {
			Location:  "Neon Forest",
			Challenge: "An ancient druidic hologram holds a fragment of code that can disrupt the Digitizers' scanning tech.",
		},
		"Neon Forest2": {
			Location:  "Neon Forest",
			Challenge: "A mythical creature, once a victim of digital replication, knows a secret path to the Digitizers' domain.",
		},
		"Neon Forest3": {
			Location:  "Neon Forest",
			Challenge: "The Guardian of the Forest is actually an ancient Digitizer who defected, possessing critical information.",
		},
		"Silicon Desert1": {
			Location:  "Silicon Desert",
			Challenge: "The spirits of the desert whisper of a buried device capable of shielding an area from digital scans.",
		},
		"Silicon Desert2": {
			Location:  "Silicon Desert",
			Challenge: "A hidden archive guarded by illusions contains the blueprint of the first Digitizer, revealing a critical vulnerability.",
		},
		"Silicon Desert3": {
			Location:  "Silicon Desert",
			Challenge: "An optical illusion created by the desert sands can camouflage essential data from the Digitizers' scans.",
		},
		"Mirror Lake1": {
			Location:  "Mirror Lake",
			Challenge: "The reflective waters can reveal the hidden location of a Digitizer's core processing unit during certain lunar phases.",
		},
		"Mirror Lake2": {
			Location:  "Mirror Lake",
			Challenge: "Submerged beneath the lake is an artifact that resonates with frequencies disruptive to the Digitizers.",
		},
		"Mirror Lake3": {
			Location:  "Mirror Lake",
			Challenge: "The lake is a natural scanner that can predict the Digitizers' next target, offering a chance to prepare defenses.",
		},
		"Cryo-Mountain1": {
			Location:  "Cryo-Mountain",
			Challenge: "Legends tell of an ice-entombed sage who predicted the arrival of the Digitizers and knew their weakness.",
		},
		"Cryo-Mountain2": {
			Location:  "Cryo-Mountain",
			Challenge: "A frozen obelisk contains an anti-digitization rune that, if deciphered, could protect entire realms.",
		},
		"Cryo-Mountain3": {
			Location:  "Cryo-Mountain",
			Challenge: "The summit's ancient observatory can pinpoint the source of the Digitizers' scanning beam.",
		},
		"Quantum Caves1": {
			Location:  "Quantum Caves",
			Challenge: "The caves are a labyrinth of shifting realities, concealing a portal to the Digitizers' homeworld.",
		},
	},