package assets

import (
	"bytes"
	"embed"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed fonts images maps music sprites
var FS embed.FS

var Images = map[string]*ebiten.Image{}
var Tracks = map[string]string{}
var Fonts = map[string]*text.GoTextFaceSource{}

func noFS(path string) string {
	return strings.TrimPrefix(path, "content/")
}

func Collect(roots ...string) ([]string, error) {
	var files []string

	for _, root := range roots {
		err := fs.WalkDir(FS, noFS(root), func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if entry.IsDir() {
				return nil
			}

			switch strings.ToLower(filepath.Ext(path)) {
			case ".ttf", ".png", ".jpg", ".jpeg", ".gif", ".mp3", ".ogg", ".wav":
				files = append(files, "content/"+path)
			}

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return files, nil
}

func Load() error {
	files, err := Collect(
		"content/fonts",
		"content/images",
		"content/maps",
		"content/music",
		"content/sprites",
	)
	if err != nil {
		return err
	}

	for _, path := range files {
		if err := LoadOne(path); err != nil {
			return err
		}
	}

	return nil
}

func LoadOne(path string) error {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".ttf":
		return loadFont(path)

	case ".mp3", ".ogg", ".wav":
		name := strings.ToLower(strings.TrimSuffix(
			filepath.Base(path),
			filepath.Ext(path),
		))

		Tracks[name] = path
		return nil

	case ".png", ".jpg", ".jpeg", ".gif":
		return loadImage(path)
	}

	return nil
}

func loadFont(path string) error {
	data, err := fs.ReadFile(FS, noFS(path))
	if err != nil {
		return err
	}

	source, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		return err
	}

	name := strings.TrimSuffix(
		filepath.Base(path),
		filepath.Ext(path),
	)

	Fonts[name] = source

	return nil
}

func loadImage(path string) error {
	file, err := FS.Open(noFS(path))
	if err != nil {
		return err
	}
	defer file.Close()

	source, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Chave = caminho relativo (sprites/duck/yellow/duck):
	// basename colidiria entre pastas (red/ x yellow/).
	name := strings.TrimSuffix(
		noFS(path),
		filepath.Ext(path),
	)

	Images[name] = ebiten.NewImageFromImage(source)

	return nil
}
