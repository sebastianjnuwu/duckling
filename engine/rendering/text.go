package ui

import (
	"image/color"
	"strings"

	"duckling/content"

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

// wrap quebra em linhas de até maxW.
func wrap(msg string, f *text.GoTextFace, maxW float64) []string {
	var out []string
	cur := ""
	for _, w := range strings.Fields(msg) {
		t := strings.TrimSpace(cur + " " + w)
		if tw, _ := text.Measure(t, f, 0); tw > maxW && cur != "" {
			out = append(out, cur)
			cur = w
		} else {
			cur = t
		}
	}
	return append(out, cur)
}

// Title desenha título com ícone: grupo centralizado, ícone
// do tamanho da fonte (24 a 40).
func Title(img *ebiten.Image, ic Icon, s string, f *text.GoTextFace, y float64, c color.Color) {
	if f == nil {
		return
	}
	is := min(max(f.Size, 24), 40)
	w, h := text.Measure(s, f, 0)
	x0 := float64(img.Bounds().Dx())/2 - (w+12+is)/2
	DrawIconSized(img, ic, x0, y+(h-is)/2, is, c)
	op := &text.DrawOptions{}
	op.GeoM.Translate(x0+is+12, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(img, s, f, op)
}
