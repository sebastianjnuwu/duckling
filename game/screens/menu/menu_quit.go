package menu

import (
	"image/color"
	"os"

	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_quit.go — modal "tem certeza?": SIM sai (os.Exit),

func (m *Menu) quitUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.state = stMain
		return
	}
	if m.qX.Update() || m.qNo.Update() {
		m.click()
		m.state = stMain
	}
	if m.qYes.Update() {
		os.Exit(0)
	}
}

func (m *Menu) quitDraw(img *ebiten.Image) {
	dim(img)
	panel(img, 360, 230, 560, 260)
	m.qX.Draw(img, "X", m.font)
	ui.DrawCenter(img, i18n.T("menu.confirm"), m.font, 330)
	m.qYes.Draw(img, i18n.T("menu.yes"), m.font)
	m.qNo.Draw(img, i18n.T("menu.no"), m.font)
}

// dim escurece a tela atrás do modal.
func dim(img *ebiten.Image) {
	ebitenutil.DrawRect(img, 0, 0, viewW, viewH, color.RGBA{0, 0, 0, 160})
}

func panel(img *ebiten.Image, x, y, w, h float64) {
	ebitenutil.DrawRect(img, x, y, w, h, color.RGBA{25, 25, 35, 255})
	ebitenutil.DrawRect(img, x+4, y+4, w-8, h-8, color.RGBA{55, 55, 75, 255})
}
