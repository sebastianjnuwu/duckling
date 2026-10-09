package actor

// anim.go — direção e quadros: o 1º toque troca a pose na hora,
// parado volta ao frame base.

func (d *Duck) Anim(dx, dy float64) {
	if d.side == nil {
		return
	}
	if dx == 0 && dy == 0 && !d.flying {
		d.side.Reset()
		d.down.Reset()
		d.up.Reset()
		d.moving = false
		return
	}
	if d.flying {
		d.dir = 0
		d.Tick()
		return
	}
	switch {
	case dy < 0:
		d.dir = 2
	case dy > 0:
		d.dir = 1
	case dx < 0:
		d.dir, d.face = 0, -1
	case dx > 0:
		d.dir, d.face = 0, 1
	}
	if !d.moving {
		d.moving = true
		d.cur().Step()
	}
	d.Tick()
}

func (d *Duck) Tick() {
	d.cur().Tick()
}
