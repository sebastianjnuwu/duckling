package menu

import (
	"fmt"
	"image/color"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// menu_room_draw.go — novo jogo na grade (menu_room_layout.go):
// fileiras uniformes, senha estilo código de sala.

func (l *RoomState) Draw(img *ebiten.Image) {
	if l.tab == 0 {
		l.roomsDraw(img)
		return
	}
	if l.tab == 2 {
		l.joinDraw(img)
		return
	}
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.Dim(img)
	ui.Panel(img, rpX, rpY, rpW, rpH)
	l.x.Draw(img, "X", l.font)
	ui.Title(img, ui.IconPlus, i18n.T("lobby.title"), l.big, 125, ui.Gold)
	l.nameRow(img)
	l.stepRow(img, 1, i18n.T("lobby.min"), fmt.Sprint(l.min), &l.minDec, &l.minInc, l.canMinDec(), l.canMinInc())
	l.stepRow(img, 2, i18n.T("lobby.max"), fmt.Sprint(l.max), &l.maxDec, &l.maxInc, l.canMaxDec(), l.canMaxInc())
	l.stepRow(img, 3, i18n.T("lobby.impostors"), fmt.Sprint(l.imp), &l.impDec, &l.impInc, l.canImpDec(), l.canImpInc())
	y := rowY(4)
	ui.DrawIcon(img, ui.IconLock, rlX-32, y+16, cream())
	if l.priv {
		left(img, i18n.T("lobby.private"), l.font, rlX, y+19)
	} else {
		left(img, i18n.T("lobby.public"), l.font, rlX, y+19)
	}
	l.privT.Draw(img)
	if l.priv {
		y := rowY(5)
		ui.DrawIcon(img, ui.IconLock, rlX-32, y+16, cream())
		left(img, i18n.T("lobby.password"), l.font, rlX, y+19)
		pill(img, 640, y+4, 140, 48)
		num(img, l.pass, l.font, 710, y+19)
		l.regen.Draw(img, "", l.font)
		l.copyB.Draw(img, "", l.font)
	}
	l.start.Draw(img, i18n.T("lobby.create"), l.font)
	l.back.Draw(img, i18n.T("shop.back"), l.font)
	l.toast.Draw(img, l.font)
}

// nameRow: rótulo + campo editável + contador.
func (l *RoomState) nameRow(img *ebiten.Image) {
	y := rowY(0)
	ui.DrawIcon(img, ui.IconNote, rlX-32, y+16, cream())
	left(img, i18n.T("lobby.name"), l.font, rlX, y+19)
	bc := color.Color(color.RGBA{48, 66, 100, 255})
	if l.nameFocus {
		bc = gold()
	}
	ebitenutil.DrawRect(img, rcX, y+4, 240, 48, bc)
	ebitenutil.DrawRect(img, rcX+4, y+8, 232, 40, color.RGBA{16, 16, 30, 255})
	t := l.name
	if l.nameFocus {
		t += "_"
	}
	left(img, t, l.font, rcX+12, y+19)
	left(img, fmt.Sprintf("%d/8", len([]rune(l.name))), l.font, rcX+250, y+19)
}

// stepRow: rótulo + [-] valor [+] com pílula.
func (l *RoomState) stepRow(img *ebiten.Image, i int, label, val string, dec, inc *ui.Btn, canDec, canInc bool) {
	y := rowY(i)
	ui.DrawIcon(img, ui.IconUser, rlX-32, y+16, cream())
	left(img, label, l.font, rlX, y+19)
	if canDec {
		dec.Draw(img, "", l.font)
	}
	pill(img, rcX+66, y+2, 124, 52)
	num(img, val, l.font, rcX+128, y+19)
	if canInc {
		inc.Draw(img, "", l.font)
	}
}
