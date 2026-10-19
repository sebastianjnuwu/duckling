# Changelog — Duckling

## [Unreleased]
### Added — Fase 3 (tarefas)
- Ação **INTERAGIR** (padrão `E`) em `engine/input/` (remapeável e
  salva em `settings.json`); o voo agora é só a tecla configurada
  (`X`), sem o atalho antigo no `E`
- Ajustes: 3ª linha de teclas na página de controles com ícone novo
  (`IconHand`) e a chave `settings.interact` em en/es/pt
- Tiled: `world.Task` ganhou `Kind` (`hold`/`order`/`carry`) e
  `Dest` (prop `to_x`/`to_y` do carry) + `Center()`/`Target()`;
  `spawn.tmx` agora tem 5 curtas (4 milho `hold` + 1 feno `carry`)
  e 2 longas (celeiro `hold`, ordenha `order`)
- `game/task/`: mecânica sem tela/rede (`hold` pausa ao soltar,
  `order` zera ao soltar, `carry` pega e entrega); `play_task.go`
  acha o ponto mais perto (46px) e roda com a ação INTERAGIR
- Protocolo v3: `task` (convidado avisa que terminou, reenviado 3x)
  e `tasks` (host manda concluídas/total); o host só aceita tarefa
  do próprio tripulante, ignora impostor/fantasma/duplicada e
  recalcula o total quando alguém cai (fantasma sai da conta)
- Impostor recebe a mesma lista (tarefa fake): a interação e a
  animação são iguais, mas o host não conta o progresso dele
- HUD de tarefas: barra global do time no topo, lista das minhas
  com ✓ (nomes traduzidos por `task.*`) e seta na borda apontando
  pra tarefa mais perto fora da tela
- Moedas no `settings.json` (`save.Coins/AddCoins`): curta +2,
  longa +5; ao concluir toca o sfx, mostra o toast (`task.done` +
  moedas) e solta um puff de penas no ponto (som próprio na Fase 8)
### Added — Fase 2 (partida e papéis)
- `game/match/`: máquina de fases (`lobby/jogando/reunião/fim`),
  sorteio de papel (AUTO: 1 impostor até 6 jogadores, 2 a partir
  de 7; ou override do dono), spawn por pato e 4 curtas + 2 longas
  por tripulante (`Plan`)
- `game/session/`: `Start/Reset/Finish`; o `role` vai filtrado por
  destino (spawn/tarefas/comparsas só pra quem é dono deles), a
  sala trava depois de começar (`net.started`) e quem cai no meio
  vira fantasma (`Ghost` no snapshot)
- `game/screens/play/`: painel do dono (COMEÇAR com contagem e
  mínimo, ENCERRAR), HUD de fase/vivos/papel (comparsas só pro
  impostor), fim mínimo com volta ao lobby e pato fantasma
  translúcido (`ui.Strip.DrawAlpha`)
- Tiled: 8 objetos `spawn` e grupo `tasks` (4 curtas + 2 longas)
  em `spawn.tmx`; `world.Map.Spawns()` e `world.Map.Tasks`
- Lobby: mínimo padrão 4 (piso 2 pra teste) e impostores AUTO (0)
  com override; `welcome` leva o mínimo da sala
- Protocolo v2 (`engine/network/msg.go`): `start/role/phase`,
  `Player.Ghost` e os campos de spawn/tarefas
### Added — Fase 1 (LAN)
- LAN jogável (Fase 1): host autoritativo + cliente em `game/session/`
  (join/welcome, snapshot 15/s, input 30/s, timeout 5s, match de
  versão) sobre UDP no `GamePort` 48218 (`engine/network/endpoint.go`);
  anúncio/descoberta continuam no 48217
- Lobby real: CRIAR vira host e anuncia na LAN; browse lista as salas
  achadas (com botão de atualizar), entra com nome + senha e mostra
  toasts de erro (`net.*`); sair da tela para a varredura
- Pato remoto: `actor.In` desacopla a física do teclado (entrada local
  ou da rede), `Interp`/`Snap` em `game/player/remote.go`; o host
  simula todos, o cliente prevê o próprio e interpola os outros
- i18n de `lobby.join`, `lobby.scanning` e `net.*` em en/es/pt
### Changed
- `Duck.Update` recebe `actor.In` (não lê mais `input` direto); a
  tecla E e o buffer de voo ficaram em `actor.LocalIn`
- `Menu.room` virou ponteiro (`RoomState` guarda mutex da varredura)
### Added
- TODO detalhado no `README.md` (Fases 1–8 com arquivos-alvo,
  números iniciais e contratos de mapa/rede por fase)
### Changed
- Física do pato com peso (`move.go`/`air.go`/`defs.go`): aceleração
  (chão 0.9/ar 0.4/voo 0.55), diagonal normalizada (1/√2),
  velocidade 3→2.6, pulo com altura variável (soltar corta em 2.5),
  buffer 0.12s + coyote 0.10s, pulo 5.6→5.2/gravidade 0.5
  (pico ~30, ainda limpa `h=25`), voo com decolagem suave
  (sobe ≤1.8/frame), mergulho segurando baixo e regen 2.0/s
### Removed
- NPCs vermelhos removidos de vez (`game/player/npc.go`,
  `game/screens/play/play_npc.go`, lista `npcs` no play,
  trombada/empurrão `slowT/bumpT/kx/ky/hx/hy` no pato):
  sem passeadores, sem freio ao encostar
### Changed
- Sprites do pato em 12 tiras por tamanho
  (`sprites/duck/yellow/{16,32,64,128}px/`: `down/up/left/right`,
  `fly_*`, `jump_*`); `NewHStrip` fatia pelo N real de cada PNG
- Docs atualizados (`defs.go`, `assets.go`, `feathers.go`, `AGENTS.md`)
  pro layout novo; ícone do `.exe`/janela usa
  `images/icons/duck/transparent/` com fallback pra `images/logo.png`
### Added
- Física 2.5D/isométrica de elevação (`game/player/floor.go`): sólidos
  com altura (cercas, pedras) agora funcionam como plataformas elevadas;
  o pato pousa no topo (`z = s.H`), caminha sobre elas, pula de cima
  e cai suavemente pela gravidade ao sair da borda
- Projeção de sombra em superfícies elevadas (`shadow.go`)
- Constantes exportadas `JumpPeak` (25) e `FlyPeak` (34) em `air.go`;
  docs em `Solid`, `map.go` e `AGENTS.md` agora mostram os limiares
  reais (`h≤25` pula, `h≤34` voa, `h>34` muro)
### Fixed
- Pato travava ou trepidava ao andar encostado em paredes/cercas:
  removido ricochete artificial em `step()`, agora desliza suavemente
- Pato caía para `z = 0` penetrando dentro de obstáculos com altura
  e ficava travado permanentemente; agora pousa e caminha sobre o topo
- `COLLISION_WALL` ignorava a prop `h` e forçava H=1000 (as cercas
  com `h=20` estavam nesse grupo); agora `solidH()` respeita `h`/`H`
- Controles com ícones nas labels (setas, espaço, voo) e REDEFINIR
  largo com ícone (cabe o RESTABLECER); confirma "RESTAURAR
  TECLAS?" (SIM/NÃO/X/ESC, padrão do sair)
- Ícone do `.exe` regenerável (`mingw32-make syso`: go-winres via
  `go run`, sem instalar nada; `build-win*` garante o `.syso`)
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
