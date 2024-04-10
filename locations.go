// Description: This file contains the definition of the Location struct and a map of Location objects representing the different locations in the game. Each Location object has a name, description, and challenge associated with it. The map is used to store and access the Location objects by their names.
package main

type Location struct {
	Name string
	Description string
	Challenge string
}

var Locations = map[string]Location{
	"Neon Forest": {
		Name: "Neon Forest",
		Description: "A dense jungle of glowing plants and animals. The air is thick with a sweet humidity, the sounds of chirping insects, and a buzz of electricity.",
		Challenge: "Navigating the maze-like paths that constantly reconfigure due to glitching influence while encountering cyber-enhanced wildlife.",
	},
	"Silicon Desert": {
		Name: "Silicon Desert",
		Description: "A vast expanse of fine silicon sand, dotted with ancient tech ruins and holographic mirages. The heat is intense, and the air shimmers with distortion.",
		Challenge: "Overcoming optical illusions and sandstorms that can erase digital memories.",
	},
	"Cryo-Mountain": {
		Name: "Cryo-Mountain",
		Description: "A towering peak surrounded by digital snowstorms, with a core of frozen data. The air is crisp, thin and cold; the ground is soft and sometimes slippery with ice.",
		Challenge: "Climbing the slippery slopes while battling against cold-based cyber creatures and avoiding data avalanches.",
	},
	"Mirror Lake": {
		Name: "Mirror Lake",
		Description: "A clear lake reflecting the constellations above, its waters hold the key to digital and astral convergence. The air is cool and fresh, and the water is clear and still.",
		Challenge: "Deciphering the reflections to reveal the path beneath the waters, while contending with reflective illusions.",
	},
	"Quantum Caves": {
		Name: "Quantum Caves",
		Description: "A network of caves where reality and virtuality merge, creating shifting dimensions and quantum puzzles. The air is heavy with the scent of earth and the sound of echoes.",
		Challenge: "Navigating the ever-changing caves and solving quantum riddles to unlock deep truths.",
	},
}


