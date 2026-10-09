package menu

import (
	"image/color"

	"duckling/src/audio"
	"duckling/src/i18n"
	"duckling/src/save"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_main.go — estado principal: 3 botões, L troca idioma,
// ESC abre confirmação de saída.

func (m *Menu) click() { _ = audio.Play("Click") }

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
		m.next("play")
	}
	if m.shop.Update() {
		m.click()
		m.toast.Show()
		m.fx.Burst(940, 328, 24,
			color.RGBA{232, 190, 72, 255},
			color.RGBA{235, 240, 235, 255},
			color.RGBA{110, 231, 200, 255})
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
