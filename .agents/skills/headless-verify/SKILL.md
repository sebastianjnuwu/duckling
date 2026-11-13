---
name: headless-verify
description: >-
  Use this skill when verifying gameplay mechanics, map collisions/tasks, UDP netcode, or save/load
  behavior at runtime using a standalone headless Go program in `tmp/` before committing.
---

# Headless Verification & Build Checks

Whenever runtime behavior is touched (physics, map loading, task reachability, UDP session state, or save persistence), prove it works with a fast headless verification program in `tmp/` in addition to `go build` and `go vet`.

## 1. Standard Compile & Static Analysis

Always run:
```powershell
go build ./...
go vet ./...
```

## 2. Writing a Headless Probe in `tmp/`

`tmp/` is gitignored. Never place test binaries or scratch `main.go` files in the repository root.

1. Create a temporary directory such as `tmp/verify_<topic>/main.go`.
2. Import the target packages (`game/map`, `game/player`, `game/session`, `engine/network`, `engine/save`, etc.) without starting `ebiten.RunGame`.
3. Run the probe:
   ```powershell
   go run ./tmp/verify_<topic>
   ```
4. **Safety with `engine/save`**: If testing `save.Store` or `save.Load`, always back up and restore the real `settings.json` in `os.UserConfigDir()` so user coins, hats, and keybinds are not overwritten.

## 3. Common Probe Scenarios

- **Map & Task Reachability**: Call `world.Open("spawn")`, inspect `m.Spawns()`, `m.Tasks`, `m.Solids`, and verify that task coordinates and `carry` targets (`to_x`, `to_y`) are not inside impassable solids (`H > actor.JumpPeak`).
- **UDP Session / Netcode**: Spin up a host (`session.Host`) and 1–2 clients (`session.Join`) over loopback (`127.0.0.1`), tick `Update()`, and assert phase transitions (`Lobby` -> `Play` -> `Meeting` -> `End`), role assignment, task completion, or voting.
- **Physics & Elevation**: Instantiate `actor.Duck`, feed synthetic `actor.In` inputs across frames, and assert `X`, `Y`, `Z` against elevated solids (`world.Solid{H}`).

## 4. Clean Up

Always delete temporary probe directories and binaries from `tmp/` once verification succeeds:
```powershell
Remove-Item -Recurse -Force tmp/verify_<topic>
```
