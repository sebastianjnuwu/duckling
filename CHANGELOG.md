# Changelog — Duckling

## [Unreleased]
### Added
- Controles remapeáveis (`engine/input` + página nas opções):
  WASD/setas, pulo e voo; salva em `settings.json`
- Controles em 2 páginas (mover/ações) com título; shift, F1-F12
  e mouse até X2 (5 botões); tecla repetida troca de lugar
- Botão RESET (volta WASD/espaço/X); título único por página
- `AGENTS.md` detalhado (contratos de mapa, boot, regras);
  testes `internal/arch/` removidos
- Reestrutura Engine/Game/Content (`src/ui`→`engine/rendering`,
  `src/audio`→`engine/audio`, `src/i18n`→`engine/i18n`,
  router→`engine/scene`, telas→`game/screens/*`,
  `src/actor`→`game/player`, `src/world`→`game/map`,
  `assets`→`content`); nomes dos packages intactos
- Config local (`game/save/`: idioma, música/sfx on+vol em
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
- Sombra de voo no pato (`game/player/shadow.go`: blob pixel arredondado
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
- Mapa TMX do Tiled via `go-tiled` em `game/map/` (`Load` +
  `Draw` em 1 chamada, spawn por objeto, `maps/` já no embed)
- Tela `play`: JOGAR abre o mapa spawn (pato anda no WASD,
  câmera segue, ESC volta, sem footer)
- Pato direcional em `game/player/` (`actor.Duck`; a pasta recebe
  todos os personagens): anda de frente nas 4 direções
  (espelha no A/D, fileira 2 no W), parado congela no frame base;
  as fileiras de lado da sheet são voo (só o menu usa)
- Física no actor: `Duck.Update` lê WASD/setas, move, anima e
  prende nos limites (a tela só importa e chama)
- Colisão: grupo `collision` no Tiled vira sólidos; bateu, o
  pato quica pra trás com atrito (passo fatiado em 1px)
- Animação: 1º toque troca a pose na hora (`Strip.Step`),
  parado congela no frame base; `game/player/anim.go` separado
- Pulo e voo (`game/player/air.go`): ESPAÇO pula e X voa
  (recarrega no chão, passa por cima dos baixos, fileira 0);
  sem penas no pulo/pouso
- Penas (`game/player/feathers.go`): amarelas+vermelhas no pulo,
  no pouso forte e rastro no voo (`Sparks` com gravidade e
  desenho com câmera)
- Props custom como API (`world`): `zoom` no mapa, `h` nos
  sólidos — tudo opcional, mapa antigo continua igual
- Barreira absoluta: grupo `COLLISION_WALL` ou classe `BARRIER`
  (H 1000 sempre: nem pula, nem voa por cima)
- Fix áudio mudo: chamadas usavam `"Loop"`/`"Click"` e as faixas
  são minúsculas (mapa é case-sensitive); chave normalizada
- SFX de pulo (`jump.mp3` no ESPAÇO)
- Modo online/lan: JOGAR abre modal (globo + ícone LAN novo);
  online avisa em breve, LAN entra no jogo (solo até o netcode)
- Sala de espera (`game/screens/lobby/`, mock offline): min/max
  2-8, privada/pública com senha gerada, lista com HOST em
  destaque, START entra no jogo
- Curral: tabelinha de salas (criar + entrar, cadeado nas
  privadas) antes da config; textos em clima de fazenda
- Lobby começa vazio ("NENHUM CURRAL POR PERTO"), ícone de mais
  no criar, voltar sem ícone; ícone `IconPlus` novo
- Curral com X pra fechar, paginação de 4 em 4 (`<`/`>` + `1/2`),
  ícone em toda linha (cadeado nas privadas, rede nas abertas)
- Curral virou modal do menu (fundo não some); grade sem
  número mágico (`menu_room_layout.go`); "NOVO JOGO", senha
  grande estilo código de sala, ícone no nome
- Browse no mesmo painel da config (X respira igual); senha
  contida (pílula normal, sem exagero)
- Nome da sala (campo editável, 3-8 caracteres, CRIAR só entra
  válido) e senha na direita em pílula; layout da config refeito
- Impostores na config (1 a cada 6: 6→1, 12→2, 14→3, trava no
  mínimo); lobby fatiada em `lobby_{cfg,draw,rooms,rule,ui}.go`
- Modal refinado (`ui.Panel`): sombra, canto recortado, filete
  de latão e régua dourada — mesmo idioma dos botões
- X desgrudado da borda em todos os modais + `ui/theme.go`
  (paleta central, o "CSS" do jogo; `Panel` já usa)
- X menor com ar e `Btn.Push()`: fechar modal mostra o afundar
  antes de trocar de tela; respiro igual (36x20) em todos
- Senha alinhada na coluna dos controles; CRIAR da tabela
  alargado (texto não estoura)
- Títulos com ícone (`ui.Title`) e palavras de fazenda: sair
  da fazenda, ajustes da fazenda, curral, modo de jogo, sede
- Título com ícone do tamanho da fonte; config com ícone criar
- Botão desabilitado (`Disabled`: sem hover/clique, apagado) no
  ONLINE e na LOJA; `bevel` foi pro `theme.go`
- Ícones custom: PNG branco em `content/images/icons/<nome>.png`
  (tintado sozinho, cai no procedural se faltar)
- Ícones por nome via fonte (`MaterialSymbolsRounded` embutida,
  codepoints verificados um a um; PNG > fonte > procedural)
- Copiar senha de verdade (`engine/clip`, Win32 sem cgo):
  refresh troca, copiar joga no clipboard (testado ida e volta)
- Toast com entrada e pilha: desliza de baixo com fade-in,
  até 3 empilhados (sai o velho); `wrap` foi pro `text.go`
- Toast na senha ("NOVA SENHA XXXXX GERADA", "SENHA COPIADA")
  e senha alfanumérica sem 0/O/1/I
- Só fonte: procedural removido, todo ícone vem da fonte
  (PNG custom continua passando na frente se existir)
- Steppers somem no limite; paginação some com <2 páginas;
  `cap`→`top` (sombreava builtin); Makefile com linux64/32
  (um arquivo por plataforma) + revisão geral (i18n 100%)
- Config: botão vira CRIAR, voltar desgrudado, X no canto;
  steppers e labels de segurança com ícones
- Rede (`engine/network`, testado local): mensagens JSON UDP
  (`hello/join/deny/state`) + descoberta LAN (anuncia/escuta)
- Modal fecha ao escolher; ESC no jogo pede confirmação
  (`play.quit` + SIM/NÃO); `ui.Dim/Panel` reutilizáveis
- Bump empurra: trombou em NPC, o pato quica pra trás (testado
  perto/longe) + `bumped.mp3` no encosto
- Fix música ignorando mudo: `PlayLoop` tocava sem olhar o
  `musicOn` (só respeitava ao religar); agora nasce pausado
- Player não trava player: corpos não bloqueiam (só o `Bump`
  de encostar); `Def.Height` fica pra uso futuro
- Fix: offset da camada de objetos somado na origem (muros
  saiam deslocados e bloqueavam lugar errado)
- Descoberta: `world.List/Open` por nome (`maps/<nome>/<nome>.tmx`),
  camada vale por nome OU classe, resto ignorado sem erro
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
