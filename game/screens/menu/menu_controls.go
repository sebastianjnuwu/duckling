package menu

import (
	"duckling/engine/input"
	"duckling/engine/save"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_controls.go — página 2: remapear teclas. Clica na
// ação, aperta a tecla nova. ESC cancela (não sai).

var keyRows = []struct {
	act input.Action
	key string
}{
	{input.Up, "settings.up"},
	{input.Down, "settings.down"},
	{input.Left, "settings.left"},
	{input.Right, "settings.right"},
	{input.Jump, "settings.jump"},
	{input.Fly, "settings.fly"},
}

func (m *Menu) controlsUpdate() {
	s := &m.set
	if s.cap >= 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			s.cap = -1
			return
		}
		if k, ok := input.Scan(); ok {
			input.Swap(input.Action(s.cap), k)
			s.cap = -1
			m.click()
			save.Store()
		}
		return
	}
	lo, hi := 0, 4
	if s.page != movePage() {
		lo, hi = 4, 6
	}
	for i := lo; i < hi; i++ {
		if s.keys[i].Update() {
			m.click()
			s.cap = i
			return
		}
	}
	if s.reset.Update() {
		m.click()
		input.Reset()
		save.Store()
	}
}

// keyName mostra a tecla ou "?" capturando.
func (m *Menu) keyName(i int) string {
	if m.set.cap == i {
		return "?"
	}
	return input.Name(input.Key(keyRows[i].act))
}
