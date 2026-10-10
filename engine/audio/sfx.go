package audio

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// sfx.go — efeito tiro único (ex: click). Toca uma vez e morre;
// nunca é religado pelo Update. Sem asset ou off: no-op.
// Cliques com <90ms de intervalo caem (anti-spam).

var (
	sfx     *audio.Player
	lastSfx time.Time
)

func Play(name string) error {
	if !sfxOn {
		return nil
	}
	if time.Since(lastSfx) < 100*time.Millisecond {
		return nil
	}
	lastSfx = time.Now()
	if sfx != nil {
		sfx.Close()
		sfx = nil
	}
	p, err := decode(name)
	if err != nil || p == nil {
		return err
	}
	sfx = p
	p.SetVolume(sfxVol)
	p.Play()
	return nil
}

func applySfxVolume() {
	if sfx != nil {
		sfx.SetVolume(sfxVol)
	}
}

func applySfxEnabled() {
	if sfx == nil {
		return
	}
	if sfxOn {
		sfx.Play()
	} else {
		sfx.Pause()
	}
}
