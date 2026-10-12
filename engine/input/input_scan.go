package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// input_scan.go — varre teclas e botões do mouse, devolve o 1º.
// Botão de mouse vira código 1000+n (mesmo tipo, salva igual).

// mouseBit marca botão de mouse no código da tecla.
const mouseBit = 1000

func isMouse(k ebiten.Key) bool { return int(k) >= mouseBit }

func mouseBtn(k ebiten.Key) ebiten.MouseButton { return ebiten.MouseButton(int(k) - mouseBit) }

func mouseKey(b ebiten.MouseButton) ebiten.Key { return ebiten.Key(mouseBit + int(b)) }

func Scan() (ebiten.Key, bool) {
	for k := ebiten.KeyA; k <= ebiten.KeyZ; k++ {
		if inpututil.IsKeyJustPressed(k) {
			return k, true
		}
	}
	for k := ebiten.Key0; k <= ebiten.Key9; k++ {
		if inpututil.IsKeyJustPressed(k) {
			return k, true
		}
	}
	for _, k := range []ebiten.Key{
		ebiten.KeyArrowUp, ebiten.KeyArrowDown,
		ebiten.KeyArrowLeft, ebiten.KeyArrowRight,
		ebiten.KeySpace, ebiten.KeyEnter, ebiten.KeyTab,
		ebiten.KeyBackspace, ebiten.KeyShiftLeft, ebiten.KeyShiftRight,
		ebiten.KeyControlLeft, ebiten.KeyControlRight,
		ebiten.KeyAltLeft, ebiten.KeyAltRight,
	} {
		if inpututil.IsKeyJustPressed(k) {
			return k, true
		}
	}
	for k := ebiten.KeyF1; k <= ebiten.KeyF12; k++ {
		if inpututil.IsKeyJustPressed(k) {
			return k, true
		}
	}
	for _, b := range []ebiten.MouseButton{
		ebiten.MouseButtonLeft, ebiten.MouseButtonMiddle, ebiten.MouseButtonRight,
	} {
		if inpututil.IsMouseButtonJustPressed(b) {
			return mouseKey(b), true
		}
	}
	for _, c := range []int{3, 4} {
		if xPressed(c) {
			return ebiten.Key(mouseBit + c), true
		}
	}
	return 0, false
}
