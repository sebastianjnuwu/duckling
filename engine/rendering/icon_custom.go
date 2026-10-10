package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// icon_custom.go — ícones PNG no lugar dos procedurais.
// Jogue `images/icons/<nome>.png` BRANCO no content (nomes
// abaixo); sem arquivo, cai no procedural de sempre.

// iconFiles liga o enum ao caminho (chave do assets.Images).
var iconFiles = map[Icon]string{
	IconPlay:  "images/icons/play",
	IconShop:  "images/icons/shop",
	IconGear:  "images/icons/gear",
	IconDoor:  "images/icons/door",
	IconNote:  "images/icons/note",
	IconSpk:   "images/icons/spk",
	IconGlobe: "images/icons/globe",
	IconLan:   "images/icons/lan",
	IconLock:  "images/icons/lock",
	IconPlus:  "images/icons/plus",
	IconMinus: "images/icons/minus",
	IconUser:  "images/icons/user",
}

var (
	iconSrc func(string) *ebiten.Image
	custom  = map[Icon]*ebiten.Image{}
)

// ProvideIcons entrega o buscador de imagens (chamar 1x no boot).
func ProvideIcons(src func(string) *ebiten.Image) { iconSrc = src }

// customIcon devolve o PNG do ícone, ou nil (usa o procedural).
func customIcon(ic Icon) *ebiten.Image {
	if img, ok := custom[ic]; ok {
		return img
	}
	if iconSrc == nil {
		return nil
	}
	img := iconSrc(iconFiles[ic])
	if img != nil {
		custom[ic] = img
	}
	return img
}

// drawCustom desenha o PNG centrado na caixa, tintado.
func drawCustom(dst, src *ebiten.Image, x, y, size float64, c color.Color) {
	w, h := float64(src.Bounds().Dx()), float64(src.Bounds().Dy())
	s := min(size/w, size/h)
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterNearest}
	op.GeoM.Scale(s, s)
	op.GeoM.Translate(x+(size-w*s)/2, y+(size-h*s)/2)
	op.ColorScale.ScaleWithColor(c)
	dst.DrawImage(src, op)
}
