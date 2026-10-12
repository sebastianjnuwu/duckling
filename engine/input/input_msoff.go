//go:build !windows

package input

// input_msoff.go — fora do Windows, sem botão lateral.

func xDown(int) bool    { return false }
func xPressed(int) bool { return false }
