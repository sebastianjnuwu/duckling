package world

import (
	"image"

	"duckling/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

// map.go — mapa TMX do Tiled via go-tiled, pré-renderizado uma vez.
// Spawn vem do objeto "spawn" (ou do centro visível, se não houver).

type Map struct {
	W, H   int
	Areas  []Area
	Solids []image.Rectangle
	img    *ebiten.Image
	view   image.Rectangle
	sx, sy float64
}

// Load lê ex: maps/spawn/spawn.tmx de dentro de assets.FS.
func Load(p string) (*Map, error) {
	doc, err := tiled.LoadFile(p, tiled.WithFileSystem(assets.FS))
	if err != nil {
		return nil, err
	}
	r, err := render.NewRendererWithFileSystem(doc, assets.FS)
	if err != nil {
		return nil, err
	}
	if err := r.RenderVisibleLayers(); err != nil {
		return nil, err
	}
	m := &Map{
		W: doc.Width * doc.TileWidth, H: doc.Height * doc.TileHeight,
		img: ebiten.NewImageFromImage(r.Result),
	}
	// Centro do que APARECE (pixels com alpha), não do mapa vazio:
	// o spawn padrão e a câmera nascem no meio da ilha desenhada.
	m.view = visualBounds(r.Result)
	m.sx = float64(m.view.Min.X+m.view.Max.X) / 2
	m.sy = float64(m.view.Min.Y+m.view.Max.Y) / 2
	for _, g := range doc.ObjectGroups {
		for _, o := range g.Objects {
			if o.Name == "spawn" || o.Class == "spawn" {
				m.sx, m.sy = o.X, o.Y
			}
			if ar, ok := objRect(o); ok && ar.Dx() >= 128 && ar.Dy() >= 128 {
				m.Areas = append(m.Areas, Area{Name: o.Name, Rect: ar})
			}
			// Grupo "collision": retângulos sólidos (cercas, paredes, água).
			if g.Name == "collision" {
				if sr, ok := objRect(o); ok {
					m.Solids = append(m.Solids, sr)
				}
			}
		}
	}
	return m, nil
}
