package menu

import (
	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/engine/save"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_main.go — estado principal: 3 botões, L troca idioma,
// ESC abre confirmação de saída.

func (m *Menu) click() { _ = audio.Play("click") }

func (m *Menu) mainUpdate() {
	m.toast.Update()
	m.fx.Update()
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		i18n.Toggle()
		save.Store()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.state = stQuit
		return
	}
	if m.play.Update() {
		m.click()
		m.state = stMode
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
