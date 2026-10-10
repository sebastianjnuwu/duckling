# AGENTS.md — Duckling (Go + Ebiten)

Idle 2D em Go 1.26 + Ebiten v2. Lógico 1280x720, fullscreen.

## Estrutura
- `main.go` — bootstrap Ebiten, nada de lógica aqui.
- `engine/scene/` — router (`UI()`, `Add`, troca interna).
- `engine/rendering/` — controles pixel (`button`, `toggle`,
  `slider`, `icon`, `text`, `cover`, `anim`). Sem estado de jogo.
- `engine/audio/` — 2 canais (música e sfx, volume + liga/desliga).
- `engine/i18n/` — `i18n.go` + `locales/en|es|pt.json`.
- `game/screens/loading|menu|play|shop/` — uma pasta por tela.
- `game/player/` — personagens (`actor.Duck` + física e animação).
- `game/map/` — mapa TMX via go-tiled (`Load`, `Draw`, áreas).
- `game/save/` — persistência.
- `content/` — `embed` (fonts, images, maps, music, sprites).

## Regras
1. Arquivo curto: ideal ≤75 linhas, máximo 100. Estourou? Divide.
2. Menos código > mais código. Lib padrão + Ebiten
   (`ebitenutil`, `inpututil`, `text/v2`), sem reinventar.
3. Uma responsabilidade por arquivo. Sem lógica de jogo no
   `main.go`, no router e nos controles `rendering/`.
4. i18n obrigatório: todo texto visível usa `i18n.T("chave")`
   com chave em `en.json`, `es.json` E `pt.json`. Sem string
   hardcoded em `Draw`. Chaves: `dominio.nome` (ex: `menu.play`).
5. i18n é `embed` (`//go:embed locales/*.json`), sem
   `os.ReadFile` com path relativo. API: `T`, `Set`, `Toggle`,
   `Lang` (+ `Langs`, `Names` se o seletor precisar).
6. Telas implementam `Update() error` + `Draw(*ebiten.Image)`.
   Nova tela = pasta própria + `g.Add("nome", ...)` no `UI()`.
7. `go build ./...` (e `go vet ./...`) antes de terminar.
   Binário de teste vai em `tmp/`, nunca na raiz.
8. CHANGELOG antes do commit: toda mudança entra no
   `CHANGELOG.md` antes do `git add`. Sem changelog, sem commit.

## Comandos
- Run: `go run .` (ou `air`, config em `.air.toml`)
- Build win64: `GOARCH=amd64 go build -ldflags="-s -w" -o tmp/duckling-win64.exe .`
- Build win32: `go build -ldflags="-s -w" -o tmp/duckling-win32.exe .` (toolchain padrão é 386)
- Pack (UPX, ~-45%): `tmp/upxbin/upx.exe --best tmp/duckling-win*.exe`
- Build web: `GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o tmp/web/duckling.wasm .` + servir `tmp/web/`
