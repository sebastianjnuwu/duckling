package save

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/engine/input"
)

// save.go — config local (idioma, música, sfx) em JSON.
// Carrega no boot, grava a cada mudança; no futuro vira servidor.

const Version = 1

type Data struct {
	V        int            `json:"v"`
	Lang     string         `json:"lang"`
	MusicOn  bool           `json:"musicOn"`
	MusicVol float64        `json:"musicVol"`
	SfxOn    bool           `json:"sfxOn"`
	SfxVol   float64        `json:"sfxVol"`
	Keys     map[string]int `json:"keys"`
	Coins    int            `json:"coins"`
	Hats     []int          `json:"hats,omitempty"`
	Hat      int            `json:"hat,omitempty"`
}

var last []byte

func path() (string, bool) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", false
	}
	return filepath.Join(dir, "duckling", "settings.json"), true
}

// Load lê o disco; sem arquivo (ou corrompido), padrão atual.
// Corrompido vira settings.corrupt.json pra não perder a prova.
func Load() Data {
	d := Current()
	d.V = Version
	p, ok := path()
	if !ok {
		return d
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return d
	}
	if json.Unmarshal(b, &d) != nil {
		_ = os.Rename(p, p+".corrupt")
	}
	return d
}

// Apply joga o salvo no idioma, áudio e teclas.
func Apply(d Data) {
	i18n.Set(d.Lang)
	audio.SetMusicOn(d.MusicOn)
	audio.SetMusicVolume(d.MusicVol)
	audio.SetSfxOn(d.SfxOn)
	audio.SetSfxVolume(d.SfxVol)
	input.LoadCodes(d.Keys)
	coins = d.Coins
	loadHats(d.Hats)
	hat = d.Hat
	if !HatOwned(hat) {
		hat = 0
	}
	last, _ = json.MarshalIndent(d, "", "  ")
}

// Current fotografa o estado atual.
func Current() Data {
	return Data{Version, i18n.Lang, audio.MusicOn(), audio.MusicVolume(), audio.SfxOn(), audio.SfxVolume(), input.Codes(), coins, hatList(), hat}
}

// Store grava o estado atual (atômico, só se mudou).
// Sem disco, no-op silencioso.
func Store() {
	p, ok := path()
	if !ok {
		return
	}
	b, _ := json.MarshalIndent(Current(), "", "  ")
	if bytes.Equal(b, last) {
		return
	}
	last = b
	_ = os.MkdirAll(filepath.Dir(p), 0o755)
	t := p + ".tmp"
	if os.WriteFile(t, b, 0o600) == nil {
		_ = os.Rename(t, p)
	}
}
