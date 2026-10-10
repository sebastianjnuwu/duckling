package menu

import (
	"duckling/content"
	"duckling/engine/audio"
	"duckling/engine/rendering"
)

// menu_media.go — música e pato preguiçosos: entram quando dá.

func (m *Menu) ensureMedia() {
	// Música entra no menu (lazy: garante loop de áudio rodando).
	if !m.bgm {
		m.bgm = true
		_ = audio.PlayLoop("loop")
	}
	// Pato carrega lazy: imagens só existem após o loading.
	if m.duck == nil {
		if img := assets.Images["sprites/duck/yellow/duck-4x"]; img != nil {
			m.duck = ui.NewDuck(img, 1, 128, 24)
		}
	} else {
		m.duck.Tick()
	}
}
