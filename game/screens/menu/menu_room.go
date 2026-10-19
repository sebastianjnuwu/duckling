package menu

import (
	"sync"

	"duckling/engine/rendering"
	world "duckling/game/map"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// lobby.go — sala de espera: browse (tabelinha de salas da LAN) e
// config (min/max, senha, lista). A rede preenche a tabela.

type room struct {
	n, max int
	priv   bool
	ip     string
	name   string
}

type RoomState struct {
	next                           func(string)
	onBack                         func()
	font, big                      *text.GoTextFace
	tab                            int
	page                           int
	mu                             sync.Mutex
	scan                           chan struct{}
	rooms                          []room
	x, prevB, nextB                ui.Btn
	rowB                           [4]ui.Btn
	create, back, backB            ui.Btn
	reloadB                        ui.Btn
	min, max, top                  int
	imp                            int
	priv                           bool
	pass                           string
	name                           string
	nameFocus                      bool
	toast                          ui.Toast
	start                          ui.Btn
	minDec, minInc, maxDec, maxInc ui.Btn
	impDec, impInc                 ui.Btn
	privT                          ui.Toggle
	regen, nameBox, copyB          ui.Btn
	// join (tab 2)
	sel          room
	pname        string
	jPass        string
	jFocus       int
	joinB, backJ ui.Btn
}

func NewRoomState(font *text.GoTextFace, next func(string)) (*RoomState, error) {
	top := 16
	if mp, err := world.Open("spawn"); err == nil {
		top = mp.MaxPlayers()
	}
	l := &RoomState{next: next, font: font, min: 4, max: top, top: top, imp: 0,
		toast:   ui.Toast{Right: true},
		onBack:  func() {},
		big:     &text.GoTextFace{Source: font.Source, Size: 34},
		pass:    genPass(),
		name:    "CURRAL",
		pname:   "DUCK",
		nameBox: btn(rcX, rowY(0)+4, 240, 48, ui.IconNone),
		create:  btn(390, 450, 240, 56, ui.IconPlus),
		back:    btn(690, rBtnY, 170, 56, ui.IconDoor),
		backB:   btn(660, 450, 180, 56, ui.IconDoor),
		x:       btn(rpX+rpW-80, rpY+20, 44, 44, ui.IconNone),
		reloadB: btn(rpX+rpW-150, rpY+20, 48, 48, ui.IconRefresh),
		prevB:   btn(320, 455, 56, 48, ui.IconNone),
		nextB:   btn(904, 455, 56, 48, ui.IconNone),
		start:   btn(420, rBtnY, 250, 56, ui.IconPlay),
		joinB:   btn(360, rBtnY, 240, 56, ui.IconPlay),
		backJ:   btn(660, rBtnY, 170, 56, ui.IconDoor),
		minDec:  btn(rcX, rowY(1), 56, 56, ui.IconMinus),
		minInc:  btn(rcX+200, rowY(1), 56, 56, ui.IconPlus),
		maxDec:  btn(rcX, rowY(2), 56, 56, ui.IconMinus),
		maxInc:  btn(rcX+200, rowY(2), 56, 56, ui.IconPlus),
		impDec:  btn(rcX, rowY(3), 56, 56, ui.IconMinus),
		impInc:  btn(rcX+200, rowY(3), 56, 56, ui.IconPlus),
		privT:   ui.Toggle{X: rcX, Y: rowY(4) + 4, W: 120, H: 48},
		regen:   btn(rcX+150, rowY(5)+4, 48, 48, ui.IconRefresh),
		copyB:   btn(rcX+210, rowY(5)+4, 48, 48, ui.IconCopy),
	}
	// Começa vazio: sala só existe se alguém criar (a rede preenche).
	for k := range l.rowB {
		l.rowB[k] = btn(320, float64(220+k*70), 640, 56, ui.IconLan)
	}
	return l, nil
}
