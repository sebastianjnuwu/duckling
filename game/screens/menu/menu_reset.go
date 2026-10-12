package menu

import (
	"duckling/engine/i18n"
	"duckling/engine/input"
	"duckling/engine/rendering"
	"duckling/engine/save"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_reset.go — REDEFINIR confirma antes de voltar
// as teclas ao padrão (SIM reseta, NÃO/X/ESC fecha).

func (m *Menu) resetConfirmUpdate() {
	s := &m.set
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.confirm = false
		return
	}
	if s.cNo.Update() {
		m.click()
		s.confirm = false
	}
	if s.cX.Push() {
		m.click()
		s.confirm = false
	}
	if s.cYes.Update() {
		m.click()
		input.Reset()
		save.Store()
		s.confirm = false
	}
}

func (m *Menu) resetConfirmDraw(img *ebiten.Image) {
	s := &m.set
	ui.Dim(img)
	ui.Panel(img, 360, 230, 560, 260)
	s.cX.Draw(img, "X", m.font)
	ui.Title(img, ui.IconRefresh, i18n.T("settings.resetConfirm"), m.font, 330, ui.Cream)
	s.cYes.Draw(img, i18n.T("menu.yes"), m.font)
	s.cNo.Draw(img, i18n.T("menu.no"), m.font)
}
