package menu

import (
	"image/color"
	"math"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// menu_draw.go — principal: bg, título no topo,
// patão na esquerda com sombra, 3 botões na direita.

func (m *Menu) Draw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.DrawCenterColored(img, i18n.T("app.title"), m.big, 47, shadow())
	ui.DrawCenterColored(img, i18n.T("app.title"), m.big, 44, color.RGBA{120, 120, 120, 255})
	ui.DrawCenterColored(img, i18n.T("app.title"), m.big, 40, gold())
	ui.DrawCenterColored(img, i18n.T("app.title"), m.big, 38, color.RGBA{
		255, 255, 255,
		uint8(180 + math.Sin(m.ft*4)*60),
	})
	m.drawDuck(img)
	m.drawBtn(img, m.play, i18n.T("menu.play"), 0.15)
	m.drawBtn(img, m.shop, i18n.T("menu.shop"), 0.25)
	m.drawBtn(img, m.opt, i18n.T("menu.options"), 0.35)
	m.drawBtn(img, m.quit, i18n.T("menu.quit"), 0.45)
	m.fx.Draw(img)

	switch m.state {	case stSettings:
		m.settingsDraw(img)
	case stQuit:
		m.quitDraw(img)
	case stMode:
		m.modeDraw(img)
	}
	m.toast.Draw(img, i18n.T("play.soon"), m.font)
}

// Patão na esquerda: entra voando da esquerda com pop,
// depois só flutuação vertical leve + sombra embaixo.
func (m *Menu) drawDuck(img *ebiten.Image) {
	if m.duck == nil {
		return
	}
	p := clamp01(m.ft / 0.7)
	e := backOut(p)
	bob := math.Sin(m.ft * 2)
	x := -220 + (340+220)*e
	y := 370 + bob*10
	m.drawGround(img, x, bob, p)
	m.duck.Draw(img, x, y, 1.25+1.25*e, 1)
}

// drawBtn desliza o botão da direita com atraso, escalonado.
func (m *Menu) drawBtn(img *ebiten.Image, b ui.Btn, label string, delay float64) {
	p := clamp01((m.ft - delay) / 0.5)
	b.X += (1 - easeOut(p)) * 700
	if b.Icon == ui.IconPlay {
		b.ISize = 32
	}
	b.Draw(img, label, m.font)
}

func clamp01(v float64) float64 { return max(0, min(1, v)) }
func easeOut(p float64) float64 { return 1 - (1-p)*(1-p)*(1-p) }
func backOut(p float64) float64 {
	c := 1.70158
	return 1 + (c+1)*(p-1)*(p-1)*(p-1) + c*(p-1)*(p-1)
}

func gold() color.Color   { return color.RGBA{255, 215, 0, 255} }
func shadow() color.Color { return color.RGBA{40, 40, 40, 255} }
