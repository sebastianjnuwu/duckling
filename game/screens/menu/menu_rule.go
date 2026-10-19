package menu

import (
	"math/rand/v2"
)

// menu_rule.go — regras da sala: nome, senha, teto de
// impostores e steppers que somem no limite.

// validName: nome da sala com 3 a 8 caracteres.
func validName(s string) bool {
	n := len([]rune(s))
	return n >= 3 && n <= 8
}

// passChars: sem 0/O/1/I/L pra ninguém confundir.
const passChars = "ABCDEFGHJKMNPQRSTUVWXYZ23456789"

// genPass: senha de 5 letras e números.
func genPass() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = passChars[rand.IntN(len(passChars))]
	}
	return string(b)
}

// maxImp: teto do override manual (o auto segue a regra do host).
// imp = 0 significa AUTO: 1 impostor até 6 jogadores, 2 a partir
// de 7 (ver game/match).
func (l *RoomState) maxImp() int { return max(1, l.min/3) }

// No limite o stepper some (igual dos dois lados).
func (l *RoomState) canMinDec() bool { return l.min > 2 }
func (l *RoomState) canMinInc() bool { return l.min < l.top }
func (l *RoomState) canMaxDec() bool { return l.max > l.min }
func (l *RoomState) canMaxInc() bool { return l.max < l.top }
func (l *RoomState) canImpDec() bool { return l.imp > 0 }
func (l *RoomState) canImpInc() bool { return l.imp < l.maxImp() }
