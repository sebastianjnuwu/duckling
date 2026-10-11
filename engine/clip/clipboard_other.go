//go:build !windows

package clip

import "errors"

// clipboard_other.go — fora do Windows, copiar não faz nada.

func Copy(string) error { return errors.New("clip: só Windows") }
