package world

import (
	"image"

	"github.com/lafriks/go-tiled"
)

// map_geo.go — geometria: spawn, área visível, bbox de objetos.

// Area é um retângulo (px) de objeto do Tiled com ≥128px por lado.
type Area struct {
	Name string
	Rect image.Rectangle
}

// Solid é um retângulo bloqueante com altura H (px de mundo):
// o pato só colide se estiver abaixo dela (z < H).
// H vem da propriedade custom `h` no Tiled (falta = 1000, muro).
type Solid struct {
	Rect image.Rectangle
	H    float64
}

func (m *Map) Spawn() (float64, float64) {
	return m.sx, m.sy
}

// View é o retângulo visível do mapa (bbox dos pixels desenhados).
func (m *Map) View() image.Rectangle {
	return m.view
}

// visualBounds varre a imagem atrás de pixels opacos.
func visualBounds(img *image.NRGBA) image.Rectangle {
	b := img.Bounds()
	x0, y0 := b.Max.X, b.Max.Y
	x1, y1 := b.Min.X, b.Min.Y
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if img.NRGBAAt(x, y).A > 0 {
				x0, y0 = min(x0, x), min(y0, y)
				x1, y1 = max(x1, x), max(y1, y)
			}
		}
	}
	if x1 < x0 {
		return b
	}
	return image.Rect(x0, y0, x1+1, y1+1)
}

// objRect devolve o bbox do objeto (retângulo, polígono ou linha).
func objRect(o *tiled.Object) (image.Rectangle, bool) {
	x0, y0 := o.X, o.Y
	x1, y1 := o.X+o.Width, o.Y+o.Height
	for _, pg := range o.Polygons {
		if pg.Points == nil {
			continue
		}
		for _, pt := range *pg.Points {
			x, y := o.X+pt.X, o.Y+pt.Y
			x0, y0 = min(x0, x), min(y0, y)
			x1, y1 = max(x1, x), max(y1, y)
		}
	}
	if x1 > x0 && y1 > y0 {
		return image.Rect(int(x0), int(y0), int(x1), int(y1)), true
	}
	return image.Rectangle{}, false
}
