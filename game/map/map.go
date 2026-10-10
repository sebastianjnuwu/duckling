package world

import (
	"image"

	"duckling/content"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

// map.go — mapa TMX do Tiled via go-tiled, pré-renderizado uma vez.
//
// Contrato (propriedades custom do Tiled, tudo opcional — mapa
// antigo sem nada disso continua igual):
//   grupo "collision" -> retângulos sólidos (prop "h": altura,
//     sem h = 1000, nem pula nem voa por cima)
//   grupo "areas"     -> zonas ≥128px (teclas 1-5)
//   objeto "spawn"    -> nascimento (sem ele, centro visível)
//   mapa prop "zoom"  -> zoom de gameplay (sem ela, Zoom)

type Map struct {
	W, H   int
	Areas  []Area
	Solids []Solid
	zoom   float64
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
		zoom: Zoom,
		img:  ebiten.NewImageFromImage(r.Result),
	}
	// Prop "zoom" do mapa (Map -> Properties): zoom de gameplay
	// desse mapa. Sem ela, vale o Zoom padrão.
	if doc.Properties != nil {
		if v := doc.Properties.GetFloat("zoom"); v > 0 {
			m.zoom = v
		}
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
					h := o.Properties.GetFloat("h")
					if h <= 0 {
						h = 1000
					}
					m.Solids = append(m.Solids, Solid{Rect: sr, H: h})
				}
			}
		}
	}
	return m, nil
}
