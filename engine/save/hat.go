package save

import "sort"

// hat.go — chapéus: o que já foi comprado e o que está na
// cabeça. Entra no mesmo settings.json (Current/Apply).

var (
	hats = map[int]bool{0: true} // 0 = sem chapéu (sempre)
	hat  int
)

// Hat é o chapéu equipado (0 = sem chapéu).
func Hat() int { return hat }

// HatOwned diz se o chapéu já foi comprado.
func HatOwned(id int) bool { return id == 0 || hats[id] }

// BuyHat desconta o preço, marca como comprado e equipa.
func BuyHat(id, price int) bool {
	if id <= 0 || HatOwned(id) || !SpendCoins(price) {
		return false
	}
	hats[id] = true
	hat = id
	return true
}

// EquipHat troca o chapéu (precisa ter comprado).
func EquipHat(id int) bool {
	if !HatOwned(id) {
		return false
	}
	hat = id
	return true
}

// hatList devolve os comprados em ordem (pro JSON).
func hatList() []int {
	out := make([]int, 0, len(hats))
	for id := range hats {
		out = append(out, id)
	}
	sort.Ints(out)
	return out
}

// loadHats aplica a lista salva (0 sempre existe).
func loadHats(ids []int) {
	hats = map[int]bool{0: true}
	for _, id := range ids {
		if id > 0 {
			hats[id] = true
		}
	}
}
