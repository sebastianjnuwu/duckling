package lobby

import (
	"fmt"
	"image/color"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// lobby_draw.go — passo 2 com hierarquia: título ouro, valores em pílula.

func (l *Lobby) Draw(img *ebiten.Image) {
	if l.tab == 0 {
		l.roomsDraw(img)
		return
	}
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.Dim(img)
	ui.Panel(img, 280, 110, 720, 510)
	l.x.Draw(img, "X", l.font)
	ui.DrawCenterColored(img, i18n.T("lobby.title"), l.big, 135, gold())
	left(img, i18n.T("lobby.name"), l.font, 340, 199)
	bc := color.Color(color.RGBA{48, 66, 100, 255})
	if l.nameFocus {
		bc = gold()
	}
	ebitenutil.DrawRect(img, 600, 182, 240, 48, bc)
	ebitenutil.DrawRect(img, 604, 186, 232, 40, color.RGBA{16, 16, 30, 255})
	t := l.name
	if l.nameFocus {
		t += "_"
	}
	left(img, t, l.font, 612, 199)
	left(img, fmt.Sprintf("%d/8", len([]rune(l.name))), l.font, 850, 199)
	pill(img, 706, 242, 124, 52)
	pill(img, 706, 302, 124, 52)
	pill(img, 706, 362, 124, 52)
	ui.DrawIcon(img, ui.IconUser, 308, 256, cream())
	left(img, i18n.T("lobby.min"), l.font, 340, 259)
	if l.canMinDec() {
		l.minDec.Draw(img, "", l.font)
	}
	num(img, fmt.Sprint(l.min), l.font, 768, 259)
	if l.canMinInc() {
		l.minInc.Draw(img, "", l.font)
	}
	ui.DrawIcon(img, ui.IconUser, 308, 316, cream())
	left(img, i18n.T("lobby.max"), l.font, 340, 319)
	if l.canMaxDec() {
		l.maxDec.Draw(img, "", l.font)
	}
	num(img, fmt.Sprint(l.max), l.font, 768, 319)
	if l.canMaxInc() {
		l.maxInc.Draw(img, "", l.font)
	}
	ui.DrawIcon(img, ui.IconUser, 308, 366, cream())
	left(img, i18n.T("lobby.impostors"), l.font, 340, 369)
	if l.canImpDec() {
		l.impDec.Draw(img, "", l.font)
	}
	num(img, fmt.Sprint(l.imp), l.font, 768, 369)
	if l.canImpInc() {
		l.impInc.Draw(img, "", l.font)
	}
	ui.DrawIcon(img, ui.IconLock, 308, 426, cream())
	if l.priv {
		left(img, i18n.T("lobby.private"), l.font, 340, 429)
	} else {
		left(img, i18n.T("lobby.public"), l.font, 340, 429)
	}
	l.privT.Draw(img)
	if l.priv {
		ui.DrawIcon(img, ui.IconLock, 308, 481, cream())
		left(img, i18n.T("lobby.password"), l.font, 340, 484)
		pill(img, 690, 472, 120, 48)
		num(img, l.pass, l.font, 750, 484)
		l.regen.Draw(img, "", l.font)
	}
	l.start.Draw(img, i18n.T("lobby.create"), l.font)
	l.back.Draw(img, i18n.T("shop.back"), l.font)
}
