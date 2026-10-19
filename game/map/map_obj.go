package world

import (
	"github.com/lafriks/go-tiled"
)

// map_obj.go — varre os objeto-grupos do Tiled e distribui por
// contrato (spawn, tasks, areas, collision); desconhecido é
// ignorado sem erro. Offset da camada soma na origem do objeto.

func (m *Map) scan(doc *tiled.Map) {
	for _, g := range doc.ObjectGroups {
		for _, o := range g.Objects {
			oc := *o
			oc.X += g.OffsetX
			oc.Y += g.OffsetY
			o := &oc

			if o.Name == "spawn" || o.Class == "spawn" {
				m.takeSpawn(o)
				continue
			}
			if groupIs(g, "tasks") {
				m.takeTask(o)
				continue
			}
			if ar, ok := objRect(o); ok && ar.Dx() >= 128 && ar.Dy() >= 128 {
				m.Areas = append(m.Areas, Area{Name: o.Name, Rect: ar})
			}
			// Grupo "collision": retângulos sólidos (cercas, paredes, água).
			if groupIs(g, "collision") {
				if sr, ok := objRect(o); ok {
					m.Solids = append(m.Solids, Solid{Rect: sr, H: solidH(o.Properties)})
				}
			}
			// "COLLISION_WALL" ou classe "BARRIER": barreira (sem h = 1000).
			if groupIs(g, "COLLISION_WALL") || o.Class == "BARRIER" {
				if sr, ok := objRect(o); ok {
					m.Solids = append(m.Solids, Solid{Rect: sr, H: solidH(o.Properties)})
				}
			}
		}
	}
}
