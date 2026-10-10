# Changelog — Duckling

## [Unreleased]

## [0.2.0] - 2026-10-10
### Added
- Reestrutura Engine/Game/Content (`src/ui`→`engine/rendering`,
  `src/audio`→`engine/audio`, `src/i18n`→`engine/i18n`,
  router→`engine/scene`, telas→`game/screens/*`,
  `src/actor`→`game/player`, `src/world`→`game/map`,
  `assets`→`content`); nomes dos packages intactos
- Config local (`src/save/`: idioma, música/sfx on+vol em
  `settings.json` versionado, carrega no boot e grava a cada
  mudança — escrita atômica, só se mudou, corrompido vira backup)
- Câmera mira o sprite (não os pés); no alto o pato ganha folga
  na trava de cima pra cabeça não sair da tela (sem mostrar o vazio)
- Colisão com altura (`world.Solid{H}`): objeto `collision` no
  Tiled aceita propriedade custom `h` (cerca=12, muro=vazio=1000);
  pulando acima de H o pato passa por cima; voando vale o mesmo
  (cruzeiro a 34px: muro com H maior barra até o voo; colado
  em paredão, nem levanta — o voo corta e ele escorrega no muro)
- Voo com tato: decolagem que amortece no cruzeiro, flap forte
  subindo e planeio no cruzeiro/queda, teto que desce no último
  1s de fôlego (plana até o chão), queda com velocidade terminal,
  quique só no 1º encosto e câmera com altitude amortecida
- Voo direcional: voando mira pra onde vai (lado = fileira de
  voo espelhada, frente/costas reaproveitadas com flap)
- Sombra de voo no pato (`actor/shadow.go`: blob pixel arredondado
  em 3 faixas perto/médio/alto + contato discreto no chão, flap 2x mais
  rápido e cruzeiro em ~34px; altura do sprite agora escala com
  o zoom; ESPAÇO pula, E/X segura ou alterna pra voar (1.3x))
- Flap de verdade (`ui.NewFlap`: alterna fileira asa-baixa com
  asa-erguida; `Def.FlapUp`; subindo rápido, cruzeiro devagar)
- NPCs vermelhos (`actor.NPC`, `Cast["red"]`): 3 passeando perto
  do spawn; trombou no chão, freia a meio passo por 1.2s (o NPC
  é quem sai da frente, com 1.5s de carência anti-trava; NPC
  que encosta no muro já vira (sem trimilique); trombada com
  tropeço (pulinho), NPC foge 1s assustado e puff de pena)
- Toast `ui.Toast` reutilizável (JOGAR e LOJA avisam EM BREVE,
  com quebra de linha responsiva) e faíscas no botão da loja
- Mapa TMX do Tiled via `go-tiled` em `src/world/` (`Load` +
  `Draw` em 1 chamada, spawn por objeto, `maps/` já no embed)
- Tela `play`: JOGAR abre o mapa spawn (pato anda no WASD,
  câmera segue, ESC volta, sem footer)
- Pato direcional em `src/actor/` (`actor.Duck`; a pasta recebe
  todos os personagens): anda de frente nas 4 direções
  (espelha no A/D, fileira 2 no W), parado congela no frame base;
  as fileiras de lado da sheet são voo (só o menu usa)
- Física no actor: `Duck.Update` lê WASD/setas, move, anima e
  prende nos limites (a tela só importa e chama)
- Colisão: grupo `collision` no Tiled vira sólidos; bateu, o
  pato quica pra trás com atrito (passo fatiado em 1px)
- Animação: 1º toque troca a pose na hora (`Strip.Step`),
  parado congela no frame base; `actor/anim.go` separado
- Pulo e voo (`actor/air.go`): ESPAÇO pula, segurar X voa 3s
  (recarrega no chão, passa por cima dos sólidos, fileira 0)
- Props custom como API (`world`): `zoom` no mapa, `h` nos
  sólidos — tudo opcional, mapa antigo continua igual
- Zoom 1.4x no mapa (`world.Zoom`, tiles ~22px, ilha respira
  na tela com `bg.png` em volta)
- Zoom 4x estilo Stardew no pato (tiles de 64px, ~20x11
  tiles por tela; `world.Zoom` continua o botão do zoom)
- Elenco data-driven (`actor.Cast`: fileiras, escala e
  velocidade por personagem)
- Sprites em subpastas (`duck/red|yellow/`, normal + 4x): chave
  da imagem virou o caminho (`sprites/duck/yellow/duck-4x`),
  `Def` ganhou `Cell/Cols` e o jogo usa o 4x (`128px`, `Scale 0.25`)
- Intro cinematográfica: abre no mapa todo e desce em 2.5s
  (smoothstep) até o pato; pato escala junto com o zoom
- Áreas no mapa: object layer `areas` no Tiled vira 5 mini-mapas
  (teclas 1-5 pulam a câmera entre elas, com trava de borda)
- Spawn e câmera no centro VISÍVEL da ilha (varredura de alpha,
  não do mapa vazio) e `ui.NewDuck(img, fileira)` com os 5
  personagens do spritesheet (24/10/8/13/10 frames)
### Fixed
- Loading abre com o fundo no frame 1 (`bg.png` junto da fonte,
  sem tela preta na espera ritmada)
- Android: APK instala do 7 ao 16 (minSdk 24, targetSdk 35
  via gomobile patchado, assinatura v2/v3, ícone `assets/icon.png`,
  landscape travado, `make android` faz tudo)
### Removed
- Suporte Android (só Windows + web): saíram `platform/`,
  `assets/icon.png` e os alvos `android*` do Makefile
### Docs
- README reescrito (pitch, controles, builds, estrutura)

## [0.1.0] - 2026-10-07
### Added
- Menu: título, patão animado na grama, JOGAR/LOJA/OPÇÕES/SAIR
  com botões e ícones pixel (hover ouro, clique afunda)
- Loading com barra, título e música em loop
- Loja (placeholder com volta)
- Opções em páginas: idioma (BR/US/ES em botões-bandeira),
  música e SFX com liga/desliga + volume próprio de arrastar
- i18n PT/EN/ES embutido (`embed`), troca com L ou bandeiras
- Áudio em 2 canais: música (loop) e sfx (tiro único)
- Builds: win64, win32, web (wasm) e Android (APK minSdk 24,
  targetSdk 35, landscape)
- Makefile (`run`, `vet`, `build-*`, `web`, `android`, `pack`,
  `size`, `clean`) e AGENTS.md
