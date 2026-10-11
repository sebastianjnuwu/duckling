package play

import (
	"image/color"

	"duckling/content"
	"duckling/game/player"
	"duckling/engine/rendering"
	"duckling/game/map"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play.go — tela do jogo: o pato anda no WASD, a câmera segue.
// ESC volta, 1-5 pula entre as áreas do mapa.

type Play struct {
	next func(string)
	font *text.GoTextFace
	mp   *world.Map
	d    *actor.Duck
	npcs []*actor.NPC
	area int
	dbg  bool
	confirm bool
	yes, no, qx ui.Btn
	zoom, zFrom, zT float64
	camX, camY, camZ float64
	puff ui.Sparks
}

func NewPlay(font *text.GoTextFace, next func(string)) (*Play, error) {
	mp, err := world.Open("spawn")
	if err != nil {
		return nil, err
	}
	p := &Play{next: next, font: font, mp: mp,
		yes: ui.Btn{X: 380, Y: 390, W: 240, H: 64},
		no:  ui.Btn{X: 660, Y: 390, W: 240, H: 64},
		qx:  ui.Btn{X: 840, Y: 250, W: 44, H: 44},
	}
	sx, sy := mp.Spawn()
	p.d = actor.New(assets.Images["sprites/duck/yellow/duck-4x"], sx, sy, actor.Cast["duck"])
	p.spawnNPCs(sx, sy, assets.Images["sprites/duck/red/duck-4x"])
	// Intro: nasce mostrando o mapa todo, desce até o pato.
	p.zoom = min(1280/float64(mp.W), 720/float64(mp.H))
	p.zFrom = p.zoom
	p.follow()
	return p, nil
}

func (p *Play) Update() error {
	// ESC no jogo abre confirmação em vez de sair direto.
	if p.confirm {
		p.confirmUpdate()
		return nil
	}
	// Pato carrega lazy: imagens só existem após o loading.
	if !p.d.Ok() {
		p.d.Load(assets.Images["sprites/duck/yellow/duck-4x"])
	}
	x0, y0, x1, y1 := p.bounds()
	p.d.Update(x0, y0, x1, y1, p.mp.Solids)
	red := assets.Images["sprites/duck/red/duck-4x"]
	for _, n := range p.npcs {
		if !n.D.Ok() {
			n.D.Load(red)
		}
		n.Update(x0, y0, x1, y1, p.mp.Solids)
		if n.Bump(p.d) {
			mx, my := (p.d.Px+n.D.Px)/2, (p.d.Py+n.D.Py)/2
			p.puff.Burst((mx-p.camX)*p.zoom, (my-p.camY)*p.zoom, 14,
				color.RGBA{235, 240, 235, 255},
				color.RGBA{250, 220, 120, 255})
		}
	}
	p.puff.Update()
	p.follow()
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
