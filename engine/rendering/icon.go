package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// icon.go — todo ícone vem da fonte (icon_font.go).

type Icon int

const (
	IconNone Icon = iota
	IconPlay
	IconShop
	IconGear
	IconDoor
	IconNote
	IconSpk
	IconGlobe
	IconLan
	IconLock
	IconPlus
	IconMinus
	IconUser
	IconCopy
	IconRefresh
	IconUp
	IconDown
	IconLeft
	IconRight
	IconJump
	IconFly
	IconHand
	IconKill
)

// DrawIcon desenha o ícone da fonte numa caixa 24x24 em (x,y).
func DrawIcon(img *ebiten.Image, ic Icon, x, y float64, c color.Color) {
	drawFont(img, ic, x, y, 24, c)
}
