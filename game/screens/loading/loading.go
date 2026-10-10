package loading

import (
	"duckling/content"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Loading struct {
	font     *text.GoTextFace
	big      *text.GoTextFace
	files    []string
	done     int
	shown    float64
	ticks    float64
	spin     float64
	finished bool
	wait     float64
	// elapsed/pace ritmam os jobs (~2s de 0→100%) + 3s de espera
	// = 5s mínimos de loading, mesmo com tudo carregado rápido.
	elapsed float64
	pace    float64
	goTo    func(string)

	pop     float64
	popVel  float64
	wasHeld bool
}

func NewLoading(next func(string)) (*Loading, error) {
	// Fonte primeiro (síncrona): o Draw usa Font()/BigFont() no frame 1,
	// antes dos jobs rodarem. Sem isso a Source vem nil e crasha.
	if err := assets.LoadOne("content/fonts/PressStart2P.ttf"); err != nil {
		return nil, err
	}
	// BG na hora: frame 1 já abre com fundo, sem preto.
	if err := assets.LoadOne("content/images/bg.png"); err != nil {
		return nil, err
	}
	files, err := assets.Collect(
		"content/fonts",
		"content/images",
		"content/music",
		"content/sprites",
	)
	if err != nil {
		return nil, err
	}

	pace := 0.0
	if len(files) > 0 {
		pace = 2.0 / float64(len(files))
	}

	return &Loading{
		files: files,
		pace:  pace,
		goTo:  next,
	}, nil
}

func (l *Loading) Update() error {
	l.ticks += 1.0 / 60.0
	l.spin += 2.5 / 60.0
	l.elapsed += 1.0 / 60.0

	// Um job por vez, ritmado: o % sobe aos poucos em vez de pular
	// de 0 para 100 no primeiro segundo.
	if !l.finished && l.done < len(l.files) && l.elapsed >= float64(l.done+1)*l.pace {
		if err := assets.LoadOne(l.files[l.done]); err != nil {
			return err
		}

		l.done++
	}

	// % conta de 1 em 1pp por frame até o progresso real: sobe
	// liso em vez de pular de arquivo em arquivo.
	if !l.finished && len(l.files) > 0 {
		target := float64(l.done) / float64(len(l.files)) * 100
		if l.shown < target {
			l.shown = min(target, l.shown+1.0)
		}
	}

	if !l.finished && l.done >= len(l.files) {
		l.finished = true
		l.shown = 100
		l.wait = 0
	}

	if l.finished {
		l.wait += 1.0 / 60.0

		// Sai no 100%: o título já está no topo (igual ao menu),
		// então o menu entra por cima com pato/botões animando.
		if l.wait >= 0.2 && l.elapsed >= 2 {
			l.goTo("menu")
		}
	}

	l.updatePop()

	return nil
}

func (l *Loading) updatePop() {
	held := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && l.mouseOnTitle()

	if held && !l.wasHeld {
		l.popVel = 0.35
	}
	l.wasHeld = held

	if held {
		l.popVel = min(l.popVel+0.03, 0.5)
	}

	l.popVel = (l.popVel - l.pop*0.15) * 0.85
	l.pop = clamp(l.pop+l.popVel, 0, 1)
	if l.pop == 0 || l.pop == 1 {
		l.popVel = 0
	}
}

func (l *Loading) Font() *text.GoTextFace {
	if l.font == nil {
		l.font = &text.GoTextFace{
			Source: assets.Fonts["PressStart2P"],
			Size:   18,
		}
	}
	return l.font
}

func (l *Loading) BigFont() *text.GoTextFace {
	if l.big == nil {
		l.big = &text.GoTextFace{
			Source: assets.Fonts["PressStart2P"],
			Size:   95,
		}
	}
	return l.big
}
