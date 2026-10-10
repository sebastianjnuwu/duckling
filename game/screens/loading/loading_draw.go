package loading

import (
	"fmt"
	"image/color"
	"math"

	"duckling/content"
	"duckling/engine/i18n"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (l *Loading) Draw(img *ebiten.Image) {
	h := float64(img.Bounds().Dy())

	if bg := assets.Images["images/bg"]; bg != nil {
		drawCover(img, bg)
	}

	font := l.Font()
	big := l.BigFont()

	title := i18n.T("app.title")

	if !l.finished {
		title += dots(l.ticks)
	}

	centerY := h/2 - big.Size/2

	if l.finished {
		centerY = 40.0
	}

	centerY -= l.pop * 25
	scale := 1 + l.pop*0.15

	drawTitleScaled(img, title, big, centerY, scale)

	if !l.finished {
		drawSpinnerPct(img, l.spin, fmt.Sprintf("%d%%", int(l.shown)), font)
	}
}

func (l *Loading) mouseOnTitle() bool {
	mx, my := ebiten.CursorPosition()
	w, h := ebiten.WindowSize()

	title := i18n.T("app.title")
	big := l.BigFont()

	tw, th := text.Measure(title, big, big.Size)
	cy := float64(h)/2 - big.Size/2

	return float64(mx) >= float64(w)/2-tw/2-20 &&
		float64(mx) <= float64(w)/2+tw/2+20 &&
		float64(my) >= cy-20 &&
		float64(my) <= cy+th+20
}

func dots(t float64) string {
	// Troca a cada 0.25s: ciclo cheio em 1s.
	return []string{"", ".", "..", "..."}[int(t/0.25)%4]
}

func drawTitleScaled(
	img *ebiten.Image,
	title string,
	font *text.GoTextFace,
	y, scale float64,
) {
	w := float64(img.Bounds().Dx())
	tw, _ := text.Measure(title, font, font.Size)
	x := w/2 - tw*scale/2

	drawScaled(img, title, font, x, y+7, scale, color.RGBA{40, 40, 40, 255})
	drawScaled(img, title, font, x, y+4, scale, color.RGBA{120, 120, 120, 255})
	drawScaled(img, title, font, x, y, scale, color.RGBA{255, 215, 0, 255})
	drawScaled(img, title, font, x, y-2, scale, color.RGBA{
		255, 255, 255,
		uint8(180 + math.Sin(float64(ebiten.Tick())/3)*60),
	})
}

func drawScaled(
	img *ebiten.Image,
	value string,
	font *text.GoTextFace,
	x, y, scale float64,
	c color.Color,
) {
	tw, _ := text.Measure(value, font, font.Size)

	op := &text.DrawOptions{}
	op.GeoM.Translate(-tw*scale/2, 0)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(tw*scale/2, 0)
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(c)

	text.Draw(img, value, font, op)
}

func drawSpinnerPct(
	img *ebiten.Image,
	angle float64,
	pct string,
	font *text.GoTextFace,
) {
	cx := 70.0
	cy := float64(img.Bounds().Dy()) - 70.0
	orbit := 45.0
	dotR := 4.0

	const trail = 8

	for i := 0; i < trail; i++ {
		a := angle - float64(i)*0.25

		drawCircle(
			img,
			cx+math.Cos(a)*orbit,
			cy+math.Sin(a)*orbit,
			dotR-float64(i)*0.3,
			color.RGBA{255, 215, 0, uint8(255 - i*28)},
		)
	}

	drawCircle(
		img,
		cx+math.Cos(angle)*orbit,
		cy+math.Sin(angle)*orbit,
		dotR,
		color.RGBA{255, 255, 255, 255},
	)

	tw, th := text.Measure(pct, font, font.Size)

	op := &text.DrawOptions{}
	op.GeoM.Translate(cx-tw/2, cy-th/2)
	op.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})

	text.Draw(img, pct, font, op)
}

func drawCover(img, src *ebiten.Image) {
	sw := float64(src.Bounds().Dx())
	sh := float64(src.Bounds().Dy())
	dw := float64(img.Bounds().Dx())
	dh := float64(img.Bounds().Dy())

	scale := math.Max(dw/sw, dh/sh)
	w := sw * scale
	h := sh * scale

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate((dw-w)/2, (dh-h)/2)

	img.DrawImage(src, op)
}

func drawCircle(img *ebiten.Image, cx, cy, r float64, c color.Color) {
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if x*x+y*y <= r*r {
				px := int(cx + x)
				py := int(cy + y)

				if px >= 0 && py >= 0 &&
					px < img.Bounds().Dx() &&
					py < img.Bounds().Dy() {
					img.Set(px, py, c)
				}
			}
		}
	}
}

func clamp(v, lo, hi float64) float64 {
	return max(lo, min(hi, v))
}
