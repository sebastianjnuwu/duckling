package save

// coins.go — saldo local de moedas (a loja gasta na Fase 7).
// Mesmo arquivo de settings: entra em Current/Apply.

var coins int

// Coins é o saldo atual.
func Coins() int { return coins }

// AddCoins soma ao saldo (só positivo).
func AddCoins(n int) {
	if n > 0 {
		coins += n
	}
}
