package menu

import (
	"duckling/content"
	"duckling/engine/audio"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// menu.go — principal: título, pato voando, 3 botões. Estados:
// main, settings (modal), quit (confirmação). ESC volta/sai.

// Layout lógico 1280x720 (igual ao screen.go).
const (
	viewW = 1280.0
	viewH = 720.0
)

const (
	stMain = iota
	stSettings
	stQuit
)

type Menu struct {
	next                  func(string)
	font                  *text.GoTextFace
	big                   *text.GoTextFace
	sml                   *text.GoTextFace
	play, shop, opt, quit ui.Btn
	set                   Settings
	// quit
	qYes, qNo, qX ui.Btn
	state         int
	duck          *ui.Strip
	ft            float64
	bgm           bool
	toast         ui.Toast
	fx            ui.Sparks
}

func btn(x, y, w, h float64, ic ui.Icon) ui.Btn { return ui.Btn{X: x, Y: y, W: w, H: h, Icon: ic} }

func NewMenu(font *text.GoTextFace, next func(string)) *Menu {
	cx := viewW / 2
	return &Menu{
		next: next,
		font: font,
		big:  &text.GoTextFace{Source: font.Source, Size: 95},
		sml:  &text.GoTextFace{Source: font.Source, Size: 24},
		// patão na esquerda, pilha uniforme na direita:
		// 4x400x72 = mesma altura do pato, centro em 370.
		play: btn(740, 208, 400, 72, ui.IconPlay),
		shop: btn(740, 292, 400, 72, ui.IconShop),
		opt:  btn(740, 376, 400, 72, ui.IconGear),
		quit: btn(740, 460, 400, 72, ui.IconDoor),
		set:  newSettings(),
		qYes: btn(cx-260, 390, 240, 64, ui.IconNone),
		qNo:  btn(cx+20, 390, 240, 64, ui.IconNone),
		qX:   btn(860, 242, 48, 48, ui.IconNone),
	}
}

func (m *Menu) Update() error {
	// Música entra no menu (lazy: garante loop de áudio rodando).
	if !m.bgm {
		m.bgm = true
		_ = audio.PlayLoop("Loop")
	}
	// Pato carrega lazy: imagens só existem após o loading.
	if m.duck == nil {
		if img := assets.Images["sprites/duck/yellow/duck-4x"]; img != nil {
			m.duck = ui.NewDuck(img, 1, 128, 24)
		}
	} else {
		m.duck.Tick()
	}
	m.ft += 1.0 / 60.0

	switch m.state {
	case stSettings:
		m.settingsUpdate()
	case stQuit:
		m.quitUpdate()
	default:
		m.mainUpdate()
	}
	return nil
}
