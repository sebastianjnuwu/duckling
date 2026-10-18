package menu

import (
	"strconv"
	"unicode"

	"duckling/engine/i18n"
	"duckling/engine/network"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_room_join.go — tab 2: entrar numa sala achada na LAN
// (nome do jogador + senha, quando a sala é privada).

const (
	jNameY = 215.0
	jPassY = 285.0
)

func (l *RoomState) joinUpdate() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || l.backJ.Update() {
		l.click()
		l.jFocus, l.tab = 0, 0
		return nil
	}
	if l.joinB.Update() {
		l.click()
		l.doJoin()
		return nil
	}
	if clicked(rcX, jNameY+4, 240, 48) {
		l.jFocus = 1
	}
	if l.sel.priv && clicked(rcX, jPassY+4, 240, 48) {
		l.jFocus = 2
	}
	typeInto(l.jFocus == 1, &l.pname, 8)
	typeInto(l.jFocus == 2, &l.jPass, 5)
	return nil
}

func (l *RoomState) doJoin() {
	name := l.pname
	if !validName(name) {
		name = "DUCK"
	}
	addr := l.sel.ip + ":" + strconv.Itoa(network.GamePort)
	s, err := session.Join(addr, name, l.jPass)
	if err != nil {
		l.toast.Show(i18n.T(err.Error()))
		return
	}
	l.scanStop()
	l.jFocus = 0
	session.Set(s)
	l.next("play")
}

// clicked diz se o mouse acabou de clicar dentro da caixa.
func clicked(x, y, w, h float64) bool {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false
	}
	mx, my := ebiten.CursorPosition()
	return float64(mx) >= x && float64(mx) <= x+w && float64(my) >= y && float64(my) <= y+h
}

// typeInto digita no campo em foco (letras/números, backspace).
func typeInto(focus bool, s *string, max int) {
	if !focus {
		return
	}
	for _, r := range ebiten.AppendInputChars(nil) {
		if len([]rune(*s)) < max && (unicode.IsLetter(r) || unicode.IsDigit(r)) {
			*s += string(r)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(*s) > 0 {
		rs := []rune(*s)
		*s = string(rs[:len(rs)-1])
	}
}
