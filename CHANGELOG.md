# Changelog — Duckling

## [Unreleased]
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
