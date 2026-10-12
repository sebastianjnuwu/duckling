package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// icon_font.go — ícones pela fonte Material (só o nome, sem
// arquivo por ícone). Precedência: PNG custom, fonte, procedural.

// iconRunes liga o enum ao codepoint (verificado na sheet).
var iconRunes = map[Icon]rune{
	IconPlay:  0xE037,
	IconShop:  0xE8CC,
	IconGear:  0xE8B8,
	IconDoor:  0xE879,
	IconNote:  0xE405,
	IconSpk:   0xE050,
	IconGlobe: 0xE80B,
	IconLan:   0xE157,
	IconLock:  0xE897,
	IconPlus:    0xE145,
	IconMinus:   0xE15B,
	IconUser:    0xE7FD,
	IconCopy:    0xE14D,
	IconRefresh: 0xE863,
	IconUp:      0xE5D8,
	IconDown:    0xE5DB,
	IconLeft:    0xE5C4,
	IconRight:   0xE5C5,
	IconJump:    0xE256,
	IconFly:     0xE539,
}

var (
	iconFontSrc func() *text.GoTextFaceSource
	iconFaces   = map[float64]*text.GoTextFace{}
)

// ProvideIconFont entrega a fonte (chamar 1x no boot).
func ProvideIconFont(src func() *text.GoTextFaceSource) { iconFontSrc = src }

// drawFont desenha o glifo centrado na caixa, tintado.
func drawFont(dst *ebiten.Image, ic Icon, x, y, size float64, c color.Color) bool {
	r, ok := iconRunes[ic]
	if !ok || iconFontSrc == nil {
		return false
	}
	f, ok := iconFaces[size]
	if !ok {
		src := iconFontSrc()
		if src == nil {
			return false
		}
		f = &text.GoTextFace{Source: src, Size: size}
		iconFaces[size] = f
	}
	w, h := text.Measure(string(r), f, 0)
	op := &text.DrawOptions{}
	op.GeoM.Translate(x+(size-w)/2, y+(size-h)/2)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(dst, string(r), f, op)
	return true
}
