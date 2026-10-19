package actor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// hat_draw.go — chapéus em retângulos sobre o sprite (frame
// 128px, cabeça no topo em y=28). Trocar por arte não muda a API.

var (
	hatRed   = color.RGBA{206, 76, 76, 255}
	hatRedD  = color.RGBA{148, 46, 54, 255}
	hatGold  = color.RGBA{232, 190, 72, 255}
	hatGoldD = color.RGBA{166, 126, 50, 255}
	hatStraw = color.RGBA{228, 198, 116, 255}
)

// DrawHat desenha h na cabeça de um pato cujo frame está
// centrado em (x,y) com escala k (px de tela por px do frame).
// dir segue o pato (0 baixo, 1 cima, 2 esquerda, 3 direita);
// air desloca o chapéu no frame (voo sobe, pulo desce).
func DrawHat(img *ebiten.Image, h Hat, dir int, x, y, k, air float64) {
	hatRects(img, h, dir, x, y, k, air, 1)
}

// DrawHatGhost é o mesmo chapéu, translúcido (fantasma).
func DrawHatGhost(img *ebiten.Image, h Hat, dir int, x, y, k, air float64) {
	hatRects(img, h, dir, x, y, k, air, 0.45)
}

func hatRects(img *ebiten.Image, h Hat, dir int, x, y, k, air, alpha float64) {
	if h == HatNone || !HatValid(int(h)) || k <= 0 {
		return
	}
	cx, top := 62.0, 28.0+air
	switch dir {
	case 1:
		cx = 63
	case 2:
		cx = 55
	case 3:
		cx = 73
	}
	rect := func(fx, fy, fw, fh float64, c color.RGBA) {
		c.A = uint8(float64(c.A) * alpha)
		ebitenutil.DrawRect(img, x+(fx-64)*k, y+(fy-64)*k, fw*k, fh*k, c)
	}
	switch h {
	case HatCap:
		rect(cx-20, top-9, 40, 12, hatRed)
		rect(cx-25, top+1, 48, 6, hatRedD)
		rect(cx-3, top-11, 6, 3, hatGold)
	case HatStraw:
		rect(cx-18, top-10, 36, 14, hatStraw)
		rect(cx-29, top+1, 58, 7, hatStraw)
		rect(cx-18, top+1, 36, 3, hatGoldD)
	case HatCrown:
		rect(cx-16, top-2, 32, 10, hatGold)
		rect(cx-16, top-10, 7, 9, hatGold)
		rect(cx-4, top-14, 8, 13, hatGold)
		rect(cx+9, top-10, 7, 9, hatGold)
		rect(cx-2, top+1, 4, 4, hatRed)
	}
}

// hatAir desloca o chapéu conforme a ação (voo sobe, pulo desce).
func (d *Duck) hatAir() float64 {
	switch {
	case d.flying:
		return -8
	case d.z > d.ground+0.1:
		return 4
	}
	return 0
}

// drawHat usa o estado do pato pra acertar o chapéu na ação.
func (d *Duck) drawHat(img *ebiten.Image, x, y, k float64) {
	DrawHat(img, d.hat, d.dir, x, y, k, d.hatAir())
}
