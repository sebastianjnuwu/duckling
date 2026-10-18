package actor

import (
	"duckling/engine/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// net.go — entrada de um pato (local ou vinda da rede) e o
// estado mínimo que trafega: posição, altitude, direção e voo.

const (
	BitUp uint8 = 1 << iota
	BitDown
	BitLeft
	BitRight
	BitJump
	BitFly
)

// In é a entrada de um frame. Edge = apertou agora (pulo/voo).
type In struct {
	Up, Down, Left, Right bool
	Jump, Fly             bool
	JumpEdge, FlyEdge     bool
}

// LocalIn lê o teclado: WASD/setas, pulo e voo (E é alternativo).
func LocalIn() In {
	return In{
		Up: input.Held(input.Up), Down: input.Held(input.Down),
		Left: input.Held(input.Left), Right: input.Held(input.Right),
		Jump:     input.Held(input.Jump),
		Fly:      input.Held(input.Fly) || ebiten.IsKeyPressed(ebiten.KeyE),
		JumpEdge: input.JustPressed(input.Jump),
		FlyEdge:  input.JustPressed(input.Fly) || inpututil.IsKeyJustPressed(ebiten.KeyE),
	}
}

// Encode vira bits segurados; a borda (Edge) o host deriva.
func (in In) Encode() uint8 {
	var b uint8
	if in.Up {
		b |= BitUp
	}
	if in.Down {
		b |= BitDown
	}
	if in.Left {
		b |= BitLeft
	}
	if in.Right {
		b |= BitRight
	}
	if in.Jump {
		b |= BitJump
	}
	if in.Fly {
		b |= BitFly
	}
	return b
}

// DecodeIn monta a entrada dos bits segurados + bits de borda.
func DecodeIn(cur, edge uint8) In {
	return In{
		Up: cur&BitUp != 0, Down: cur&BitDown != 0,
		Left: cur&BitLeft != 0, Right: cur&BitRight != 0,
		Jump: cur&BitJump != 0, Fly: cur&BitFly != 0,
		JumpEdge: edge&BitJump != 0, FlyEdge: edge&BitFly != 0,
	}
}
