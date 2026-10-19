package play

import (
	"fmt"

	"duckling/engine/i18n"
	"duckling/game/match"

	"github.com/hajimehoshi/ebiten/v2"
)

// play_host.go — painel do dono: COMEÇAR (trancado até o mínimo)
// no lobby e ENCERRAR durante a partida.

func (p *Play) hostUpdate() {
	s := p.sess
	if s == nil || !s.IsHost() {
		return
	}
	switch p.phase {
	case match.Lobby:
		p.startB.Disabled = s.Count() < s.Min()
		if !p.startB.Disabled && p.startB.Update() {
			p.startMatch()
		}
	case match.Playing:
		if p.endB.Update() {
			s.Finish()
		}
	}
}

func (p *Play) hostDraw(img *ebiten.Image) {
	s := p.sess
	if s == nil || !s.IsHost() {
		return
	}
	switch p.phase {
	case match.Lobby:
		label := fmt.Sprintf("%s  %d/%d", i18n.T("match.start"), s.Count(), s.Min())
		p.startB.Draw(img, label, p.font)
	case match.Playing:
		p.endB.Draw(img, i18n.T("match.stop"), p.font)
	}
}

// startMatch sorteia o plano e abre a partida (host).
func (p *Play) startMatch() {
	plan := match.Build(p.sess.PlanIDs(), p.mp.Spawns(), p.mp.Tasks, p.sess.Impostors())
	p.sess.Start(plan)
	// O local é posicionado pelo syncPhase; os remotos vão na hora.
	for id, at := range plan.Spawn {
		if id == 0 {
			continue
		}
		p.remote(id).Snap(at.X, at.Y, 0, 0, false)
	}
}
