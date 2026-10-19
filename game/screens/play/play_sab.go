package play

import (
	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/game/match"
	"duckling/game/session"
)

// play_sab.go — sabotagens do impostor: botões no HUD, pedido ao
// host e o aviso (toast/som) que todo mundo vê.

func (p *Play) sabUpdate() {
	if p.sess == nil || p.phase != match.Playing || !p.imp() || p.ghost() {
		p.sabLamp.Disabled, p.sabGate.Disabled = true, true
		return
	}
	cd := p.sess.SabotageCooldown()
	p.sabLamp.Disabled = cd > 0 || p.sess.LampLeft() > 0
	p.sabGate.Disabled = cd > 0 || p.sess.GateLeft() > 0
	if p.sabLamp.Update() {
		p.doSab(session.SabLamp)
	}
	if p.sabGate.Update() {
		p.doSab(session.SabGate)
	}
}

// doSab manda pro host (ou aplica, se for o host).
func (p *Play) doSab(kind string) {
	if p.sess.IsHost() {
		p.sess.Sabotage(p.sess.ID(), kind)
		return
	}
	p.sess.SendSab(kind)
}

// sabFx avisa com toast/som quando uma sabotagem começa.
func (p *Play) sabFx() {
	if p.sess == nil {
		return
	}
	lamp := p.sess.LampLeft() > 0
	gate := p.sess.GateLeft() > 0
	if lamp && !p.lampOn {
		p.toast.Show(i18n.T("sab.lamp.off"))
		_ = audio.Play("bumped")
	}
	if gate && !p.gateOn {
		p.toast.Show(i18n.T("sab.gate.closed"))
		_ = audio.Play("bumped")
	}
	p.lampOn, p.gateOn = lamp, gate
}
