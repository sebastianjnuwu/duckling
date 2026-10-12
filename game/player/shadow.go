package actor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// shadow.go — sombra blob em 3 faixas de altura (perto, médio,
// alto): quanto mais alto, menor e mais clara. Arredondada em
// 3 fileiras de pixel pra combinar com a arte do jogo.
// x, y = pés na tela; s = zoom; relZ = altura acima do piso;
// ground = elevação do piso onde a sombra projeta.
func drawShadow(img *ebiten.Image, x, y, s, relZ, ground float64) {
	sc, a := 0.7, uint8(35)
	switch {
	case relZ >= 22:
		sc, a = 0.45, 40
	case relZ >= 10:
		sc, a = 0.7, 70
	case relZ >= 0.5:
		sc, a = 0.9, 75
	}
	w, h := 24*s*sc, 7*s*sc
	gy := y - ground*s + 8*s
	c := color.RGBA{0, 0, 0, a}
	ebitenutil.DrawRect(img, x-w*0.31, gy-h/2, w*0.62, h, c)
	ebitenutil.DrawRect(img, x-w/2, gy-h*0.32, w, h*0.64, c)
}
