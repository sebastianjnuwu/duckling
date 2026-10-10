package lobby

import (
	"fmt"
	"math/rand/v2"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// lobby_cfg.go — passo 2: min/max, privada com senha, lista.

func (l *Lobby) configUpdate() error {
	if l.x.Push() || l.back.Update() {
		l.click()
		l.tab = 0
		return nil
	}
	if l.start.Update() && validName(l.name) {
		l.click()
		l.next("play")
		return nil
	}
	if l.canMinDec() && l.minDec.Update() {
		l.click()
		l.min = max(6, l.min-1)
		l.imp = min(l.imp, l.maxImp())
	}
	if l.canMinInc() && l.minInc.Update() {
		l.click()
		l.min = min(l.top, l.min+1)
		l.max = max(l.max, l.min)
		l.imp = min(l.imp, l.maxImp())
	}
	if l.canMaxDec() && l.maxDec.Update() {
		l.click()
		l.max = min(l.top, max(6, l.max-1))
		l.min = min(l.min, l.max)
	}
	if l.canMaxInc() && l.maxInc.Update() {
		l.click()
		l.max = min(l.top, l.max+1)
	}
	if l.canImpDec() && l.impDec.Update() {
		l.click()
		l.imp = max(1, l.imp-1)
	}
	if l.canImpInc() && l.impInc.Update() {
		l.click()
		l.imp = min(l.maxImp(), l.imp+1)
	}
	if l.privT.Update() {
		l.click()
		l.priv = l.privT.On
	}
	if l.priv && l.regen.Update() {
		l.click()
		l.pass = fmt.Sprintf("%05d", rand.IntN(100000))
	}
	if l.nameBox.Update() {
		l.click()
		l.nameFocus = true
	}
	if l.nameFocus {
		for _, r := range ebiten.AppendInputChars(nil) {
			if len([]rune(l.name)) < 8 && (unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ') {
				l.name += string(r)
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(l.name) > 0 {
			rs := []rune(l.name)
			l.name = string(rs[:len(rs)-1])
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			l.nameFocus = false
		}
	}
	return nil
}
