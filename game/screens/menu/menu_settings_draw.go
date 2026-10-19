package menu

import (
	"image/color"

	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// menu_settings_draw.go — desenha os controles,
// as setas disponíveis e os indicadores de página.

func (m *Menu) settingsDraw(img *ebiten.Image) {
	s := &m.set
	cream := color.RGBA{235, 240, 235, 255}

	ui.Dim(img)
	ui.Panel(img, 320, 120, 640, 480)

	s.x.Draw(img, "X", m.font)
	if s.page == 0 || (s.page == 1 && (audio.MusicOn() || audio.SfxOn())) {
		ui.Title(img, ui.IconGear, i18n.T("settings.title"), m.sml, 150, ui.Gold)
	}

	if s.page == 0 {
		left(img, i18n.T("settings.lang"), m.font, 380, 249)
		ui.DrawIcon(img, ui.IconGlobe, 348, 246, cream)

		ui.FlagButton(img, s.langA, "br", i18n.Lang == "pt")
		ui.FlagButton(img, s.langB, "us", i18n.Lang == "en")
		ui.FlagButton(img, s.langC, "es", i18n.Lang == "es")

		ui.DrawIcon(img, ui.IconNote, 348, 346, cream)
		left(img, i18n.T("settings.music"), m.font, 380, 349)
		s.mus.Draw(img)

		ui.DrawIcon(img, ui.IconSpk, 348, 446, cream)
		left(img, i18n.T("settings.sfx"), m.font, 380, 449)
		s.sfx.Draw(img)
	} else if s.page == 1 && (audio.MusicOn() || audio.SfxOn()) {
		y := 230.0

		if audio.MusicOn() {
			ui.DrawIcon(img, ui.IconSpk, 348, y+16, cream)
			left(img, i18n.T("settings.musicVol"), m.font, 380, y+19)
			s.mvol.Draw(img)
			y += 100
		}

		if audio.SfxOn() {
			ui.DrawIcon(img, ui.IconSpk, 348, y+16, cream)
			left(img, i18n.T("settings.sfxVol"), m.font, 380, y+19)
			s.svol.Draw(img)
		}
	} else {
		ui.Title(img, ui.IconGear, i18n.T("settings.controls"), m.sml, 150, ui.Gold)
		s.reset.Draw(img, i18n.T("settings.reset"), m.font)
		if s.page == movePage() {
			for i := 0; i < 4; i++ {
				y := 226.0 + float64(i)*62
				ui.DrawIcon(img, keyRows[i].ic, 348, y+11, cream)
				left(img, i18n.T(keyRows[i].key), m.font, 380, y+14)
				s.keys[i].Draw(img, m.keyName(i), m.font)
			}
		} else {
			for i := 4; i < len(keyRows); i++ {
				y := 279.0 + float64(i-4)*58
				ui.DrawIcon(img, keyRows[i].ic, 348, y+11, cream)
				left(img, i18n.T(keyRows[i].key), m.font, 380, y+14)
				s.keys[i].Draw(img, m.keyName(i), m.font)
			}
		}
	}

	mp := maxPage()

	if mp > 0 {
		if s.page > 0 {
			s.prev.Draw(img, "<", m.font)
		}

		if s.page < mp {
			s.next.Draw(img, ">", m.font)
		}

		for i := 0; i <= mp; i++ {
			c := color.RGBA{70, 90, 130, 255}
			if i == s.page {
				c = color.RGBA{232, 190, 72, 255}
			}

			ebitenutil.DrawRect(
				img,
				620+float64(i)*28,
				566,
				12,
				12,
				c,
			)
		}
	}

	if s.confirm {
		m.resetConfirmDraw(img)
	}
}

// maxPage retorna a última página (controles sempre por último).
func maxPage() int {
	if audio.MusicOn() || audio.SfxOn() {
		return 3
	}
	return 2
}

// movePage: página do movimento (1 sem áudio, 2 com).
func movePage() int {
	if audio.MusicOn() || audio.SfxOn() {
		return 2
	}
	return 1
}

// left desenha texto alinhado à esquerda em (x, y).
func left(
	img *ebiten.Image,
	s string,
	f *text.GoTextFace,
	x, y float64,
) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(img, s, f, op)
}
