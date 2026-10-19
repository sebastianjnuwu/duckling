package menu

import (
	"duckling/engine/input"
	"duckling/engine/rendering"
	"duckling/engine/save"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_controls.go — página 2: remapear teclas. Clica na
// ação, aperta a tecla nova. ESC cancela (não sai).

var keyRows = []struct {
	act input.Action
	key string
	ic  ui.Icon
}{
	{input.Up, "settings.up", ui.IconUp},
	{input.Down, "settings.down", ui.IconDown},
	{input.Left, "settings.left", ui.IconLeft},
	{input.Right, "settings.right", ui.IconRight},
	{input.Jump, "settings.jump", ui.IconJump},
	{input.Fly, "settings.fly", ui.IconFly},
	{input.Interact, "settings.interact", ui.IconHand},
	{input.Kill, "settings.kill", ui.IconKill},
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
		lo, hi = 4, len(keyRows)
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
		s.confirm = true
	}
}

// keyName mostra a tecla ou "?" capturando.
func (m *Menu) keyName(i int) string {
	if m.set.cap == i {
		return "?"
	}
	return input.Name(input.Key(keyRows[i].act))
}
