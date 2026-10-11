//go:build windows

package clip

import (
	"errors"
	"syscall"
	"unsafe"
)

// clipboard_windows.go — copiar texto (só Windows, sem cgo).

var (
	user32            = syscall.NewLazyDLL("user32.dll")
	kernel32          = syscall.NewLazyDLL("kernel32.dll")
	pOpenClipboard    = user32.NewProc("OpenClipboard")
	pEmptyClipboard   = user32.NewProc("EmptyClipboard")
	pSetClipboardData = user32.NewProc("SetClipboardData")
	pCloseClipboard   = user32.NewProc("CloseClipboard")
	pGlobalAlloc      = kernel32.NewProc("GlobalAlloc")
	pGlobalLock      = kernel32.NewProc("GlobalLock")
	pGlobalUnlock    = kernel32.NewProc("GlobalUnlock")
	pRtlMoveMemory   = kernel32.NewProc("RtlMoveMemory")
)

const (
	cfUnicodeText = 13
	gmemMoveable  = 0x0002
)

// Copy joga o texto na área de transferência.
func Copy(text string) error {
	u16, err := syscall.UTF16FromString(text)
	if err != nil {
		return err
	}
	if r, _, _ := pOpenClipboard.Call(0); r == 0 {
		return errors.New("clip: abrir")
	}
	defer pCloseClipboard.Call()
	pEmptyClipboard.Call()
	h, _, _ := pGlobalAlloc.Call(gmemMoveable, uintptr(len(u16))*2)
	if h == 0 {
		return errors.New("clip: memória")
	}
	p, _, _ := pGlobalLock.Call(h)
	if p == 0 {
		return errors.New("clip: travar")
	}
	pRtlMoveMemory.Call(p, uintptr(unsafe.Pointer(&u16[0])), uintptr(len(u16))*2)
	pGlobalUnlock.Call(h)
	if r, _, _ := pSetClipboardData.Call(cfUnicodeText, h); r == 0 {
		return errors.New("clip: colar")
	}
	return nil
}
