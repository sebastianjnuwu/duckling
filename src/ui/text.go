package ui

import (
	"image/color"

	"duckling/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// text.go — fonte do FS embutido + texto centralizado. Todo texto
// visível passa por aqui (ou DrawLeft) com chave i18n, sem hardcode.

// Face monta a fonte pelo nome do arquivo (ex: "PressStart2P").
func Face(name string, size float64) *text.GoTextFace {
	src, ok := assets.Fonts[name]
	if !ok {
		return nil
	}
	return &text.GoTextFace{Source: src, Size: size}
}

func DrawCenter(img *ebiten.Image, s string, f *text.GoTextFace, y float64) {
	DrawCenterColored(img, s, f, y, color.White)
}

func DrawCenterColored(img *ebiten.Image, s string, f *text.GoTextFace, y float64, c color.Color) {
	if f == nil {
		return
	}
	w, _ := text.Measure(s, f, 0)
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(img.Bounds().Dx())/2-w/2, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(img, s, f, op)
}
