package screen

import (
	"image/color"

	"duckling/src/audio"
	"duckling/src/screen/loading"
	"duckling/src/screen/menu"
	"duckling/src/screen/shop"

	"github.com/hajimehoshi/ebiten/v2"
)

type Screen interface {
	Update() error
	Draw(*ebiten.Image)
}

type Game struct {
	m   map[string]Screen
	cur string
	nxt string
}

// UI registra as telas do jogo.
func UI() (*Game, error) {
	g := &Game{
		m: map[string]Screen{},
	}

	ld, err := loading.NewLoading(g.switchTo)
	if err != nil {
		return nil, err
	}

	g.Add("loading", ld)
	mn := menu.NewMenu(ld.Font(), g.switchTo)
	g.Add("menu", mn)
	g.Add("shop", shop.NewShop(ld.Font(), g.switchTo))

	g.cur = "loading"

	return g, nil
}

// Add registra uma tela pelo nome.
func (g *Game) Add(name string, s Screen) {
	g.m[name] = s
}

// switchTo agenda a troca para uma tela registrada.
func (g *Game) switchTo(name string) {
	if _, ok := g.m[name]; ok {
		g.nxt = name
	}
}

// Update atualiza a tela atual e aplica a troca pendente.
func (g *Game) Update() error {
	audio.Update()
	if err := g.m[g.cur].Update(); err != nil {
		return err
	}

	if g.nxt != "" {
		g.cur, g.nxt = g.nxt, ""
	}

	return nil
}

// Draw desenha a tela atual.
func (g *Game) Draw(img *ebiten.Image) {
	img.Fill(color.RGBA{8, 10, 15, 255})
	g.m[g.cur].Draw(img)
}

// Layout define a resolução lógica do jogo.
func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 720
}
