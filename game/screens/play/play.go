package play

import (
	"duckling/content"
	"duckling/engine/rendering"
	"duckling/game/map"
	"duckling/game/player"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play.go — tela do jogo: o pato anda no WASD, a câmera segue.
// LAN: o host simula todo mundo; o cliente manda input e interpola.

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
	toast            ui.Toast
}

func NewPlay(font *text.GoTextFace, next func(string)) (*Play, error) {
	mp, err := world.Open("spawn")
	if err != nil {
		return nil, err
	}
	p := &Play{next: next, font: font, mp: mp,
		yes:   ui.Btn{X: 380, Y: 390, W: 240, H: 64},
		no:    ui.Btn{X: 660, Y: 390, W: 240, H: 64},
		qx:    ui.Btn{X: 840, Y: 250, W: 44, H: 44},
		rem:   map[int]*actor.Duck{},
		rname: map[int]string{},
		toast: ui.Toast{Right: true},
	}
	sx, sy := mp.Spawn()
	p.d = actor.New(assets.Images, sx, sy, actor.Cast["duck"])
	p.sess = session.Get()
	// Intro: nasce mostrando o mapa todo, desce até o pato.
	p.zoom = min(1280/float64(mp.W), 720/float64(mp.H))
	p.zFrom = p.zoom
	p.follow()
	return p, nil
}

func (p *Play) Update() error {
	p.syncSession()
	if p.confirm {
		p.confirmUpdate()
		return nil
	}
	// Pato carrega lazy: imagens só existem após o loading.
	if !p.d.Ok() {
		p.d.Load(assets.Images)
	}
	x0, y0, x1, y1 := p.bounds()
	p.d.Update(x0, y0, x1, y1, p.mp.Solids, actor.LocalIn())
	p.netUpdate(x0, y0, x1, y1)
	p.toast.Update()
	p.puff.Update()
	p.follow()
	if p.leaveIfBroken() {
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		p.confirm = true
	}
	// F3 liga o raio-X das colisões.
	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		p.dbg = !p.dbg
	}
	keys := []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5}
	for i, k := range keys {
		if inpututil.IsKeyJustPressed(k) && i < len(p.mp.Areas) {
			p.area = i
			r := p.mp.Areas[i].Rect
			p.d.Px, p.d.Py = float64(r.Min.X+r.Max.X)/2, float64(r.Min.Y+r.Max.Y)/2
			p.follow()
		}
	}
	return nil
}
