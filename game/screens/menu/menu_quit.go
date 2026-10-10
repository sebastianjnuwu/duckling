package menu

import (
	"os"

	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_quit.go — modal "tem certeza?": SIM sai (os.Exit),

func (m *Menu) quitUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.state = stMain
		return
	}
	if m.qNo.Update() {
		m.click()
		m.state = stMain
	}
	if m.qX.Push() {
		m.click()
		m.state = stMain
	}
	if m.qYes.Update() {
		os.Exit(0)
	}
}

func (m *Menu) quitDraw(img *ebiten.Image) {
	ui.Dim(img)
	ui.Panel(img, 360, 230, 560, 260)
	m.qX.Draw(img, "X", m.font)
	ui.DrawCenter(img, i18n.T("menu.confirm"), m.font, 330)
	m.qYes.Draw(img, i18n.T("menu.yes"), m.font)
	m.qNo.Draw(img, i18n.T("menu.no"), m.font)
}
