package tritslab

import (
	"math"
)

type TritsPlayerResponse struct {
	PlayerAddr *TritsAddress
	Game       *TritsGame
	Amount     uint64
}

func NewTritsPlayerResponse(game *TritsGame, amount uint64, my_addr *TritsAddress) *TritsPlayerResponse {
	r := new(TritsPlayerResponse)
	r.PlayerAddr = my_addr
	r.Game = game
	r.Amount = amount
	return r
}

type TritsPlayer struct {
	Addr   *TritsAddress // Player's address
	banker *TritsBanker  // Connection to the banker
}

func NewTritsPlayer(addr *TritsAddress, banker *TritsBanker) *TritsPlayer {
	player := new(TritsPlayer)
	player.banker = banker
	player.Addr = addr
	return player
}

// Composition of the base TritsPlayer
func NewTritsPlayerFactory(addr *TritsAddress, banker *TritsBanker, strategy string) *TritsPlayer {
	switch strategy {
	case "dumb":
		return NewTritsPlayer(addr, banker)
	}
	return nil
}

// Choose the table
func (p *TritsPlayer) ChooseTable() int8 {
	b := RandByte()
	return int8(b % 23)
}

// Choose amount
func (p *TritsPlayer) ChooseAmount() uint64 {
	pocket := p.Balance()
	if NOMINAL > 0 && NOMINAL < pocket {
		return NOMINAL
	} else {
		b := RandByte()
		pocket := p.Balance()
		percent := int8(b%5 + 1) // Bet no more between 1 - 5 percent of the pocket
		percent = 1              // Try with 1 percent
		return uint64(math.Round(float64(percent) / 100 * float64(pocket)))
	}
}

func (p *TritsPlayer) Balance() uint64 {
	return p.banker.Tell(p.Addr)
}

// Borrow money if necessary
func (p *TritsPlayer) Borrow(max_borrow uint64) uint64 {
	pocket := p.Balance()
	if max_borrow > 100*pocket {
		return max_borrow
	} else {
		return 0
	}
}

// Bet or not - some strategy should be implemented here
func (p *TritsPlayer) Bet(desk []*TritsGame) []*TritsPlayerResponse {
	var responses []*TritsPlayerResponse
	pocket := p.Balance()
	if pocket == 0 { // No money no honey
		return nil
	}

	// Try to find an uninitiated game
	var i int = 0
	for i < GAMES_ON_TABLE && desk[i].Nominal != 0 {
		i++
	}
	if i < GAMES_ON_TABLE { // Found an uninitiated game
		amount := p.ChooseAmount()
		if amount == 0 { // This guy is out of money
			if TD {
				l(LOG_NOTICE, LogName(p.Addr), " is broke !")
			}
			return nil
		}
		res := NewTritsPlayerResponse(desk[i], amount, p.Addr)
		responses = append(responses, res)
		if TD {
			l(LOG_NOTICE, LogName(p.Addr), " decides to init new game ", LogName(desk[i].ThisGame), ", with ", amount)
		}
		return responses
	}

	// No uninitiated game found, try to find the inbalance of 2 which I can afford
	var j int = 0
	for j < GAMES_ON_TABLE { // Loop through all games
		if (desk[j].GetInbalance() == 2) && (desk[j].Nominal < pocket) {
			break // If inbalance is 2 and the player can afford it
		}
		j++
	}
	if j < GAMES_ON_TABLE { // Found a good inbalance, bet
		res := NewTritsPlayerResponse(desk[j], desk[j].Nominal, p.Addr)
		responses = append(responses, res)
		if TD {
			l(LOG_NOTICE, LogName(p.Addr), " has ", pocket, " and finds inbalance 2 ", LogName(desk[j].ThisGame), ", nominal ", desk[j].Nominal)
		}
		return responses
	}

	// No good inbalance found, bet on the first one I can afford
	var k int = 0
	for k < GAMES_ON_TABLE {
		if desk[k].Nominal < pocket {
			break
		}
		k++
	}
	if k < GAMES_ON_TABLE { // Found something to bet on
		res := NewTritsPlayerResponse(desk[k], desk[k].Nominal, p.Addr)
		responses = append(responses, res)
		if TD {
			l(LOG_NOTICE, LogName(p.Addr), " has ", pocket, " and decides to play on ", LogName(desk[k].ThisGame), ", with ", desk[k].Nominal)
		}
		return responses
	}

	if TD {
		l(LOG_NOTICE, LogName(p.Addr), " has got ", pocket, " only and can't afford any game")
	}
	// I seem to be out of game
	return nil
}

// Human name for logging
func (p *TritsPlayer) Name() string {
	var name string = ""
	switch p.Addr.Raw() {
	case LenderAddr:
		name = "Lender"
	case BankAddr:
		name = "Bank"
	case NeoAddr:
		name = "Neo"
	case TrinityAddr:
		name = "Trinity"
	case AgentAddr:
		name = "Agent"
	case KeymakerAddr:
		name = "Keymaker"
	case MorpheusAddr:
		name = "Morheus"
	case NiobeAddr:
		name = "Niobe"
	case OracleAddr:
		name = "Oracle"
	case PersephoneAddr:
		name = "Persephone"
	case TwinsAddr:
		name = "Twins"
	case BugsAddr:
		name = "Bugs"
	case AnalystAddr:
		name = "Analyst"
	case SeraphAddr:
		name = "Seraph"
	case ArchitectAddr:
		name = "Architect"
	case BaneAddr:
		name = "Bane"
	default:
		name = p.Addr.Human()
	}
	return name
}

// Helper function to get player name by address statically
func PlayerName(addr *TritsAddress) string {
	b := NewTritsBanker(0)
	p := NewTritsPlayer(addr, b)
	return p.Name()
}
