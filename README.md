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

O que já está pronto: 1 pato com física própria (corre/acelera/pula/voa) na fazenda (`play`), menu com salas mock, loja "EM BREVE", rede UDP pronta mas desplugada, save de config, i18n em 3 idiomas. Falta o jogo em si: gente, papéis, tarefas, morte e voto. A ordem abaixo é a sugerida (cada fase destrava a próxima). Números como cooldowns e preços são chute inicial — o único lugar da verdade é o código citado em cada item.

### Fase 1 — Multiplayer de verdade (base de tudo)
- [ ] Modelo host-autoritativo: o dono da sala simula e distribui; protocolo em `engine/network/msg.go` com `hello/join/leave/state/event` + versão do protocolo (cliente velho não entra)
- [ ] Anúncio LAN valendo (`discover.go`): nome da sala, nº de jogadores, tem senha?, mapa, intervalo de broadcast e porta fixa documentados
- [ ] Salas reais no curral: criar (nome/senha/min/max/mapa), entrar com senha, sair; se o dono sai, a sala fecha (sem migração por enquanto)
- [ ] `menu_rooms.go` listando as salas achadas na rede (hoje é mock), com cadeado, paginação e toast de erro ("senha errada", "sala cheia")
- [ ] Snapshot de estado 10–15/s com interpolação nos patos remotos; eventos (pulo/kill/tarefa/voto) por mensagem com reenvio simples
- [ ] Timeout de 5s sem pacote = caiu (vira fantasma na partida, some no lobby) + toast avisando quem caiu
- [ ] Nome do jogador: campo de texto no lobby (hoje não existe input de texto — fazer um mínimo com teclado físico)
- [ ] Escolha de cor no lobby com trava de repetido (cores = pastas `sprites/duck/<cor>/`, hoje só existe `yellow`)
- [ ] Teste com 2 instâncias na mesma máquina antes de fechar a fase

### Fase 2 — Partida e papéis
- [ ] Máquina de estados da partida (ex: `game/match/` novo): `lobby → jogando → reunião → fim`; rede só sincroniza dentro dela
- [ ] Dono aperta "começar" com no mínimo 4 patos; sala trava pra quem chega no meio (vira espectador ou volta ao menu)
- [ ] Sorteio pelo host: 1 impostor com 4–6 jogadores, 2 com 7+; impostores se conhecem, tripulante não sabe de nada
- [ ] Segredo de verdade: o host nunca manda a lista de impostores pro cliente tripulante (mensagens filtradas por destino, não é só esconder na UI)
- [ ] Contrato novo no Tiled (documentar no `AGENTS.md`): vários objetos `spawn` — um por pato, longe um do outro
- [ ] Distribuir tarefas por tripulante (ex: 4 curtas + 2 longas, sem repetir o mesmo ponto pra todo mundo)
- [ ] Morte vira fantasma: anda e vê tudo, mas não fala, não vota, não faz tarefa, não é visto como vivo
- [ ] Quem cai no meio da partida vira fantasma automático (não some o corpo do jogo)

### Fase 3 — Tarefas (o trabalho da fazenda)
- [ ] Contrato Tiled: objetos de tarefa (`nome`, `tipo`, `tempo`) documentados no `AGENTS.md`, um por ponto da fazenda
- [ ] 3 tipos pra começar: segurar-E com barra (colher milho), levar item de A até B (feno no celeiro), ordem certa (ordenhar na sequência)
- [ ] Nova ação INTERAGIR no `engine/input` (remapeável, padrão E) — atenção: E hoje é voo alternativo (`air.go` lê `KeyE` direto); voo passa a ser só X e E vira interagir
- [ ] HUD das minhas tarefas + seta na borda apontando a próxima + barra global de progresso (vitória tripulante em 100%)
- [ ] Moedas: curta +2, longa +5, vitória +10 (números iniciais, mexer depois)
- [ ] Impostor tem tarefas fake: animação igual, mas sem progresso e sem moeda
- [ ] Som de tarefa pronta + toast + puff de pena (reaproveitar `feathers.go`/`toast.go`)

### Fase 4 — Impostor: faca, corpo e cooldown
- [ ] Nova ação MATAR no `engine/input` (remapeável, padrão F): só impostor, alcance ~24px, cooldown 25s com radial no HUD
- [ ] Kill trava o impostor 0.5s + som + puff de pena da vítima; corpo caído no chão (arte nova: `duck_dead.png` por cor)
- [ ] 1ª leva de sabotagens (com cooldown próprio): fechar porteiras (tranca as `areas` por 30s) e apagar a lamparina (overlay escuro com raio de visão — controle novo em `engine/rendering`)
- [ ] Vitória do(s) impostor(es) quando vivos deles == vivos tripulantes (fantasma não conta)
- [ ] Cooldowns pausam durante a reunião e resetam ao voltar pra fazenda

### Fase 5 — Reunião: corpo achado, dedo apontado
- [ ] Nova tela `game/screens/meet/` (padrão da casa: `Update/Draw` + `Add` no `UI()`): motivo no topo (quem achou o corpo e onde / quem apertou o botão do celeiro)
- [ ] Gatilhos: ação REPORTAR perto de corpo + botão de emergência no celeiro (objeto no Tiled, 1 uso por jogador por partida)
- [ ] Lista de vivos com avatar (usar `images/icons/duck/transparent/`?), nome e cor; morto/fantasma aparece riscado e mudo
- [ ] Discussão 60s com chat só de frases prontas (~12 frases × `en/es/pt.json`, regra do i18n) — sem teclado livre
- [ ] Votação aberta 30s com contagem ao vivo, pular permitido; empate ou maioria skip = ninguém sai
- [ ] Ejetado revela o papel ("era / não era o impostor"), vira fantasma, sabotagem ativa é desfeita

### Fase 6 — Fim de jogo
- [ ] Nova tela `game/screens/over/`: quem venceu + motivo (tarefas 100% / impostores zerados / tripulantes ≤ impostores)
- [ ] Revelar a lista completa de papéis + stats da partida (tarefas feitas, kills, quem votou em quem)
- [ ] Moedas creditadas e salvas; botões revanche (mesma sala, volta ao lobby) e menu
- [ ] Condição checada no host a cada morte/ejeção/tarefa (não no cliente — anti-"vitória fantasma")

### Fase 7 — Loja valendo
- [ ] Desligar o `Disabled` do botão (`menu.go:71`) e tirar o "EM BREVE" (`shop.go`): grade de itens com preço, preview e botão comprar/equipar
- [ ] Moeda e inventário no `settings.json` (`engine/save`): moedas, skin equipada, itens comprados — `Load/Apply` no boot, `Store` a cada mudança
- [ ] Catálogo inicial: 4 cores de pato (exige desenhar `sprites/duck/<cor>/` com as mesmas 12 tiras) + 3 chapéus (arte nova desenhada sobre o sprite no `Draw`)
- [ ] Preços iniciais: cor 50, chapéus 30/80/150; toast "sem moedas" quando não dá
- [ ] Preview usando o pato do menu (`menu_media.go`); skin sincronizada na rede (campo no `hello`/`state` da Fase 1)
- [ ] Chapéu aparece em todos: reunião, fazenda e fantasma

### Fase 8 — Conteúdo e polimento
- [ ] LOD de sprites: usar `16/32/64px` conforme o zoom (hoje só `128px` desenha; o resto só pesa o binário) + dar destino a `icons/duck/bottom/` (avatares da reunião? favicon web?)
- [ ] Sons que faltam, um por evento: kill, corpo achado, sino da emergência, voto, ejetado, vitória, derrota, tarefa pronta, sabotagem (2 canais + anti-spam, padrão do `sfx.go`)
- [ ] Tela/modal "como jogar" + placas de tutorial na 1ª fazenda; i18n de TODA string nova (`en/es/pt.json`, sem hardcoded no `Draw`)
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