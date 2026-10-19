package play

import (
	"duckling/engine/input"
)

// play_report.go — a tecla de interagir também serve pro mundo:
// corpo por perto = reportar; sino do celeiro = reunião de
// emergência (1 por jogador). Sem alvo, a tecla é da tarefa.

const (
	bodyReach = 74.0
	bellReach = 58.0
)

// actUpdate divide a tecla de interagir entre reportar e a tarefa.
func (p *Play) actUpdate() {
	if !p.reportUpdate() {
		p.taskUpdate(1.0 / 60.0)
	}
	p.killUpdate()
}

// reportUpdate devolve true quando gastou a tecla num alvo.
func (p *Play) reportUpdate() bool {
	if p.sess == nil || !p.canAct() || !input.JustPressed(input.Interact) {
		return false
	}
	if id, ok := p.nearBody(); ok {
		p.doReport(id)
		return true
	}
	if p.bellUsed || !p.nearBell() {
		return false
	}
	p.bellUsed = true
	p.doBell()
	return true
}

// nearBody acha o corpo mais perto do pato local.
func (p *Play) nearBody() (int, bool) {
	for _, b := range p.sess.Bodies() {
		dx, dy := float64(b.X)-p.d.Px, float64(b.Y)-p.d.Py
		if dx*dx+dy*dy <= bodyReach*bodyReach {
			return b.ID, true
		}
	}
	return 0, false
}

// nearBell diz se o pato local está no sino.
func (p *Play) nearBell() bool {
	for _, b := range p.mp.Bells {
		dx, dy := float64(b.X)-p.d.Px, float64(b.Y)-p.d.Py
		if dx*dx+dy*dy <= bellReach*bellReach {
			return true
		}
	}
	return false
}

// doReport manda o corpo pro host (host aplica direto).
func (p *Play) doReport(id int) {
	if p.sess.IsHost() {
		p.sess.Report(p.sess.ID(), id)
		return
	}
	p.sess.SendReport(id)
}

// doBell pede a reunião de emergência.
func (p *Play) doBell() {
	if p.sess.IsHost() {
		p.sess.CallBell(p.sess.ID())
		return
	}
	p.sess.SendBell()
}
