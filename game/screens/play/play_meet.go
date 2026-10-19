package play

import (
	"duckling/engine/audio"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// play_meet.go — a reunião por cima do mundo congelado: votar,
// falar (frases prontas) e avisar o resultado da ejeção.

// meetUpdate atende os cliques da reunião e o aviso de início.
func (p *Play) meetUpdate() {
	s := p.sess
	if s == nil {
		return
	}
	m := s.Meet()
	if e, ok := s.TakeEject(); ok {
		p.ejectFx(e)
	}
	if !m.Active {
		p.meetOn = false
		return
	}
	p.confirm = false
	if !p.meetOn {
		p.meetOn = true
		_ = audio.Play("bumped")
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		p.meetClick(m)
	}
}

func (p *Play) meetClick(m session.MeetInfo) {
	x, y := ebiten.CursorPosition()
	fx, fy := float64(x), float64(y)
	if !p.canAct() {
		return
	}
	if m.Open {
		if inRect(fx, fy, mtSkipX, mtSkipY, mtSkipW, mtSkipH) {
			p.doVote(session.SkipVote)
			return
		}
		roster := p.meetRoster()
		for i, e := range roster {
			if e.ID == p.sess.ID() || e.Ghost {
				continue
			}
			rx, ry, rw, rh := meetRowRect(i, len(roster))
			if inRect(fx, fy, rx, ry, rw, rh) {
				p.doVote(e.ID)
				return
			}
		}
	}
	for i := range session.SayKeys {
		sx, sy, sw, sh := meetSayRect(i)
		if inRect(fx, fy, sx, sy, sw, sh) {
			p.doSay(session.SayKeys[i])
			return
		}
	}
}

// doVote manda o voto (host aplica direto).
func (p *Play) doVote(id int) {
	if p.sess.IsHost() {
		p.sess.Vote(p.sess.ID(), id)
		return
	}
	p.sess.SendVote(id)
}

// doSay manda a frase pronta.
func (p *Play) doSay(key string) {
	if p.sess.IsHost() {
		p.sess.Say(p.sess.ID(), key)
		return
	}
	p.sess.SendSay(key)
}
