package actor

// remote.go — o pato de outro jogador: chega por snapshot e é
// interpolado; o local só é corrigido se divergir muito.

// Interp puxa o pato remoto pra posição da rede (suave).
func (d *Duck) Interp(x, y, z float64, dir int, fly bool) {
	dx, dy := x-d.Px, y-d.Py
	if dx*dx+dy*dy > 0.01 {
		d.Px += dx * 0.25
		d.Py += dy * 0.25
	}
	d.z, d.dir, d.flying = z, dir, fly
	if d.down == nil {
		return
	}
	if dx != 0 || dy != 0 || fly {
		d.Tick()
		return
	}
	// Parado: volta ao frame base.
	d.down.Reset()
	d.up.Reset()
	d.left.Reset()
	d.right.Reset()
}

// Snap cola o pato na posição autoritativa (correção dura).
func (d *Duck) Snap(x, y, z float64, dir int, fly bool) {
	d.Px, d.Py, d.z, d.dir, d.flying = x, y, z, dir, fly
}

// Dir e Flying expõem o estado pra montar o snapshot.
func (d *Duck) Dir() int     { return d.dir }
func (d *Duck) Flying() bool { return d.flying }
