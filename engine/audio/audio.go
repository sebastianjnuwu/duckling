package audio

import (
	"strings"

	"duckling/content"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

// audio.go — contexto + 2 canais (música e sfx), cada um
// com volume e liga/desliga próprios.

const sampleRate = 44100

var (
	ctx      *audio.Context
	musicVol = 0.5
	musicOn  = true
	sfxVol   = 0.5
	sfxOn    = true
)

func Init() error {
	ctx = audio.NewContext(sampleRate)
	return nil
}

func Update() {
	tickMusic()
}

// decode abre o asset e devolve o player parado.
// Sem asset: (nil, nil) = no-op silencioso.
func decode(name string) (*audio.Player, error) {
	path, ok := assets.Tracks[name]
	if !ok {
		return nil, nil
	}
	f, err := assets.FS.Open(strings.TrimPrefix(path, "content/"))
	if err != nil {
		return nil, err
	}
	stream, err := mp3.DecodeWithSampleRate(sampleRate, f)
	if err != nil {
		f.Close()
		return nil, err
	}
	p, err := ctx.NewPlayer(stream)
	if err != nil {
		f.Close()
		return nil, err
	}
	return p, nil
}

func MusicVolume() float64 { return musicVol }

func SetMusicVolume(v float64) {
	musicVol = min(1, max(0, v))
	applyMusicVolume()
}

func MusicOn() bool { return musicOn }

func SetMusicOn(v bool) {
	musicOn = v
	applyMusicEnabled()
}

func SfxVolume() float64 { return sfxVol }

func SetSfxVolume(v float64) {
	sfxVol = min(1, max(0, v))
	applySfxVolume()
}

func SfxOn() bool { return sfxOn }

func SetSfxOn(v bool) {
	sfxOn = v
	applySfxEnabled()
}
