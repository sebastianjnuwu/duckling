# Duckling

Idle 2D em Go + Ebiten: menu com pato animado, loja,
opções (3 idiomas, música/SFX separados) e loading com música.

## Rodar
- `go run .` (ou `air`)
- Requer Go 1.26.

## Build (`make ...`)
- `build-win64`, `build-win32`, `web`, `android` (pede gomobile
  + SDK/NDK), `pack` (UPX), `size`, `clean`
- Saída em `tmp/`.

## Controles
- Mouse nos botões, `L` troca idioma, `ESC`/`X` fecha modal.
