package menu

import (
	"duckling/src/audio"
	"duckling/src/i18n"
	"duckling/src/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_settings.go — modal config em páginas: 0 = chaves
// (idioma, música, sfx), 1 = volumes. Setas trocam.

type Settings struct {
	langA, langB, langC ui.Btn
	mus                 ui.Toggle
	mvol                ui.Slider
	sfx                 ui.Toggle
	svol                ui.Slider
	x                   ui.Btn
	prev, next          ui.Btn
	page                int
}

func newSettings() Settings {
	return Settings{
		langA: ui.Btn{X: 640, Y: 230, W: 80, H: 56}, langB: ui.Btn{X: 730, Y: 230, W: 80, H: 56}, langC: ui.Btn{X: 820, Y: 230, W: 80, H: 56},
		mus: ui.Toggle{X: 780, Y: 334, W: 120, H: 48, On: audio.MusicOn()}, sfx: ui.Toggle{X: 780, Y: 434, W: 120, H: 48, On: audio.SfxOn()},
		mvol: ui.Slider{X: 640, Y: 238, W: 260, H: 40, Pct: int(audio.MusicVolume() * 100)}, svol: ui.Slider{X: 640, Y: 238, W: 260, H: 40, Pct: int(audio.SfxVolume() * 100)},
		x: ui.Btn{X: 900, Y: 132, W: 48, H: 48}, prev: ui.Btn{X: 380, Y: 525, W: 48, H: 40}, next: ui.Btn{X: 852, Y: 525, W: 48, H: 40},
	}
}

func (m *Menu) settingsUpdate() {
	s := &m.set
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.state = stMain
		return
	}
	if s.x.Update() {
		m.click()
		m.state = stMain
	}
	mp := maxPage()
	s.page = min(mp, max(0, s.page))
	if mp > 0 {
		if s.prev.Update() {
			m.click()
			s.page = max(0, s.page-1)
		}
		if s.next.Update() {
			m.click()
			s.page = min(mp, s.page+1)
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
		}
		s.sfx.Y = 434
		if s.sfx.Update() {
			m.click()
			audio.SetSfxOn(s.sfx.On)
		}
		return
	}
	y := 230.0
	if audio.MusicOn() {
		s.mvol.Y = y + 8
		if s.mvol.Update() {
			audio.SetMusicVolume(float64(s.mvol.Pct) / 100)
		}
		y += 100
	}
	if audio.SfxOn() {
		s.svol.Y = y + 8
		if s.svol.Update() {
			audio.SetSfxVolume(float64(s.svol.Pct) / 100)
		}
	}
}

// pick idioma no clique.
func (m *Menu) pick(b *ui.Btn, code string) {
	if b.Update() {
		m.click()
		i18n.Set(code)
	}
}
