# Duckling

Idle 2D em Go + Ebiten: um patinho, uma loja e muita
configuração. Menu animado em pixel-art, 3 idiomas e
áudio separado em música + efeitos.

- **Menu** — patão animado na grama, JOGAR / LOJA / OPÇÕES /
  SAIR com botões e ícones pixel (dourado no hover).
- **Loading** — barra de progresso com música em loop.
- **Loja** — placeholder pronto pra itens
- **Opções em páginas** — idioma (BR/US/ES em bandeiras),
  música e SFX com liga/desliga + volume de arrastar.
- **i18n** — PT/EN/ES embutido no binário.

## Rodar

Requer Go 1.26.

```sh
go run .      # roda direto
air           # com live reload (.air.toml)
```

## Build (`mingw32-make ...`)

| Alvo          | Saída                    |
|---------------|--------------------------|
| `build-win64` | `tmp/duckling-win64.exe` |
| `build-win32` | `tmp/duckling-win32.exe` |

Suporte: Windows 10 e 11, 32-bit e 64-bit. 

## Estrutura

```text
main.go              bootstrap Ebiten
src/screen/          router + uma pasta por tela
src/ui/              botões, toggle, slider, ícones, texto
src/audio/           música (loop) e sfx em canais separados
src/i18n/            PT/EN/ES embutido
assets/              fonts, images, music, sprites (embed)
```

<div align="center">
  <img src="assets/images/screen_1.png" width="80%">
</div>