package menu

import (
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
	stMode
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
	// mode (online/lan)
	online, lan, modeX ui.Btn
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
	m := &Menu{
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
		qX:   btn(848, 248, 44, 44, ui.IconNone),
		online: btn(400, 320, 200, 64, ui.IconGlobe),
		lan:    btn(660, 320, 200, 64, ui.IconLan),
		modeX:  btn(848, 248, 44, 44, ui.IconNone),
	}
	// Em breve: sem clique, apagados.
	m.online.Disabled = true
	m.shop.Disabled = true
	return m
}

func (m *Menu) Update() error {
	m.ensureMedia()
	m.ft += 1.0 / 60.0

	switch m.state {
	case stSettings:
		m.settingsUpdate()
	case stQuit:
		m.quitUpdate()
	case stMode:
		m.modeUpdate()
	default:
		m.mainUpdate()
	}
	return nil
}
