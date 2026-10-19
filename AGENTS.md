# AGENTS.md — Duckling (Go + Ebiten)

Idle 2D em Go 1.26 + Ebiten v2. Lógico 1280x720, fullscreen.
Jogo de dedução social: patos com tarefas, um impostor no meio.

## Estrutura (Engine / Game / Content)
- `main.go` — bootstrap: áudio, save, ícone, `screen.UI()`,
  `ebiten.RunGame`. Nada de lógica aqui.
- `engine/scene/` — router (`UI()`, `Add`, troca interna).
  É o ÚNICO lugar que conhece as telas (exceção à regra).
- `engine/rendering/` — controles pixel (`button`, `toggle`,
  `slider`, `icon`, `text`, `cover`, `anim`, `toast`,
  `particles`, `theme`). Sem estado de jogo, só desenho.
- `engine/audio/` — contexto + 2 canais (música e sfx, cada um
  com volume + liga/desliga) + `music.go` (loop) + `sfx.go`
  (tiro único, anti-spam 100ms, sem handle aberto).
- `engine/i18n/` — `i18n.go` + `locales/en|es|pt.json`.
- `engine/save/` — config local (`settings.json` em
  UserConfigDir): idioma, áudio, teclas e moedas. `Load/Apply`
  no boot, `Store` a cada mudança. Sem arquivo = defaults
  ligados.
- `engine/network/` — mensagens JSON por UDP (`msg.go`:
  `hello/join/deny/welcome/input/state/leave` + `start/role/
  phase/task/tasks/kill/sabotage`) + descoberta LAN (`discover.go`: anuncia e
  escuta no 48217) + socket da partida (`endpoint.go`: host e
  cliente no 48218).
- `game/session/` — partida LAN: host autoritativo (simula todos
  e distribui snapshot 15/s com corpos e cooldown) ou cliente
  (manda input 30/s e interpola). `Set/Get` guarda a sessão atual
  (nil = solo).
- `game/match/` — regras da partida: fases (`lobby/jogando/
  reunião/fim`), papéis (tripulante/impostor) e o plano que o
  host sorteia (spawn, 4+2 tarefas por jogador — fake pro
  impostor — e comparsas).
- `game/task/` — mecânica de uma tarefa sem tela/rede:
  `hold` (segurar pausa no soltar), `order` (soltar zera) e
  `carry` (pegar na origem, entregar no destino).
- `game/screens/loading|menu|play|shop/` — uma pasta
  por tela (`Update() error` + `Draw(*ebiten.Image)`).
  O curral (salas + novo jogo) é modal do menu.
- `game/player/` — `actor.Duck` (`duck.go` sprite,
  `move.go` física, `air.go` pulo/voo, `anim.go` direção,
  `defs.go` elenco, `color.go` matiz do pato, `feathers.go` penas).
- `game/map/` — TMX via go-tiled (`map.go` Load, `map_geo.go`
  geometria, `map_draw.go` 1 DrawImage, `maps.go` List/Open).
- `content/` — `embed` (fonts, images, maps, music, sprites).
  Chave de imagem = caminho (`images/bg`); faixa de áudio
  sempre minúscula (`Tracks` normaliza).

## Direção dos imports (lei)
- `engine/*` NUNCA importa `game/*`. Exceção única: `scene`
  registra as telas, esse é o trabalho dela.
- `content` é folha: não importa nada do projeto.
- `game/*` nunca importa `main`. Telas conversam por
  callback (`next func(string)`), nunca se importam.
- Nomes de package NÃO seguem a pasta (`engine/rendering`
  é `package ui`, `game/player` é `package actor`,
  `game/map` é `package world`, `content` é `package assets`).
  Não "conserte" isso: só o caminho do import muda.

## Contrato do mapa (Tiled, tudo opcional)
- Mapa em `maps/<nome>/<nome>.tmx` (`world.Open("nome")`).
- Tileset externo TEM que ser `.tsx` (JSON dá `EOF` no load).
- Grupo `collision` = sólidos (prop `h` = altura, sem `h` = 1000).
- Grupo `COLLISION_WALL` ou classe `BARRIER` = barreira
  absoluta (H 1000: nem pula, nem voa).
- Grupo `areas` (retângulos ≥128px) = zonas das teclas 1-5.
- Grupo `gates` = porteiras que a sabotagem fecha; viram colisão
  total (H 1000) e desenho enquanto durar.
- Objeto `spawn` = nascimento (sem ele, centro visível).
  Vários = um ponto por pato (o host distribui no começar).
- Grupo `tasks` = pontos de tarefa: prop `long` (bool), `time`
  (s) e `kind` = `hold` (segurar), `order` (segurar sem soltar,
  soltar zera) ou `carry` (levar de A até `to_x`/`to_y`);
  sem `kind` = `hold`. O host dá 4 curtas + 2 longas por
  tripulante.
- Props do mapa: `zoom` (gameplay), `maxplayers` (teto da sala).
- Offset da camada soma na origem do objeto (o código faz).
- Camada/tile desconhecido é ignorado sem erro.

## Boot e timing (pega-IA clássicos)
- Assets carregam AOS POUCOS no loading (`LoadOne` por job).
  Quem usa imagem/som no `New*` pega nil: carregue LAZY
  (padrão do pato no menu/play) ou depois do loading.
- `save.Apply(save.Load())` roda no `main` ANTES do `UI()`,
  então `newSettings()` já nasce com o estado salvo.
- `audio.PlayLoop` respeita `musicOn` (nasce pausado se off).
- `Tracks` é case-sensitive na origem: chame minúsculo
  (`"loop"`, não `"Loop"`); o decode normaliza mesmo assim.

## Regras
1. Arquivo curto: ideal ≤75 linhas, máximo 100. Estourou? Divide
   por responsabilidade (`lobby.go` virou 6 arquivos, é assim).
2. Menos código > mais código. Lib padrão + Ebiten
   (`ebitenutil`, `inpututil`, `text/v2`), sem reinventar.
3. Uma responsabilidade por arquivo. Sem lógica de jogo no
   `main.go`, no router e nos controles `rendering/`.
4. Sem reestruturar, renomear package/função pública, mudar
   assinatura usada em >1 lugar ou somar dependência sem pedir.
5. i18n obrigatório: todo texto visível usa `i18n.T("chave")`
   com chave em `en.json`, `es.json` E `pt.json`. Sem string
   hardcoded em `Draw`. Chaves: `dominio.nome` (ex: `menu.play`).
6. i18n é `embed` (`//go:embed locales/*.json`), sem
   `os.ReadFile` com path relativo. API: `T`, `Set`, `Toggle`,
   `Lang` (+ `Langs`, `Names` se o seletor precisar).
7. Nova tela = pasta própria em `game/screens/` + método
   `Update/Draw` + `g.Add("nome", ...)` no `UI()`.
8. `tmp/` é gitignore: binário de teste e programa headless
   vão em `tmp/`, nunca na raiz. Apague depois de usar.
9. Verificar antes de terminar: `go build ./...`, `go vet ./...`
   — tudo verde. Dúvida de runtime? programa headless em `tmp/`
   prova (load, física, codec, save com backup do real).
10. CHANGELOG antes do commit: toda mudança entra no
    `CHANGELOG.md` antes do `git add`. Sem changelog, sem commit.

## Comandos
- Run: `go run .` (ou `air`, config em `.air.toml`)
- Syso (ícone do .exe): `mingw32-make syso` (do `logo.png`; `build-win*` já garante)
- Verificação completa: `go build ./... && go vet ./...`
- Build win64: `GOARCH=amd64 go build -ldflags="-s -w" -o tmp/duckling-win64.exe .`
- Build win32: `go build -ldflags="-s -w" -o tmp/duckling-win32.exe .` (toolchain padrão é 386)
- Build linux64: `GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o tmp/duckling-linux64 .`
- Build linux32: `GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o tmp/duckling-linux32 .`
- Ou `mingw32-make build-win64|build-win32|build-linux64|build-linux32` (um arquivo por plataforma)
- Pack (UPX, ~-45%): `tmp/upxbin/upx.exe --best tmp/duckling-win*.exe`
- Build web: `GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o tmp/web/duckling.wasm .` + servir `tmp/web/`
