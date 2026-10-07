package menu

import (
	"duckling/src/audio"
	"duckling/src/i18n"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_main.go — estado principal: 3 botões, L troca idioma,
// ESC abre confirmação de saída.

func (m *Menu) click() { _ = audio.Play("click") }

func (m *Menu) mainUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		i18n.Toggle()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.state = stQuit
		return
	}
	if m.play.Update() {
		m.click()
		m.next("play") // no-op até existir a tela play
	}
	if m.shop.Update() {
		m.click()
		m.next("shop")
	}
	if m.opt.Update() {
		m.click()
		m.state = stSettings
	}
	if m.quit.Update() {
		m.click()
		m.state = stQuit
	}
}
