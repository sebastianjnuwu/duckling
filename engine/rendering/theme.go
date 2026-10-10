package ui

import "image/color"

// theme.go — paleta central (o "CSS" do jogo, já que Ebiten
// desenha pixel em Go e não tem folha de estilo): muda aqui,
// muda em tudo que usa as vars. Migração aos poucos.

var (
	Ink    = color.RGBA{16, 16, 30, 255}
	Shadow = color.RGBA{0, 0, 0, 130}
	Gold   = color.RGBA{232, 190, 72, 255}
	Brass  = color.RGBA{170, 135, 60, 255}
	Cream  = color.RGBA{235, 240, 235, 255}
	Slate  = color.RGBA{38, 48, 74, 255}
	Gray   = color.RGBA{110, 115, 135, 255}
)

// bevel deriva claro/escuro da cor da face.
func bevel(face color.RGBA) (hi, mid, lo color.Color) {
	hi = color.RGBA{uint8(min(255, int(face.R)+90)), uint8(min(255, int(face.G)+90)), uint8(min(255, int(face.B)+90)), 255}
	mid = color.RGBA{uint8(max(0, int(face.R)-30)), uint8(max(0, int(face.G)-30)), uint8(max(0, int(face.B)-30)), 255}
	lo = color.RGBA{20, 20, 30, 255}
	return hi, mid, lo
}
