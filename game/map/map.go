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
// antigo sem nada disso continua igual). Nome de mapa resolve
// por convenção (maps/<nome>/<nome>.tmx); camada vale por NOME
// ou CLASSE, o resto é ignorado sem erro:
//   grupo "collision" -> retângulos sólidos (prop "h": altura,
//     sem h = 1000; h≤25 pula, h≤34 voa, h>34 muro)
//   grupo "COLLISION_WALL" ou classe "BARRIER" -> barreira
//     (sem h = 1000; com h, respeita o valor)
//   grupo "areas"     -> zonas ≥128px (teclas 1-5)
//   objeto "spawn"    -> nascimento; vários = um ponto por pato
//   grupo "tasks"     -> pontos de tarefa (prop "long" = longa)
//   grupo "gates"     -> porteiras que a sabotagem fecha (retângulos)
//   mapa prop "zoom"  -> zoom de gameplay (sem ela, Zoom)
//   mapa prop "maxplayers" -> teto de jogadores (sem ela, 16)

type Map struct {
	W, H   int
	Areas  []Area
	Solids []Solid
	Tasks  []Task
	Gates  []image.Rectangle
	zoom   float64
	maxp   int
	img    *ebiten.Image
	view   image.Rectangle
	sx, sy float64
	spawns []image.Point
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
		zoom: Zoom, maxp: 16,
		img: ebiten.NewImageFromImage(r.Result),
	}
	// Props do mapa (Map -> Properties): zoom e teto de jogadores.
	// Sem elas, valem os padrões.
	if doc.Properties != nil {
		if v := doc.Properties.GetFloat("zoom"); v > 0 {
			m.zoom = v
		}
		if v := doc.Properties.GetInt("maxplayers"); v > 0 {
			m.maxp = v
		}
	}
	// Centro do que APARECE (pixels com alpha), não do mapa vazio:
	// o spawn padrão e a câmera nascem no meio da ilha desenhada.
	m.view = visualBounds(r.Result)
	m.sx = float64(m.view.Min.X+m.view.Max.X) / 2
	m.sy = float64(m.view.Min.Y+m.view.Max.Y) / 2
	m.scan(doc)
	return m, nil
}
