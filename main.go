package main

import (
	"image"
	_ "image/png"
	"log"

	"duckling/content"
	"duckling/engine/audio"
	"duckling/game/save"
	"duckling/engine/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Duckling")

	if err := audio.Init(); err != nil {
		log.Fatal(err)
	}
	save.Apply(save.Load())

	if f, err := assets.FS.Open("images/logo.png"); err == nil {
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
