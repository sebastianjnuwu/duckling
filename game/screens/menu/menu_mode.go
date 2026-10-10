package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_mode.go — modal online/lan: online avisa em breve,
// lan entra no jogo (solo até o multiplayer chegar).

func (m *Menu) modeUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.state = stMain
		return
	}
	if m.modeX.Push() {
		m.click()
		m.state = stMain
	}
	if m.online.Update() {
		m.click()
		m.toast.Show()
		m.state = stMain
	}
	if m.lan.Update() {
		m.click()
		m.state = stMain
		m.next("lobby")
	}
}
