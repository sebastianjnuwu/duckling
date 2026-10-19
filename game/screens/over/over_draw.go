package over

import (
	"fmt"
	"image/color"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/game/match"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
)

// over_draw.go — fundo escurecido, painel e o título do desfecho.

var (
	winCrew = color.RGBA{120, 210, 150, 255}
	winImp  = color.RGBA{230, 120, 120, 255}
	cream   = color.RGBA{235, 240, 235, 255}
	dimmer  = color.RGBA{132, 142, 158, 210}
)

func (o *Over) Draw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.Dim(img)
	ui.Panel(img, 150, 60, 980, 640)
	s := o.sess
	if s == nil {
		return
	}
	o.head(img, s)
	o.roles(img, s)
	o.votes(img, s)
	o.footer(img, s)
}

// head desenha vencedor, motivo e os números da partida.
func (o *Over) head(img *ebiten.Image, s *session.Session) {
	r := s.Result()
	title, col := i18n.T("over.stop"), ui.Gold
	switch r.Winner {
	case match.WinCrew:
		title, col = i18n.T("over.crew"), winCrew
	case match.WinImp:
		title, col = i18n.T("over.imp"), winImp
	}
	ui.DrawCenterColored(img, title, o.big, 92, col)
	if r.Reason != "" {
		ui.DrawCenterColored(img, i18n.T("over.why."+r.Reason), o.font, 134, cream)
	}
	done, total := s.GlobalTasks()
	line := fmt.Sprintf("%s %d/%d    %s %d",
		i18n.T("over.tasks"), done, total, i18n.T("over.kills"), s.Kills())
	ui.DrawCenterColored(img, line, o.face(12), 170, dimmer)
}
