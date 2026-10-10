package audio

import "github.com/hajimehoshi/ebiten/v2/audio"

// music.go — trilha de fundo em loop. É o único player que o
// Update religa (tickMusic). SFX passa longe daqui.

var (
	music     *audio.Player
	musicLoop = false
)

func PlayLoop(name string) error {
	StopMusic()
	p, err := decode(name)
	if err != nil || p == nil {
		return err
	}
	music, musicLoop = p, true
	p.SetVolume(musicVol)
	if musicOn {
		p.Play()
	}
	return nil
}

func StopMusic() {
	musicLoop = false
	if music != nil {
		music.Close()
		music = nil
	}
}

func MusicPlaying() bool { return music != nil && music.IsPlaying() }

func tickMusic() {
	if musicLoop && music != nil && !music.IsPlaying() && musicOn {
		music.Rewind()
		music.Play()
	}
}

func applyMusicVolume() {
	if music != nil {
		music.SetVolume(musicVol)
	}
}

func applyMusicEnabled() {
	if music == nil {
		return
	}
	if musicOn {
		music.Play()
	} else {
		music.Pause()
	}
}
