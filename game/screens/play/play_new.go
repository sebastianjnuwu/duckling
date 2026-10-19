package play

import (
	"duckling/content"
	"duckling/engine/rendering"
	"duckling/game/map"
	"duckling/game/player"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play_new.go — monta a tela: mapa, pato, botões e câmera inicial.

func NewPlay(font *text.GoTextFace, next func(string)) (*Play, error) {
	mp, err := world.Open("spawn")
	if err != nil {
		return nil, err
	}
	p := &Play{next: next, font: font, mp: mp,
		yes:    ui.Btn{X: 380, Y: 390, W: 240, H: 64},
		no:     ui.Btn{X: 660, Y: 390, W: 240, H: 64},
		qx:     ui.Btn{X: 840, Y: 250, W: 44, H: 44},
		startB: ui.Btn{X: 520, Y: 22, W: 240, H: 48, Icon: ui.IconPlay},
		endB:   ui.Btn{X: 1050, Y: 22, W: 210, H: 44, Icon: ui.IconDoor},
		backB:  ui.Btn{X: 500, Y: 430, W: 280, H: 56, Icon: ui.IconPlay},
		rem:    map[int]*actor.Duck{},
		rname:  map[int]string{},
		rghost: map[int]bool{},
		toast:  ui.Toast{Right: true},
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
