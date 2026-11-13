# Duckling — Workspace Rules

Always enforce these core constraints when working in this repository:

- **File Length Limit**: Every `.go` file must ideally be **≤75 lines**, with a strict maximum of **100 lines**. Split files by responsibility immediately if a change would exceed 100 lines.
- **Import Direction Law**:
  - `engine/*` must **never** import `game/*` (sole exception: `engine/scene` registers screens).
  - `content` (`package assets`) is a leaf package and imports nothing from the project.
  - `game/*` never imports `main`, and screens (`game/screens/*`) never import each other.
  - Do not rename packages to match folder names (`engine/rendering` is `package ui`, `game/player` is `package actor`, `game/map` is `package world`, `content` is `package assets`).
- **Mandatory i18n**: All user-facing text in `Draw` must use `i18n.T("domain.key")` with entries in `engine/i18n/locales/en.json`, `es.json`, and `pt.json`.
- **Pre-Commit & GitHub Workflow**:
  - Update `CHANGELOG.md` under `## [Unreleased]` before every commit.
  - Run `go build ./...` and `go vet ./...` before committing.
  - Never commit directly to `idle`; always use a dedicated branch (`feat/...`, `fix/...`, `docs/...`), open a Pull Request, and merge via the `github-workflow` skill.
