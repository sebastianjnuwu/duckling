package menu

import (
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// menu_mode_draw.go — modal online/lan no mesmo molde do quit.

func (m *Menu) modeDraw(img *ebiten.Image) {
	ui.Dim(img)
	ui.Panel(img, 360, 230, 560, 260)
	m.modeX.Draw(img, "X", m.font)
	m.online.Draw(img, i18n.T("menu.online"), m.font)
	m.lan.Draw(img, i18n.T("menu.lan"), m.font)
}
