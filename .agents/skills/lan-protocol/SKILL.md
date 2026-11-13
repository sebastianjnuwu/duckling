---
name: lan-protocol
description: >-
  Use this skill when adding or modifying UDP network messages, snapshot payloads, or authoritative
  host/client synchronization across `engine/network` and `game/session`.
---

# LAN UDP Protocol & Authoritative Session

Follow this runbook when changing the LAN multiplayer protocol between `engine/network/` and `game/session/`.

## 1. Architecture & Import Boundaries

- `engine/network/` defines the wire format (`msg.go`), LAN broadcast discovery (`discover.go`, port `48217`), and UDP game endpoint (`endpoint.go`, port `48218`).
- `engine/network/` **never** imports `game/*`.
- `game/session/` implements the authoritative host (simulates all players, broadcasts snapshots at 15 Hz) and client (sends inputs at 30 Hz, interpolates remote ducks).
- Keep every `.go` file ≤75 lines (hard cap: 100 lines). Split `session` logic by domain (`session_host.go`, `session_client.go`, `session_meet.go`, `session_color.go`, etc.).

## 2. Adding or Changing a Message / Field

1. **Bump Protocol Version**: If the wire format or snapshot contract changes incompatibly, increment the protocol version constant in `engine/network/msg.go` so mismatched clients are cleanly rejected on `join`.
2. **Define Wire Structs**: Update `Msg` or payload structs in `engine/network/msg.go` using compact JSON tags.
3. **Authoritative Host Validation**:
   - The host (`game/session/`) is the single source of truth.
   - Never trust client claims blindly: validate phase (`Lobby`, `Play`, `Meeting`, `End`), player alive/ghost status, role (`Crew` vs `Impostor`), cooldowns, and proximity before applying actions (`kill`, `task`, `report`, `call`, `vote`, `sabotage`).
4. **UDP Reliability**:
   - UDP packets can drop. Critical edge-triggered client events (`input` jump/fly edges, `task`, `kill`, `phase` transitions) should either be sent redundantly (e.g., 2–3× with host deduplication) or mirrored in the periodic host `state` snapshot so clients recover automatically.
5. **Sensitive Data Filtering**:
   - Never broadcast impostor roles or teammate lists to all players during `Play`. Send `role` assignments individually per recipient, and only reveal full roles in `End` (or `Finish()`).

## 3. Verification

1. Run `go build ./... && go vet ./...`.
2. Write a headless loopback test in `tmp/` (host + 2 clients over UDP) to verify the full round-trip before removing the test file.
