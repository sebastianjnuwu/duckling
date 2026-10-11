package menu

import (
	"duckling/engine/i18n"

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
		m.toast.Show(i18n.T("play.soon"))
		m.state = stMain
	}
	if m.lan.Update() {
		m.click()
		m.state = stRooms
	}
}
