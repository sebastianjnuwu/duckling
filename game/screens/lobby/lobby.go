package lobby

import (
	"fmt"
	"math/rand/v2"

	"duckling/engine/audio"
	"duckling/engine/rendering"
	world "duckling/game/map"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// lobby.go — sala de espera: browse (tabelinha de salas) e
// config (min/max, senha, lista). Mock offline, rede depois.

type room struct {
	n, max int
	priv   bool
}

type Lobby struct {
	next func(string)
	font, big *text.GoTextFace
	tab  int
	page int
	rooms []room
	x, prevB, nextB ui.Btn
	rowB [4]ui.Btn
	create, back, backB ui.Btn
	min, max, top int
	imp int
	priv bool
	pass string
	name string
	nameFocus bool
	start ui.Btn
	minDec, minInc, maxDec, maxInc ui.Btn
	impDec, impInc ui.Btn
	privT ui.Toggle
	regen, nameBox ui.Btn
}

func btn(x, y, w, h float64, ic ui.Icon) ui.Btn { return ui.Btn{X: x, Y: y, W: w, H: h, Icon: ic} }

func NewLobby(font *text.GoTextFace, next func(string)) (*Lobby, error) {
	top := 16
	if mp, err := world.Open("spawn"); err == nil {
		top = mp.MaxPlayers()
	}
	l := &Lobby{next: next, font: font, min: 6, max: top, top: top, imp: 1,
		big: &text.GoTextFace{Source: font.Source, Size: 34},
		pass: fmt.Sprintf("%05d", rand.IntN(100000)),
		name: "CURRAL",
		nameBox: btn(600, 182, 240, 48, ui.IconNone),
		create: btn(400, 450, 220, 56, ui.IconPlus),
		back:   btn(680, 545, 180, 56, ui.IconDoor),
		backB:  btn(660, 450, 180, 56, ui.IconDoor),
		x:      btn(928, 128, 44, 44, ui.IconNone),
		prevB:  btn(320, 455, 56, 48, ui.IconNone),
		nextB:  btn(904, 455, 56, 48, ui.IconNone),
		start: btn(420, 545, 220, 56, ui.IconPlay),
		minDec: btn(640, 240, 56, 56, ui.IconMinus),
		minInc: btn(840, 240, 56, 56, ui.IconPlus),
		maxDec: btn(640, 300, 56, 56, ui.IconMinus),
		maxInc: btn(840, 300, 56, 56, ui.IconPlus),
		impDec: btn(640, 360, 56, 56, ui.IconMinus),
		impInc: btn(840, 360, 56, 56, ui.IconPlus),
		privT: ui.Toggle{X: 640, Y: 420, W: 120, H: 48},
		regen: btn(830, 472, 48, 48, ui.IconGear),
	}
	// Começa vazio: sala só existe se alguém criar (mock: CRIAR).
	for k := range l.rowB {
		l.rowB[k] = btn(320, float64(220+k*70), 640, 56, ui.IconLan)
	}
	return l, nil
}

func (l *Lobby) click() { _ = audio.Play("click") }

func (l *Lobby) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		l.nameFocus = false
		if l.tab == 0 {
			l.next("menu")
		} else {
			l.tab = 0
		}
		return nil
	}
	if l.tab == 0 {
		return l.roomsUpdate()
	}
	return l.configUpdate()
}
