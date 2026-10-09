package menu

import (
	"duckling/src/audio"
	"duckling/src/ui"
)

// menu_settings_state.go — estado do modal config em páginas:
// 0 = chaves (idioma, música, sfx), 1 = volumes. Setas trocam.

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
