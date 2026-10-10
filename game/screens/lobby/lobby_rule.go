package lobby

// lobby_rule.go — regras da sala: nome válido, teto de
// impostores e steppers que somem no limite.

// validName: nome da sala com 3 a 8 caracteres.
func validName(s string) bool {
	n := len([]rune(s))
	return n >= 3 && n <= 8
}

// maxImp: 1 impostor a cada 6 jogadores (6→1, 12→2, 14→3).
func (l *Lobby) maxImp() int { return (l.min + 5) / 6 }

// No limite o stepper some (igual dos dois lados).
func (l *Lobby) canMinDec() bool { return l.min > 6 }
func (l *Lobby) canMinInc() bool { return l.min < l.top }
func (l *Lobby) canMaxDec() bool { return l.max > 6 }
func (l *Lobby) canMaxInc() bool { return l.max < l.top }
func (l *Lobby) canImpDec() bool { return l.imp > 1 }
func (l *Lobby) canImpInc() bool { return l.imp < l.maxImp() }
