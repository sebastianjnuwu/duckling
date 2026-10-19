package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// input.go — ações remapeáveis: 1 tecla por ação + setas
// fixas nas direções. Salva em settings.json (códigos).

type Action int

const (
	Up Action = iota
	Down
	Left
	Right
	Jump
	Fly
	Interact
)

var keys = map[Action]ebiten.Key{
	Up: ebiten.KeyW, Down: ebiten.KeyS, Left: ebiten.KeyA,
	Right: ebiten.KeyD, Jump: ebiten.KeySpace, Fly: ebiten.KeyX,
	Interact: ebiten.KeyE,
}

var arrows = map[Action]ebiten.Key{
	Up: ebiten.KeyArrowUp, Down: ebiten.KeyArrowDown,
	Left: ebiten.KeyArrowLeft, Right: ebiten.KeyArrowRight,
}

// Held: tecla/botão da ação ou seta (direções).
func Held(a Action) bool {
	if k := keys[a]; isMouse(k) {
		c := int(k) - mouseBit
		if c < 3 {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButton(c)) {
				return true
			}
		} else if xDown(c) {
			return true
		}
	} else if ebiten.IsKeyPressed(k) {
		return true
	}
	k, ok := arrows[a]
	return ok && ebiten.IsKeyPressed(k)
}

// JustPressed: borda da ação (pulo/voo, sem seta).
func JustPressed(a Action) bool {
	if k := keys[a]; isMouse(k) {
		c := int(k) - mouseBit
		if c < 3 {
			return inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(c))
		}
		return xPressed(c)
	}
	return inpututil.IsKeyJustPressed(keys[a])
}

// Key devolve a tecla (pra UI).
func Key(a Action) ebiten.Key { return keys[a] }

// Set troca a tecla (vale na hora; salva quem chama Store).
func Set(a Action, k ebiten.Key) { keys[a] = k }

// Swap bota k em a; quem tinha k fica com a velha de a.
func Swap(a Action, k ebiten.Key) {
	old := keys[a]
	for o, ko := range keys {
		if o != a && ko == k {
			keys[o] = old
		}
	}
	keys[a] = k
}

// Name mostra a tecla (setas e mouse com símbolo).
func Name(k ebiten.Key) string {
	if isMouse(k) {
		return "MOUSE " + string(rune('1'+int(k)-mouseBit))
	}
	switch k {
	case ebiten.KeyArrowUp:
		return "↑"
	case ebiten.KeyArrowDown:
		return "↓"
	case ebiten.KeyArrowLeft:
		return "←"
	case ebiten.KeyArrowRight:
		return "→"
	}
	return k.String()
}
