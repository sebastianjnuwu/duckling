package actor

import (
	"image/color"

	"duckling/engine/rendering"
)

// color.go — cor do pato: matiz sobre a arte amarela (a loja
// troca por sprite própria na Fase 7).

type Color int

const (
	Yellow Color = iota
	Red
	Orange
	Green
	Cyan
	Blue
	Purple
	Pink
	ColorCount
)

var colorNames = [ColorCount]string{
	"yellow", "red", "orange", "green",
	"cyan", "blue", "purple", "pink",
}

// colorHue é a rotação (radianos) do amarelo da arte.
var colorHue = [ColorCount]float64{
	0, -0.9, -0.35, 1.2, 2.25, 2.95, 4.0, 4.9,
}

// swatch é a cor do quadradinho da UI, não o matiz aplicado.
var swatch = [ColorCount]color.RGBA{
	{246, 208, 76, 255},
	{226, 84, 84, 255},
	{236, 140, 60, 255},
	{96, 190, 96, 255},
	{72, 196, 196, 255},
	{84, 132, 232, 255},
	{164, 104, 220, 255},
	{232, 120, 196, 255},
}

// ColorValid diz se c é uma cor do elenco.
func ColorValid(c int) bool { return c >= 0 && c < int(ColorCount) }

// Hue devolve a rotação de matiz (0 = amarelo da arte).
func (c Color) Hue() float64 {
	if !ColorValid(int(c)) {
		return 0
	}
	return colorHue[c]
}

// Name devolve a chave curta da cor (i18n: color.<name>).
func (c Color) Name() string {
	if !ColorValid(int(c)) {
		return colorNames[0]
	}
	return colorNames[c]
}

// RGB devolve a cor do quadradinho no lobby.
func (c Color) RGB() color.RGBA {
	if !ColorValid(int(c)) {
		return swatch[0]
	}
	return swatch[c]
}

// strips devolve as 12 tiras (nil antes do Load).
func (d *Duck) strips() []*ui.Strip {
	return []*ui.Strip{
		d.down, d.up, d.left, d.right,
		d.flyDown, d.flyUp, d.flyLeft, d.flyRight,
		d.jumpDown, d.jumpUp, d.jumpLeft, d.jumpRight,
	}
}

// Color devolve a cor atual do pato.
func (d *Duck) Color() Color { return d.col }

// SetColor troca a cor (matiz) das tiras já carregadas.
func (d *Duck) SetColor(c Color) {
	if !ColorValid(int(c)) {
		c = Yellow
	}
	if d.col == c {
		return
	}
	d.col = c
	for _, s := range d.strips() {
		if s != nil {
			s.Hue(c.Hue())
		}
	}
}
