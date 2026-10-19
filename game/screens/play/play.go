package play

import (
	"duckling/content"
	"duckling/engine/rendering"
	"duckling/game/map"
	"duckling/game/match"
	"duckling/game/player"
	"duckling/game/session"
	"duckling/game/task"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play.go — tela do jogo: o pato anda no WASD, a câmera segue.
// LAN: o host simula todo mundo; o cliente manda input e interpola.
// A fase (match) manda: lobby livre, partida, fim.

type Play struct {
	next             func(string)
	font             *text.GoTextFace
	mp               *world.Map
	d                *actor.Duck
	area             int
	dbg              bool
	confirm          bool
	yes, no, qx      ui.Btn
	zoom, zFrom, zT  float64
	camX, camY, camZ float64
	puff             ui.Sparks
	sess             *session.Session
	rem              map[int]*actor.Duck
	rname            map[int]string
	rghost           map[int]bool
	toast            ui.Toast
	phase            match.Phase
	placed           bool
	startB, endB     ui.Btn
	backB            ui.Btn
	ts               map[int]*task.State
	done             map[int]bool
	tids             []int
}

func (p *Play) Update() error {
	p.syncSession()
	p.syncPhase()
	if p.confirm {
		p.confirmUpdate()
		return nil
	}
	// Pato carrega lazy: imagens só existem após o loading.
	if !p.d.Ok() {
		p.d.Load(assets.Images)
	}
	x0, y0, x1, y1 := p.bounds()
	if p.canMove() {
		p.d.Update(x0, y0, x1, y1, p.mp.Solids, actor.LocalIn())
	}
	p.netUpdate(x0, y0, x1, y1)
	if p.phase == match.Playing && p.canMove() {
		p.taskUpdate(1.0 / 60.0)
	}
	p.toast.Update()
	p.puff.Update()
	p.follow()
	if p.leaveIfBroken() {
		return nil
	}
	if p.phase == match.End {
		return p.overUpdate()
	}
	p.hostUpdate()
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		p.confirm = true
	}
	// F3 liga o raio-X das colisões.
	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		p.dbg = !p.dbg
	}
	p.areaKeys()
	return nil
}
