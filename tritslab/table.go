package tritslab

const GamesOnTable = 23

type TritsTable struct {
	table []*TritsGame // Croupier's table
}

func NewTritsTable() *TritsTable {
	dice := NewTritsDice()
	t := new(TritsTable)
	var i int8 = 0
	for i < GamesOnTable {
		game := NewTritsGame(t.GetCityAddress(i), TRITS_GAME_LONGEVITY, dice)
		t.table = append(t.table, game)
		i++
	}
	return t
}

// Define city names
func (t *TritsTable) GetCityName(index int8) string {
	var cities = [GamesOnTable]string{
		"Tokyo",
		"Delhi",
		"Shanghai",
		"SaoPaulo",
		"Mexico",
		"Dhaka",
		"Cairo",
		"Beijing",
		"Mumbai",
		"Osaka",
		"Karachi",
		"Chongqing",
		"Istanbul",
		"BuenosAires",
		"Kolkata",
		"Kinshasa",
		"Lagos",
		"Manila",
		"Tianjin",
		"Guangzhou",
		"RioDeJaneiro",
		"Lahore",
		"Bangalore"}
	if index >= 0 && index < GamesOnTable {
		return cities[index]
	} else {
		return ""
	}
}

// Define an address for each city
func (t *TritsTable) GetCityAddress(index int8) *TritsAddress {
	var addresses = [GamesOnTable]*TritsAddress{
		NewTritsAddress("1000000010000000100000001000000010000000"),
		NewTritsAddress("1000000110000000100000001000000010000001"),
		NewTritsAddress("1000000210000000100000001000000010000002"),
		NewTritsAddress("1000000310000000100000001000000010000003"),
		NewTritsAddress("1000000410000000100000001000000010000004"),
		NewTritsAddress("1000000510000000100000001000000010000005"),
		NewTritsAddress("1000000610000000100000001000000010000006"),
		NewTritsAddress("1000000710000000100000001000000010000007"),
		NewTritsAddress("1000000800000000100000001000000010000008"),
		NewTritsAddress("1000000910000000100000001000000010000009"),
		NewTritsAddress("1000001000000000100000001000000010000010"),
		NewTritsAddress("1000001100000000100000001000000010000011"),
		NewTritsAddress("1000001200000000100000001000000010000012"),
		NewTritsAddress("1000001300000000100000001000000010000013"),
		NewTritsAddress("1000001400000000100000001000000010000014"),
		NewTritsAddress("1000001500000000100000001000000010000015"),
		NewTritsAddress("1000001600000000100000001000000010000016"),
		NewTritsAddress("1000001700000000100000001000000010000017"),
		NewTritsAddress("1000001800000000100000001000000010000018"),
		NewTritsAddress("1000001900000000100000001000000010000019"),
		NewTritsAddress("1000020600000000100000001000000010000020"),
		NewTritsAddress("1000021600000000100000001000000010000021"),
		NewTritsAddress("1000022600000000100000001000000010000022"),
	}
	if index >= 0 && index < GamesOnTable {
		return addresses[index]
	} else {
		return nil
	}
}
