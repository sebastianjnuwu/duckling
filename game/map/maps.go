package world

import (
	"io/fs"

	assets "duckling/content"

	"github.com/lafriks/go-tiled"
)

// maps.go — descoberta: nome resolve pra maps/<nome>/<nome>.tmx.
// Camada vale por NOME ou CLASSE (o resto é ignorado, sem erro).

// List devolve os mapas embutidos (ex: ["spawn"]).
func List() ([]string, error) {
	dirs, err := fs.ReadDir(assets.FS, "maps")
	if err != nil {
		return nil, err
	}
	var out []string
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		if _, err := fs.Stat(assets.FS, "maps/"+d.Name()+"/"+d.Name()+".tmx"); err == nil {
			out = append(out, d.Name())
		}
	}
	return out, nil
}

// Open carrega pelo nome (ex: Open("spawn")).
func Open(name string) (*Map, error) {
	return Load("maps/" + name + "/" + name + ".tmx")
}

// groupIs diz se o grupo é o nomeado (por nome ou classe).
func groupIs(g *tiled.ObjectGroup, name string) bool {
	return g.Name == name || g.Class == name
}
