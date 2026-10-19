package world

import (
	"image"

	"github.com/lafriks/go-tiled"
)

// map_task.go — contrato de tarefa (grupo "tasks"): cada objeto é
// um ponto de trabalho. Props: "long" (bool) = tarefa longa;
// "time" (float, segundos); "kind" = "hold" (segurar), "order"
// (segurar na sequência, soltar zera) ou "carry" (levar até
// "to_x"/"to_y"). Sem "kind" válido, é "hold".

// Tipos de tarefa (prop "kind").
const (
	KindHold  = "hold"
	KindCarry = "carry"
	KindOrder = "order"
)

// Task é um ponto de tarefa do mapa.
type Task struct {
	Name string
	Kind string
	Long bool
	Time float64
	Rect image.Rectangle
	Dest image.Point // só "carry": destino (zero = o próprio ponto)
}

func (m *Map) takeTask(o *tiled.Object) {
	r, ok := objRect(o)
	if !ok {
		x, y := int(o.X), int(o.Y)
		r = image.Rect(x, y, x, y)
	}
	m.Tasks = append(m.Tasks, Task{
		Name: o.Name,
		Kind: taskKind(o),
		Long: o.Properties.GetBool("long"),
		Time: taskTime(o),
		Rect: r,
		Dest: taskDest(o),
	})
}

// Center devolve o centro do ponto (onde o pato interage).
func (t Task) Center() image.Point {
	return t.Rect.Min.Add(t.Rect.Max).Div(2)
}

// Target devolve onde o "carry" termina (o destino ou o próprio
// ponto).
func (t Task) Target() image.Point {
	if t.Dest != (image.Point{}) {
		return t.Dest
	}
	return t.Center()
}

// taskKind lê "kind"; desconhecido ou vazio cai em "hold".
func taskKind(o *tiled.Object) string {
	switch k := o.Properties.GetString("kind"); k {
	case KindCarry, KindOrder, KindHold:
		return k
	}
	return KindHold
}

// taskTime lê "time" (float ou int) em segundos.
func taskTime(o *tiled.Object) float64 {
	if v := o.Properties.GetFloat("time"); v > 0 {
		return v
	}
	if v := o.Properties.GetInt("time"); v > 0 {
		return float64(v)
	}
	return 0
}

// taskDest lê o destino do "carry" ("to_x"/"to_y"); sem ele, o
// destino é o próprio ponto (leva daqui pra cá).
func taskDest(o *tiled.Object) image.Point {
	x := propNum(o, "to_x")
	y := propNum(o, "to_y")
	if x == 0 && y == 0 {
		return image.Point{}
	}
	return image.Point{X: int(x), Y: int(y)}
}

// propNum lê número (float ou int) da propriedade.
func propNum(o *tiled.Object, name string) float64 {
	if v := o.Properties.GetFloat(name); v != 0 {
		return v
	}
	return float64(o.Properties.GetInt(name))
}
