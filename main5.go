package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	WalletAddress  string
	PlayerName     string
	GameTokens     int
	ArtTokens      int
	TechTokens     int
	ArtXP          int
	GameXP         int
	TechXP         int
	RiddleAttempts map[string]bool // Track riddle attempts
	RiddleScore    int             // Track riddle score
}

type Game struct {
	Players     map[string]*Player // Keyed by wallet address
	AllowList   map[string]bool    // Keyed by wallet address
	Purgatory   map[string]*Player // Keyed by wallet address
	CurrentUser string             // tracking the logged in user
}

const tippiWalletAddress = "0xTippi"

// NewGame initializes a new game environment.
func NewGame() *Game {
	game := &Game{
		Players:   make(map[string]*Player),
		AllowList: make(map[string]bool),
		Purgatory: make(map[string]*Player),
	}

	// Add "0xTippi" to the AllowList
	game.AllowList["0xTippi"] = true

	game.Players["0xTippi"] = &Player{
		WalletAddress:  "0xTippi",
		PlayerName:     "Tippi",
		GameTokens:     10,
		ArtTokens:      5,
		TechTokens:     20,
		ArtXP:          100,
		GameXP:         500,
		TechXP:         1000,
		RiddleAttempts: make(map[string]bool),
		RiddleScore:    5,
	}

	return game
}

func (g *Game) Login(walletAddress string) {
	if g.IsAllowed(walletAddress) {
		g.CurrentUser = walletAddress
		fmt.Println("Login successful.")
		// do i need more login logic?

	} else {
		fmt.Println("Login failed.")
	}
}

// SaveGame writes the current game state to a file.
func (g *Game) SaveGame(filename string) error {
	data, err := json.Marshal(g)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// LoadGame reads a game state from a file.
func LoadGame(filename string) (*Game, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var game Game
	err = json.Unmarshal(data, &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (g *Game) IsAllowed(walletAddress string) bool {
	_, allowed := g.AllowList[walletAddress]
	return allowed || walletAddress == "wallet"
}

// When adding a new player, initialize the RiddleAttempts map
func (g *Game) AddPlayer(walletAddress, playerName string) {
	if _, exists := g.AllowList[walletAddress]; !exists {
		g.Players[walletAddress] = &Player{
			WalletAddress:  walletAddress,
			PlayerName:     playerName,
			GameTokens:     5,
			ArtTokens:      1,
			TechTokens:     10,
			RiddleAttempts: make(map[string]bool), // Initialize the map
		}
		g.AllowList[walletAddress] = true
		fmt.Printf("Player %s added with starting tokens.\n", playerName)
	} else {
		fmt.Println("Player already exists.")
	}
}

// RemovePlayer removes a player from the game, moving them to Purgatory.
func (g *Game) RemovePlayer(walletAddress string) {
	if player, exists := g.Players[walletAddress]; exists {
		g.Purgatory[walletAddress] = player // Move to Purgatory
		delete(g.Players, walletAddress)    // Remove from active players
		delete(g.AllowList, walletAddress)  // Remove from allow list
		fmt.Printf("Player %s has been moved to Purgatory.\n", player.PlayerName)
	} else {
		fmt.Println("Player not found.")
	}
}

// AwardTokensXP awards tokens and XP to a player.
func (g *Game) AwardTokensXP(walletAddress string, gameTokens, artTokens, techTokens, artXP, gameXP, techXP int) {
	if player, exists := g.Players[walletAddress]; exists {
		player.GameTokens += gameTokens
		player.ArtTokens += artTokens
		player.TechTokens += techTokens
		player.ArtXP += artXP
		player.GameXP += gameXP
		player.TechXP += techXP
		fmt.Println("Awards and XP have been updated for", player.PlayerName)
	} else {
		fmt.Println("Player not found.")
	}
}

// security function to check if the current user is tippi to give access
func (g *Game) IsTippi() bool {
	return g.CurrentUser == tippiWalletAddress
}

// startTutorial encapsulates the tutorial logic.
func startTutorial(buf *bufio.Reader) {
	fmt.Println("\n--- Welcome to the Tutorial! ---")
	fmt.Println(`1. Setting Availability and Preferences
Your presence in the Astrovan isn't just about being there; it's about making sure you're there at the right time. This is where you set your game availability.`)

	// Simulate interaction or more explanations

	fmt.Println(`2. Registering for Game Sessions
Excited for an adventure? Here's how you register for the next session of Astrovan. It's simpler than dodging ScanBots.`)

	// More tutorial content...

	fmt.Println("\nTutorial completed! Are you ready to start your adventure, or do you need help (type 'help' for more commands)?")
	// Based on further user input, you can repeat or exit the tutorial.
}

// Generates a scaled bar based on the maximum value in the dataset.
func generateScaledBar(value, maxValue int) string {
	const scaleSize = 10
	if maxValue == 0 {
		return strings.Repeat(" ", scaleSize)
	}
	barLength := int((float64(value) / float64(maxValue)) * scaleSize)
	return strings.Repeat("=", barLength) + strings.Repeat(" ", scaleSize-barLength)
}

// Helper function to find the maximum of three integers.
func max(a, b, c int) int {
	return int(math.Max(math.Max(float64(a), float64(b)), float64(c)))
}

// Generates ASCII chart for the logged-in player
func (player *Player) generateChart() {
	if player == nil {
		fmt.Println("No player data available to generate chart.")
		return
	}
	// Find the max values for scaling
	maxTokens := max(player.GameTokens, player.ArtTokens, player.TechTokens)
	maxXP := max(player.ArtXP, player.GameXP, player.TechXP)

	// Generate the bars
	gameTokensBar := generateScaledBar(player.GameTokens, maxTokens)
	artTokensBar := generateScaledBar(player.ArtTokens, maxTokens)
	techTokensBar := generateScaledBar(player.TechTokens, maxTokens)

	gameXPBar := generateScaledBar(player.GameXP, maxXP)
	artXPBar := generateScaledBar(player.ArtXP, maxXP)
	techXPBar := generateScaledBar(player.TechXP, maxXP)

	// Print the bars
	fmt.Printf("Tokens\n")
	fmt.Printf("Game Tokens: [%s] %d\n", gameTokensBar, player.GameTokens)
	fmt.Printf("Art Tokens:  [%s] %d\n", artTokensBar, player.ArtTokens)
	fmt.Printf("Tech Tokens: [%s] %d\n", techTokensBar, player.TechTokens)

	fmt.Printf("\nExperience Points\n")
	fmt.Printf("Game XP: [%s] %d\n", gameXPBar, player.GameXP)
	fmt.Printf("Art XP:  [%s] %d\n", artXPBar, player.ArtXP)
	fmt.Printf("Tech XP: [%s] %d\n", techXPBar, player.TechXP)
}

func main() {
	game := NewGame()
	buf := bufio.NewReader(os.Stdin)
	fmt.Println(`Welcome to Ceptor Club's "Drive, Astrovan, Drive"!

You see Grampa the Astrovan rolling up, with your old friend Tippi at the wheel. "Hop in, no time to explain!" he shouts, and then as you take your seat, almost immediately hits the accelerator.

Smoke fills the interior and you are transported into the adventure: 
Save the multiverse from the Quantum Digitizers and their insidious ScanBots!`)
	// Tutorial Decision
	fmt.Print("\nDo you want to skip the tutorial? (Y/n): ")
	input, _ := buf.ReadString('\n')
	input = strings.TrimSpace(input)

	// Tutorial Process
	if strings.ToLower(input) == "n" {
		// Start Tutorial
		startTutorial(buf)
	} else {
		// Skip Tutorial
		fmt.Println("Skipping tutorial. Fastening seat belts...")
		// Any additional setup before starting the game can be placed here.
	}

	for {
		fmt.Print("> ")
		input, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)

		// Parse input for commands
		args := strings.Split(input, " ")
		command := args[0]

		switch command {
		case "login":
			if len(args) < 2 {
				fmt.Println("Usage: login <walletAddress>")
				continue
			}
			walletAddress := args[1]
			game.Login(walletAddress)
			// prompt user to load a game state, listing the game states available (files in the directory not ending in .go)
			files, err := ioutil.ReadDir(".")
			if err != nil {
				fmt.Println("Error reading directory:", err)
				continue
			}

			fmt.Println("Available game states:")
			for _, file := range files {
				if !file.IsDir() && !strings.HasSuffix(file.Name(), ".go") {
					fmt.Println(file.Name())
				}
			}
		case "add":
			if !game.IsTippi() {
				fmt.Println("You are not allowed to add players.")
				continue
			}
			if len(args) < 3 {
				fmt.Println("Usage: add <walletAddress> <playerName>")
				continue
			}
			walletAddress := args[1]
			playerName := strings.Join(args[2:], " ") // In case the name consists of multiple words
			game.AddPlayer(walletAddress, playerName)
		case "chart":
			if game.CurrentUser == "" {
				fmt.Println("You must be logged in to view the chart.")
			} else {
				currentPlayer, ok := game.Players[game.CurrentUser]
				if !ok || currentPlayer == nil {
					fmt.Println("Current user not found in players.")
				} else {
					currentPlayer.generateChart()
				}
			}

		case "list", "ls":
			fmt.Println("Players:")
			for _, player := range game.Players {
				fmt.Printf("%s (%s)\n", player.PlayerName, player.WalletAddress)
			}
		case "allowlist":
			fmt.Println("Allowed Wallet Addresses:")
			for walletAddress := range game.AllowList {
				fmt.Println(walletAddress)
			}
		case "remove":
			if !game.IsTippi() {
				fmt.Println("You are not allowed to remove players.")
				continue
			}
			if len(args) < 2 {
				fmt.Println("Usage: remove <walletAddress>")
				continue
			}
			walletAddress := args[1]
			game.RemovePlayer(walletAddress)
		case "award":
			if !game.IsTippi() {
				fmt.Println("You are not allowed to award tokens and XP.")
				continue
			}
			if len(args) < 8 {
				fmt.Println("Usage: award <walletAddress> <gameTokens> <artTokens> <techTokens> <artXP> <gameXP> <techXP>")
				continue
			}
			walletAddress := args[1]
			gameTokens, _ := strconv.Atoi(args[2])
			artTokens, _ := strconv.Atoi(args[3])
			techTokens, _ := strconv.Atoi(args[4])
			artXP, _ := strconv.Atoi(args[5])
			gameXP, _ := strconv.Atoi(args[6])
			techXP, _ := strconv.Atoi(args[7])
			game.AwardTokensXP(walletAddress, gameTokens, artTokens, techTokens, artXP, gameXP, techXP)
		case "help":
			fmt.Println("Commands:")
			fmt.Println("login <walletAddress> - Login to the game")
			fmt.Println("add <walletAddress> <playerName> - Add a new player (** RESTRICTED to Tippi **)")
			fmt.Println("list - List all active players")
			fmt.Println("allowlist - List all allowed wallet addresses")
			fmt.Println("remove <walletAddress> - Remove a player from the game (** RESTRICTED to Tippi **)")
			fmt.Println("award <walletAddress> <gameTokens> <artTokens> <techTokens> <artXP> <gameXP> <techXP> - Award tokens and XP to a player (** RESTRICTED to Tippi **)")
			fmt.Println("help - Display this help message")
			fmt.Println("chart - Display a chart of the logged-in player's tokens and XP")
			fmt.Println("save <filename> - Save the game state to a file (** RESTRICTED to Tippi **)")
			fmt.Println("load <filename> - Load the game state from a file (** RESTRICTED to Tippi **)")
			fmt.Println("locations - List all available locations")
			fmt.Println("read <locationName or number> - Read the description of a location")
			fmt.Println("riddle <language> - Get a riddle in the specified language (options: go, react, solidity)")
			fmt.Println("exit - Exit the game")
		case "check":
			if len(args) < 2 {
				fmt.Println("Usage: check <walletAddress>")
				continue
			}
			walletAddress := args[1]
			if player, exists := game.Players[walletAddress]; exists {
				fmt.Printf("Player: %s\n", player.PlayerName)
				fmt.Printf("Game Tokens: %d\n", player.GameTokens)
				fmt.Printf("Art Tokens: %d\n", player.ArtTokens)
				fmt.Printf("Tech Tokens: %d\n", player.TechTokens)
				fmt.Printf("Art XP: %d\n", player.ArtXP)
				fmt.Printf("Game XP: %d\n", player.GameXP)
				fmt.Printf("Tech XP: %d\n", player.TechXP)
				fmt.Printf("Riddle Score: %d\n", player.RiddleScore)
			} else {
				fmt.Println("Player not found.")
			}
		case "save":
			if !game.IsTippi() {
				fmt.Println("You are not allowed to save the game state.")
				continue
			}
			if len(args) < 2 {
				fmt.Println("Usage: save <filename>")
				continue
			}
			filename := args[1]
			err := game.SaveGame(filename)
			if err != nil {
				fmt.Println("Error saving game:", err)
			} else {
				fmt.Println("Game saved to", filename, "successfully")
			}
		case "load":
			if !game.IsTippi() {
				fmt.Println("You are not allowed to load the game state.")
				continue
			}
			if len(args) < 2 {
				fmt.Println("Usage: load <filename>")
				continue
			}
			filename := args[1]
			loadedGame, err := LoadGame(filename)
			if err != nil {
				fmt.Println("Error loading game:", err)
			} else {
				game = loadedGame
				fmt.Println("Game loaded from", filename, "successfully")
			}
		case "locations":
			fmt.Println("Choose a location by number or name:")
			i := 1
			for name := range Locations {
				fmt.Printf("%d. %s\n", i, name)
				i++
			}
		case "read":
			if len(args) < 2 {
				fmt.Println("Usage: read <locationName or number>")
				continue
			}
			input := strings.Join(args[1:], " ")
			if num, err := strconv.Atoi(input); err == nil {
				// Input is a number, find the corresponding location by index
				i := 1
				for _, loc := range Locations {
					if i == num {
						fmt.Printf("%s: %s - %s\n", loc.Name, loc.Description, loc.Challenge)
						break
					}
					i++
				}
			} else {
				// Input is a name
				if loc, ok := Locations[input]; ok {
					fmt.Printf("%s: %s - %s\n", loc.Name, loc.Description, loc.Challenge)
				} else {
					fmt.Println("Location not found.")
				}
			}
		case "riddle":
			// Ensure the player is logged in
			if game.CurrentUser == "" {
				fmt.Println("You must be logged in to attempt riddles.")
				continue
			}
			currentPlayer := game.Players[game.CurrentUser]
			if len(args) < 2 {
				fmt.Println("Usage: riddle <language> (options: go, react, solidity)")
				continue
			}
			language := args[1] // This captures the second argument, e.g., "go" or "react" or "solidity"
			// Check if the player has already attempted this riddle
			if _, ok := currentPlayer.RiddleAttempts[language]; ok {
				fmt.Println("You've already attempted this riddle. Moving on...")
				continue
			}
			switch language {
			case "go":
				fmt.Println(`Here is a Go code snippet missing a crucial part:
				
		votes ___ []string{"Dog", "Cat", "Dog", "Dog"}
		      ^^^
			What should go here?
				`)
				fmt.Print("> ")
				answer, _, _ := buf.ReadLine()
				if string(answer) == ":=" {
					fmt.Println("Correct! ':=' is used to declare and initialize 'votes'.")
					currentPlayer.GameXP += 5
					currentPlayer.TechXP += 5
					currentPlayer.RiddleAttempts[language] = true
					currentPlayer.RiddleScore++
				} else {
					fmt.Println("'riddle go' answer incorrect! Go, try again. Maybe Google or ask OG Petey...")
				}

			case "react":
				fmt.Println(`Will this React code display the winning team based on the votes from a Solidity smart contract?
				
		// React Component Snippet [Display code here]
		
				Is this correct? (yes/no)`)
				fmt.Print("> ")
				answer, _, _ := buf.ReadLine()
				if strings.ToLower(string(answer)) == "yes" {
					fmt.Println("Correct! The code correctly displays the winning team.")
					currentPlayer.GameXP += 5
					currentPlayer.TechXP += 5
					currentPlayer.RiddleAttempts[language] = true
					currentPlayer.RiddleScore++
				} else {
					fmt.Println("Incorrect. The code is properly set up to display the winning team. Do not try again")
					currentPlayer.RiddleAttempts[language] = true
				}

			case "solidity":
				fmt.Println(`Identify the vulnerability in this Solidity function: [Describe vulnerability scenario here]
				
				Given TIPPI_ADDRESS is a constant and public, how might an attacker exploit this function to change the admin from a Cat team to a Dog team?`)
				fmt.Print("> ")
				answer, _, _ := buf.ReadLine()
				// Logic to evaluate the answer for solidity riddle
				if strings.Contains(strings.ToLower(string(answer)), "reentrancy") {
					fmt.Println("Correct! The function is vulnerable to reentrancy attacks.")
					currentPlayer.GameXP += 5
					currentPlayer.TechXP += 5
					currentPlayer.RiddleAttempts[language] = true
					currentPlayer.RiddleScore++
				} else {
					fmt.Println("Incorrect. Try again... 'riddle solidity'")
				}

			default:
				fmt.Println("Unknown riddle language. Options are: go, react, solidity")
			}
		case "exit":
			fmt.Println("Exciting game. May all your hooties, and this is important, dooty!")
			return
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}
