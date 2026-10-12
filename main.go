package main

import (
	"image"
	_ "image/png"
	"log"

	"duckling/content"
	"duckling/engine/audio"
	"duckling/engine/rendering"
	"duckling/engine/save"
	"duckling/engine/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Duckling")

	if err := audio.Init(); err != nil {
		log.Fatal(err)
	}
	ui.ProvideIconFont(func() *text.GoTextFaceSource { return assets.Fonts["MaterialSymbolsRounded-Regular"] })
	save.Apply(save.Load())

	var icons []image.Image
	for _, sz := range []string{"256x256", "128x128", "64x64", "48x48", "32x32", "16x16"} {
		if f, err := assets.FS.Open("images/icons/duck/transparent/" + sz + ".png"); err == nil {
			if icon, _, err := image.Decode(f); err == nil {
				icons = append(icons, icon)
			}
			f.Close()
		}
	}
	if len(icons) > 0 {
		ebiten.SetWindowIcon(icons)
	} else if f, err := assets.FS.Open("images/logo.png"); err == nil {
		if icon, _, err := image.Decode(f); err == nil {
			ebiten.SetWindowIcon([]image.Image{icon})
		}
		f.Close()
	}

	game, err := screen.UI()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(ebiten.RunGame(game))
}
