package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// input_save.go — códigos pra settings.json (tecla ou mouse).

var actionNames = map[Action]string{
	Up: "up", Down: "down", Left: "left",
	Right: "right", Jump: "jump", Fly: "fly",
	Interact: "interact",
}

// Codes exporta pra save; LoadCodes importa (ignora lixo).
func Codes() map[string]int {
	out := map[string]int{}
	for a, n := range actionNames {
		out[n] = int(keys[a])
	}
	return out
}

func LoadCodes(m map[string]int) {
	back := map[string]Action{}
	for a, n := range actionNames {
		back[n] = a
	}
	for n, c := range m {
		if a, ok := back[n]; ok && c > 0 && c < 2000 {
			keys[a] = ebiten.Key(c)
		}
	}
}

// Reset volta aos padrões (WASD, espaço, X, E).
func Reset() {
	Set(Up, ebiten.KeyW)
	Set(Down, ebiten.KeyS)
	Set(Left, ebiten.KeyA)
	Set(Right, ebiten.KeyD)
	Set(Jump, ebiten.KeySpace)
	Set(Fly, ebiten.KeyX)
	Set(Interact, ebiten.KeyE)
}
