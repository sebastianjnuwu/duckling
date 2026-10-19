# 🦆 Duckling

<div align="center">
  <h1>Duckling</h1>
  <img alt="Duckling" src="https://raw.githubusercontent.com/sebastianjnuwu/duckling/refs/heads/idle/content/images/logo.png" width="25%" />
  <p>
    <b>🌽 Duckling</b> <i>é uma fazenda cheia de tarefas, patos suspeitos e milho voando!</i>
  </p>
  <hr>
</div>

<div align="center">
  <a href="https://github.com/sebastianjnuwu/duckling/blob/idle/LICENSE">
    <img alt="GitHub License"src="https://img.shields.io/github/license/sebastianjnuwu/cookie-clicker-brasil?style=for-the-badge&logo=apache&logoColor=4086c4&color=4086c4" />
  </a>
  <a href="https://github.com/sebastianjnuwu/duckling/archive/refs/heads/idle.zip">
    <img alt="GitHub Repo Size"src="https://img.shields.io/github/repo-size/sebastianjnuwu/cookie-clicker-brasil?style=for-the-badge&logo=databricks&logoColor=4086c4&color=4086c4" />
  </a>
 
</div>

Uma fazenda cheia de serviço e patos que não são lá muito confiáveis.

**Duckling** é um jogo 2D inspirado em descubrir quem é o impostor, feito em **Go + Ebiten**. Aqui, você cuida da fazenda, tenta descobrir quem está matando seus colegas e torce pra não ser o próximo.

Era pra todo mundo ajudar na colheita. Alguém entendeu “passar a foice” de outro jeito.

## 🌾 Como vai funcionar

Os patos da fazenda têm tarefas pra cumprir. Os impostores também, mas a deles é acabar com os outros patos sem levantar suspeitas.

Achou um corpo? Chame os outros e conte o que viu. Só cuidado: o pato que mais fala pode estar inventando tudo, e aquele que ficou quietinho pode só estar comendo milho.

O jogo ainda está em desenvolvimento. Por enquanto, estamos construindo a fazenda e dando aos patos motivos pra desconfiar uns dos outros.

## 🗺️ TODO — rumo ao jogo funcionando

O que já está pronto: 1 pato com física própria (corre/acelera/pula/voa) na fazenda (`play`), menu com salas reais na LAN (host/cliente + descoberta), loja "EM BREVE", rede UDP plugada, save de config, i18n em 3 idiomas. Falta o jogo em si: gente, papéis, tarefas, morte e voto. A ordem abaixo é a sugerida (cada fase destrava a próxima). Números como cooldowns e preços são chute inicial — o único lugar da verdade é o código citado em cada item.

### Fase 1 — Multiplayer de verdade (base de tudo)
- [x] Modelo host-autoritativo: o dono da sala simula e distribui; protocolo em `engine/network/msg.go` (`hello/join/welcome/deny/input/state/leave`) + versão do protocolo (cliente velho não entra)
- [x] Anúncio LAN valendo (`discover.go`): nome da sala, nº de jogadores, tem senha?, intervalo de 500ms; portas fixas 48217 (descoberta) e 48218 (partida)
- [x] Salas reais no curral: criar (nome/senha/min/max), entrar com senha, sair; se o dono sai, a sala fecha (sem migração por enquanto). Mapa ainda é fixo (`spawn`)
- [x] `menu_rooms.go` listando as salas achadas na rede (com botão de atualizar), com cadeado, paginação e toast de erro ("senha errada", "sala cheia")
- [x] Snapshot de estado 10–15/s com interpolação nos patos remotos (feito: 15/s + interpolação); pulo/voo vão como borda própria reenviada 1× (host filtra repetição em 120ms), kill e tarefa vão 3×; o voto usa o mesmo reenvio quando existir (Fase 5)
- [x] Timeout de 5s sem pacote = caiu (some do lobby e o cliente volta com aviso; fantasma de partida fica pra Fase 2)
- [x] Nome do jogador: campo de texto no lobby (teclado físico) — já existia e agora vale no join
- [x] Escolha de cor no lobby: 8 matizes por tint da arte amarela (`game/player/color.go`, a loja troca por sprite próprio na Fase 7), paleta no HUD do lobby e trava de repetido no host
- [x] Teste com 2 instâncias da GUI na mesma máquina (host+cliente headless também validados) — fecha a Fase 1

### Fase 2 — Partida e papéis
- [x] Máquina de estados da partida (`game/match/`): `lobby → jogando → reunião → fim`; o host sincroniza (`start/role/phase`) e a rede só anda dentro dela (reunião ainda sem gatilho: Fase 5)
- [x] Dono aperta "COMEÇAR" com no mínimo 4 patos (ajustável no lobby, piso 2 pra teste); sala trava pra quem chega no meio (`net.started`)
- [x] Sorteio pelo host: 1 impostor com 4–6 jogadores, 2 com 7+ (AUTO) ou escolha do dono; impostores se conhecem, tripulante não recebe nada
- [x] Segredo de verdade: o host manda um `role` por destino, com o spawn e as tarefas de cada um; tripulante nunca vê a lista de impostores
- [x] Contrato novo no Tiled (documentado no `AGENTS.md`): vários objetos `spawn` — um por pato, longe um do outro (8 no mapa)
- [x] Distribuir tarefas por tripulante (4 curtas + 2 longas, girando o pool); a interação/minigame é da Fase 3
- [x] Morte vira fantasma: flag no snapshot, desenho translúcido, sem sombra; hoje vale pra quem cai (kill real na Fase 4)
- [x] Quem cai no meio da partida vira fantasma automático (o pato não some do mundo)

### Fase 3 — Tarefas (o trabalho da fazenda)
- [x] Contrato Tiled: objetos de tarefa (`nome`, `tipo`, `tempo`) documentados no `AGENTS.md`, um por ponto da fazenda
- [x] 3 tipos pra começar: segurar-E com barra (colher milho), levar item de A até B (feno no celeiro), ordem certa (ordenhar na sequência)
- [x] Nova ação INTERAGIR no `engine/input` (remapeável, padrão E) — voo agora é só X
- [x] HUD das minhas tarefas + seta na borda apontando a próxima + barra global de progresso (a vitória em 100% é checada na Fase 6)
- [x] Moedas: curta +2, longa +5 (a vitória +10 fecha na tela de fim)
- [x] Impostor tem tarefas fake: animação igual, mas sem progresso e sem moeda
- [x] Som de tarefa pronta + toast + puff de pena (som reaproveitado até a Fase 8)

### Fase 4 — Impostor: faca, corpo e cooldown
- [x] Nova ação MATAR no `engine/input` (remapeável, padrão F): só impostor, alcance 46px, cooldown 25s com radial no HUD
- [x] Kill trava o impostor 0.5s + som + puff de pena da vítima; corpo caído no chão (por enquanto o próprio sprite deitado e escurecido no lugar de `duck_dead.png`)
- [x] 1ª leva de sabotagens (cooldown 30s, pausa na reunião): lamparina apagada 15s (overlay escuro com raio de visão por pato em `engine/rendering`) e porteiras fechadas 30s (grupo `gates` do Tiled)
- [x] Vitória do(s) impostor(es) quando vivos deles == vivos tripulantes (fantasma não conta)
- [x] Cooldowns pausam durante a reunião e resetam ao voltar pra fazenda

### Fase 5 — Reunião: corpo achado, dedo apontado
- [x] Reunião por cima do mundo congelado (`play_meet*.go`): motivo no topo (quem achou o corpo e o morto / quem apertou o sino), relógio da fase e o mundo parado — modal em vez de tela própria porque o host precisa continuar publicando o snapshot durante a reunião
- [x] Gatilhos: INTERAGIR perto de corpo reporta e perto do sino chama (remapeável, sem ação nova); sino no Tiled pelo grupo `bell`, 1 uso por jogador por partida
- [x] Lista de vivos com nome e cor; morto/fantasma aparece riscado e mudo (lista em até 2 colunas se passar de 8)
- [x] Discussão 60s com chat só de frases prontas (12 frases × `en/es/pt.json`, regra do i18n) — sem teclado livre
- [x] Votação aberta 30s com contagem ao vivo, pular permitido; empate ou maioria skip = ninguém sai
- [x] Ejetado revela o papel em toast/som, vira fantasma e a sabotagem ativa é desfeita; os corpos somem quando a reunião abre
- [ ] Teste manual com 2 instâncias da GUI: sino, reportar corpo, votar, ler o chat e ver a ejeção

### Fase 6 — Fim de jogo
- [x] Tela `game/screens/over/` (pasta própria): quem venceu e o motivo (tarefas 100% / impostores zerados / tripulantes ≤ impostores), no lugar do overlay mínimo que era do play
- [x] Lista completa de papéis (revelada pelo host no `phase`) + stats: tarefas feitas, mortes e os votos da última reunião (quem votou em quem / pular)
- [ ] Teste manual com 2 instâncias da GUI: ver a tela de fim, voltar pro lobby e conferir papéis/votos
- [x] Moedas creditadas e salvas (+2/+5 por tarefa, +10 pro lado vencedor na tela de fim); botões REVANCHE (mesma sala, volta ao lobby) e MENU (fecha a sala)
- [x] Condição checada no host a cada morte/ejeção/tarefa (`winLocked` no `Kill`, `MeetStep` e `DoneTask`); o cliente só lê o `Result` do `phase`/snapshot

### Fase 7 — Loja valendo
- [x] Botão LOJA ligado (`menu.go`): grade de cartões (sem chapéu + 3) com nome, pato de preview vestindo a peça, preço/estado (preço, EQUIPAR, EQUIPADO) e clique direto no cartão
- [x] Moedas e inventário no `settings.json` (`engine/save/hat.go`): moedas, chapéus comprados e o equipado — `Load/Apply` no boot, `Store` a cada compra/equipar
- [x] Catálogo: 3 chapéus desenhados por código sobre o sprite (sem arte nova); as 4 cores saíram do plano porque a paleta do lobby já dá as 8 cores de graça (matiz) — a loja vende só os chapéus
- [x] Preços: boné 30, palha 80, coroa 150 (sem chapéu = grátis); toast "sem moedas" quando não dá e "COMPRADO!" quando dá
- [x] Preview no próprio cartão (pato da loja, mesmo sprite) e chapéu sincronizado na rede (protocolo v10: `hat` no join/welcome/snapshot)
- [x] Chapéu aparece em todos: pato do menu, jogo (local e remotos), fantasma (translúcido) e lista da reunião
- [ ] Teste manual: comprar/equipar, ver o saldo cair e o chapéu em 2 instâncias (menu + jogo + reunião)

### Fase 8 — Conteúdo e polimento
- [ ] LOD de sprites: usar `16/32/64px` conforme o zoom (hoje só `128px` desenha; o resto só pesa o binário) + dar destino a `icons/duck/bottom/` (avatares da reunião? favicon web?)
- [~] Sons que faltam, um por evento: início, vitória e derrota já tocam os mp3 que existiam; kill, corpo achado, sino, voto, ejetado, tarefa pronta e sabotagem esperam os arquivos (`sfx.go` já tem 2 canais + anti-spam de 100ms)
- [~] Tela "como jogar" pronta (`game/screens/help/`, botão "?" no menu) com objetivo, papéis, tarefas, teclas remapeadas e dicas; faltam as placas de tutorial na 1ª fazenda
- [ ] 2º mapa de fazenda (celeiro + milharal) com spawns múltiplos, tasks próprias e `maxplayers`; `world.Open` já lista por pasta
- [ ] Acessibilidade: cor sempre acompanhada de símbolo/nome (daltônico), volumes já existem no settings
- [ ] Novas ações (interagir/matar/reportar) entram no remapeamento de controles (`engine/input` + 2 páginas do settings, com nomes no i18n)
- [ ] Tabela de balance inicial documentada aqui quando fechar: velocidade, voo, cooldown do kill (25s), timers da reunião (60/30s), nº de tarefas (4+2)
- [ ] Modo online (botão hoje desligado em `menu.go:70`): relay fora da LAN — por último, LAN primeiro
- [ ] Cada fase com seu teste headless em `tmp/` (física, protocolo, votação...) provando antes de marcar o checkbox

> Regra de ouro: fase fechada e verde (`go build ./... && go vet ./...`) antes de abrir a próxima. Sem changelog, sem commit.

## 🛠️ Feito com

- **Go + Ebiten** — pra fazer os patos andarem.
- **Tiled** — pra construir a fazenda.
- **Piskel** — pra desenhar os patos antes de acontecer alguma coisa com eles.

## 📋 Requisitos

- Go 1.26.
- Windows 10 ou 11.
- Tiled e Piskel, se quiser editar os mapas e as artes.

## ▶️ Bora jogar

Clone o repositório:

```bash
git clone https://github.com/sebastianjnuwu/duckling.git
```

Entre na pasta:

```bash
cd duckling
```

Baixe as dependências:

```bash
go mod download
```

Solte os patos:

```bash
go run .
```

## 📦 Compilar

Um arquivo por plataforma (saem em `tmp/`):

```bash
mingw32-make build-win64        # Windows 10/11 64-bit
mingw32-make build-win32        # Windows 10/11 32-bit
mingw32-make build-linux64      # Linux 64-bit
mingw32-make build-linux32      # Linux 32-bit
```

> Se perguntarem, você estava no celeiro.