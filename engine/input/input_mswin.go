//go:build windows

package input

import "syscall"

// input_mswin.go — laterais do mouse (X1/X2) via Win32.
// Códigos 3 e 4 (MOUSE 4/5). Sem cgo, stdlib pura.

var pGetAsyncKeyState = syscall.NewLazyDLL("user32.dll").NewProc("GetAsyncKeyState")

func xState(code int) uintptr {
	vk := uintptr(0x05)
	if code == 4 {
		vk = 0x06
	}
	r, _, _ := pGetAsyncKeyState.Call(vk)
	return r
}

// xDown: segurando agora. xPressed: borda desde a leitura.
func xDown(code int) bool    { return xState(code)&0x8000 != 0 }
func xPressed(code int) bool { return xState(code)&0x0001 != 0 }
