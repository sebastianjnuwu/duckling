---
name: new-screen
description: >-
  Use this skill when creating a new UI screen or modal in Duckling, wiring screen navigation
  in the scene router, or adding localized UI text keys in en/es/pt.
---

# Creating a New Screen & i18n

Follow this runbook when adding a new screen to `game/screens/` or adding user-facing text.

## 1. Architecture & Import Rules

- Each screen lives in its own folder: `game/screens/<name>/`.
- Screens **never** import `main` or other screens (`game/screens/*`).
- Screens communicate exclusively via navigation callbacks (`next func(string)`) passed by `engine/scene/`.
- `engine/scene/` is the **only** package allowed to import `game/screens/*`.
- Keep every file ≤75 lines (hard cap: 100 lines). Split large screens by responsibility (e.g., `<name>.go`, `<name>_draw.go`, `<name>_layout.go`).

## 2. Screen Implementation Pattern

Create `game/screens/<name>/<name>.go` with `Update() error` and `Draw(screen *ebiten.Image)`:

```go
package <name>

import "github.com/hajimehoshi/ebiten/v2"

type Screen struct {
	next func(string)
}

func New(next func(string)) *Screen {
	return &Screen{next: next}
}

func (s *Screen) Update() error {
	return nil
}

func (s *Screen) Draw(dst *ebiten.Image) {
}
```

> **Lazy Asset Loading**: Assets are loaded progressively during the `loading` screen. Do **not** fetch images or audio inside `New(...)` if `New` is called during `scene.UI()` bootstrap—load them lazily in `Update`/`Draw` or on first frame.

## 3. Register in `engine/scene/`

Wire the new screen inside `engine/scene/scene.go` within `UI()`:

```go
g.Add("<name>", <name>.New(g.Go))
```

## 4. Mandatory i18n (`en`, `es`, `pt`)

Every visible string in `Draw` must use `i18n.T("domain.key")`. Hardcoded UI strings are forbidden.

1. Add the key `<domain>.<key>` to all three embedded locale files:
   - `engine/i18n/locales/en.json`
   - `engine/i18n/locales/es.json`
   - `engine/i18n/locales/pt.json`
2. Measure text width when fitting inside fixed panels (`ui.Panel`) so Spanish and Portuguese strings do not overflow.

## 5. Verification

Run build and static checks:
```powershell
go build ./...
go vet ./...
```
