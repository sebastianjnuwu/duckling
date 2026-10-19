package play

import (
	"image/color"

	"duckling/engine/i18n"

	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// play_meet_chat.go — o log de falas e os 12 botões de frase
// pronta (fantasma não fala nem vota: botões apagados).

func (p *Play) meetChat(img *ebiten.Image, m session.MeetInfo) {
	p.textAt(img, i18n.T("meet.chat"), mtLogX, mtLogY-24, mtCream, 10)
	log := p.sess.Chats()
	if len(log) > mtLogLines {
		log = log[len(log)-mtLogLines:]
	}
	for i, c := range log {
		label := p.meetName(c.By) + ": " + i18n.T(c.Key)
		p.textAt(img, label, mtLogX, mtLogY+float64(i)*mtLogStep, mtCream, 10)
	}
	dead := !p.canAct()
	for i, key := range session.SayKeys {
		x, y, w, h := meetSayRect(i)
		face, col := color.RGBA{30, 36, 54, 235}, mtCream
		if dead {
			face, col = color.RGBA{22, 24, 32, 200}, mtDim
		}
		ebitenutil.DrawRect(img, x, y, w, h, face)
		p.centerText(img, i18n.T(key), x+w/2, y+10, col, 10)
	}
	_ = m
}
