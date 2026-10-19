package play

import (
	"fmt"
	"image/color"

	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/engine/save"
	"duckling/game/match"
)

// play_task_fx.go — fechamento da tarefa: avisa o host e paga
// (moedas, som, toast e penas).

// finishTasks conta quem completou e avisa o host.
func (p *Play) finishTasks() {
	for i, st := range p.ts {
		if st.Done() && !p.done[i] {
			p.done[i] = true
			p.report(i)
			p.reward(i)
		}
	}
}

// report manda a conclusão: o host conta a dele, o cliente avisa.
func (p *Play) report(i int) {
	if p.sess == nil {
		return
	}
	if p.sess.IsHost() {
		p.sess.DoneTask(p.sess.ID(), i)
		return
	}
	p.sess.SendTask(i)
}

const (
	coinShort = 2 // tarefa curta
	coinLong  = 5 // tarefa longa
)

// reward paga a tarefa concluída (impostor não ganha moeda).
func (p *Play) reward(i int) {
	if p.imp() {
		return
	}
	n := coinShort
	if p.mp.Tasks[i].Long {
		n = coinLong
	}
	save.AddCoins(n)
	save.Store()
	_ = audio.Play("bumped")
	at := p.mp.Tasks[i].Center()
	sx := (float64(at.X) - p.camX) * p.zoom
	sy := (float64(at.Y) - p.camY) * p.zoom
	p.puff.Burst(sx, sy-16, 12, ui.Gold, ui.Cream, color.RGBA{230, 120, 60, 255})
	p.toast.Show(fmt.Sprintf("%s +%d %s", i18n.T("task.done"), n, i18n.T("task.coins")))
}

// imp diz se o pato local é impostor.
func (p *Play) imp() bool {
	if p.sess == nil {
		return false
	}
	r, ok := p.sess.Role(p.sess.ID())
	return ok && r == match.Imp
}
