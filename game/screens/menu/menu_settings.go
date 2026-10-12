package menu

import (
	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/engine/save"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_settings.go — atualiza os controles de áudio e idioma,
// gravando as configurações a cada mudança.

func (m *Menu) settingsUpdate() {
	s := &m.set

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.state = stMain
		return
	}

	if s.x.Push() {
		m.click()
		m.state = stMain
		return
	}

	mp := maxPage()
	s.page = min(mp, max(0, s.page))

	if s.page > 0 {
		if s.prev.Update() {
			m.click()
			s.page--
		}
	}
	if s.page < mp {
		if s.next.Update() {
			m.click()
			s.page++
		}
	}

	if s.page == 0 {
		s.langA.Y, s.langB.Y, s.langC.Y = 230, 230, 230

		m.pick(&s.langA, "pt")
		m.pick(&s.langB, "en")
		m.pick(&s.langC, "es")

		s.mus.Y = 334
		if s.mus.Update() {
			m.click()
			audio.SetMusicOn(s.mus.On)
			save.Store()
		}

		s.sfx.Y = 434
		if s.sfx.Update() {
			m.click()
			audio.SetSfxOn(s.sfx.On)
			save.Store()
		}

		return
	}

	if s.page != 1 || (!audio.MusicOn() && !audio.SfxOn()) {
		m.controlsUpdate()
		return
	}

	y := 230.0

	if audio.MusicOn() {
		s.mvol.Y = y + 8
		if s.mvol.Update() {
			audio.SetMusicVolume(float64(s.mvol.Pct) / 100)
			save.Store()
		}
		y += 100
	}

	if audio.SfxOn() {
		s.svol.Y = y + 8
		if s.svol.Update() {
			audio.SetSfxVolume(float64(s.svol.Pct) / 100)
			save.Store()
		}
	}
}

// pick seleciona o idioma no clique.
func (m *Menu) pick(b *ui.Btn, code string) {
	if b.Update() {
		m.click()
		i18n.Set(code)
		save.Store()
	}
}
