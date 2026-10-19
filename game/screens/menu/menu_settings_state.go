package menu

import (
	"duckling/engine/audio"
	"duckling/engine/rendering"
)

// menu_settings_state.go — estado das configurações:
// página 0 = idioma, música e efeitos;
// página 1 = volumes.

type Settings struct {
	langA, langB, langC ui.Btn
	mus                 ui.Toggle
	mvol                ui.Slider
	sfx                 ui.Toggle
	svol                ui.Slider
	x                   ui.Btn
	prev, next          ui.Btn
	reset               ui.Btn
	confirm             bool
	cYes, cNo, cX       ui.Btn
	page                int
	keys                [8]ui.Btn
	cap                 int
}

func newSettings() Settings {
	s := Settings{
		langA: ui.Btn{
			X: 640, Y: 230, W: 80, H: 56,
		},
		langB: ui.Btn{
			X: 730, Y: 230, W: 80, H: 56,
		},
		langC: ui.Btn{
			X: 820, Y: 230, W: 80, H: 56,
		},
		mus: ui.Toggle{
			X: 780, Y: 334, W: 120, H: 48,
			On: audio.MusicOn(),
		},
		sfx: ui.Toggle{
			X: 780, Y: 434, W: 120, H: 48,
			On: audio.SfxOn(),
		},
		mvol: ui.Slider{
			X: 640, Y: 238, W: 260, H: 40,
			Pct: int(audio.MusicVolume() * 100),
		},
		svol: ui.Slider{
			X: 640, Y: 238, W: 260, H: 40,
			Pct: int(audio.SfxVolume() * 100),
		},
		x: ui.Btn{
			X: 884, Y: 140, W: 44, H: 44,
		},
		prev: ui.Btn{
			X: 380, Y: 525, W: 48, H: 40,
		},
		next: ui.Btn{
			X: 852, Y: 525, W: 48, H: 40,
		},
		reset: ui.Btn{
			X: 510, Y: 518, W: 260, H: 40, Icon: ui.IconRefresh,
		},
		page: 0,
		cap:  -1,
	}
	// Confirma do reset: mesmo modal do sair (360,230,560x260).
	s.cYes = ui.Btn{X: 380, Y: 390, W: 240, H: 64}
	s.cNo = ui.Btn{X: 660, Y: 390, W: 240, H: 64}
	s.cX = ui.Btn{X: 840, Y: 250, W: 44, H: 44}
	ys := []float64{226, 288, 350, 412, 279, 337, 395, 453}
	for i := range s.keys {
		s.keys[i] = ui.Btn{X: 700, Y: ys[i], W: 200, H: 48}
	}
	return s
}
