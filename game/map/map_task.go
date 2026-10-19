package world

import (
	"image"

	"github.com/lafriks/go-tiled"
)

// map_task.go — contrato de tarefa (grupo "tasks"): cada objeto é
// um ponto de trabalho. Props: "long" (bool) = tarefa longa;
// "time" (float/int, opcional) = segundos. Só os dados: a
// interação (segurar E, barra, moedas) é da Fase 3.

// Task é um ponto de tarefa do mapa.
type Task struct {
	Name string
	Long bool
	Time float64
	Rect image.Rectangle
}

func (m *Map) takeTask(o *tiled.Object) {
	r, ok := objRect(o)
	if !ok {
		x, y := int(o.X), int(o.Y)
		r = image.Rect(x, y, x, y)
	}
	t := Task{Name: o.Name, Long: o.Properties.GetBool("long"), Rect: r}
	if v := o.Properties.GetFloat("time"); v > 0 {
		t.Time = v
	} else if v := o.Properties.GetInt("time"); v > 0 {
		t.Time = float64(v)
	}
	m.Tasks = append(m.Tasks, t)
}
