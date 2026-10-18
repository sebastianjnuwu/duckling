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

// menu_room_join_draw.go — o modal de entrar: sala escolhida,
// nome do jogador e senha (quando a sala é privada).

func (l *RoomState) joinDraw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.Dim(img)
	ui.Panel(img, rpX, rpY, rpW, rpH)
	l.x.Draw(img, "X", l.font)
	ui.Title(img, ui.IconLan, i18n.T("lobby.join"), l.big, 120, ui.Gold)
	left(img, fmt.Sprintf("%s  %d/%d", l.sel.name, l.sel.n, l.sel.max), l.font, rlX, 170)
	l.fieldAt(img, jNameY, i18n.T("lobby.name"), l.pname, l.jFocus == 1)
	if l.sel.priv {
		l.fieldAt(img, jPassY, i18n.T("lobby.password"), l.jPass, l.jFocus == 2)
	}
	l.joinB.Draw(img, i18n.T("lobby.join"), l.font)
	l.backJ.Draw(img, i18n.T("shop.back"), l.font)
	l.toast.Draw(img, l.font)
}

// fieldAt desenha rótulo + caixa editável (padrão do nome da sala).
func (l *RoomState) fieldAt(img *ebiten.Image, y float64, label, val string, focus bool) {
	ui.DrawIcon(img, ui.IconUser, rlX-32, y+16, cream())
	left(img, label, l.font, rlX, y+19)
	bc := color.Color(color.RGBA{48, 66, 100, 255})
	if focus {
		bc = gold()
	}
	ebitenutil.DrawRect(img, rcX, y+4, 240, 48, bc)
	ebitenutil.DrawRect(img, rcX+4, y+8, 232, 40, color.RGBA{16, 16, 30, 255})
	t := val
	if focus {
		t += "_"
	}
	left(img, t, l.font, rcX+12, y+19)
}
