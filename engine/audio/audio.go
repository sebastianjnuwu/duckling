package audio

import (
	"bytes"
	"io"
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
// Lê tudo pra memória (sem handle aberto) e sem asset:
// (nil, nil) = no-op silencioso.
func decode(name string) (*audio.Player, error) {
	path, ok := assets.Tracks[strings.ToLower(name)]
	if !ok {
		return nil, nil
	}
	f, err := assets.FS.Open(strings.TrimPrefix(path, "content/"))
	if err != nil {
		return nil, err
	}
	raw, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		return nil, err
	}
	stream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	return ctx.NewPlayer(stream)
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
