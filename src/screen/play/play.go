package play

import (
	"image/color"

	"duckling/assets"
	"duckling/src/actor"
	"duckling/src/ui"
	"duckling/src/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// play.go — tela do jogo: o pato anda no WASD, a câmera segue.
// ESC volta, 1-5 pula entre as áreas do mapa.

type Play struct {
	next func(string)
	mp   *world.Map
	d    *actor.Duck
	npcs []*actor.NPC
	area int
	dbg  bool
	zoom, zFrom, zT float64
	camX, camY, camZ float64
	puff ui.Sparks
}

func NewPlay(next func(string)) (*Play, error) {
	mp, err := world.Load("maps/spawn/spawn.tmx")
	if err != nil {
		return nil, err
	}
	p := &Play{next: next, mp: mp}
	sx, sy := mp.Spawn()
	p.d = actor.New(assets.Images["sprites/duck/yellow/duck-4x"], sx, sy, actor.Cast["duck"])
	// Vermelhinhos passeando perto do spawn.
	red := assets.Images["sprites/duck/red/duck-4x"]
	for _, off := range [][2]float64{{60, 0}, {-50, 30}, {10, -60}} {
		p.npcs = append(p.npcs, actor.NewNPC(red, sx+off[0], sy+off[1], actor.Cast["red"]))
	}
	p.npcs[2].Still = true
	// Intro: nasce mostrando o mapa todo, desce até o pato.
	p.zoom = min(1280/float64(mp.W), 720/float64(mp.H))
	p.zFrom = p.zoom
	p.follow()
	return p, nil
}

func (p *Play) Update() error {
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
		p.next("menu")
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
