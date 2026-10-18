package network

import (
	"encoding/json"
)

// msg.go — mensagens UDP (JSON pequeno, debugável no olho).
// Mesmo formato serve pra LAN hoje e pro servidor externo depois.

const (
	// Port é a porta de descoberta (announce/listen da LAN).
	Port = 48217
	// GamePort é a porta da partida (join/input/state).
	GamePort = 48218
	// Version do protocolo: cliente de versão diferente não entra.
	Version = 1
)

// Player é um pato num snapshot (posição, altitude, pose).
type Player struct {
	ID   int     `json:"id"`
	Name string  `json:"name,omitempty"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Z    float64 `json:"z,omitempty"`
	Dir  int     `json:"d,omitempty"`
	Fly  bool    `json:"f,omitempty"`
}

// Msg cobre os dois papéis: hello é a descoberta; join/deny/
// welcome/input/state/leave são a partida.
type Msg struct {
	Ver     int      `json:"v,omitempty"`
	Kind    string   `json:"kind"`
	Room    string   `json:"room,omitempty"`
	Name    string   `json:"name,omitempty"`
	Pass    string   `json:"pass,omitempty"`
	Reason  string   `json:"reason,omitempty"`
	ID      int      `json:"id,omitempty"`
	Max     int      `json:"max,omitempty"`
	N       int      `json:"n,omitempty"`
	In      uint8    `json:"in,omitempty"`
	Players []Player `json:"p,omitempty"`
}

// Encode serializa pra mandar no UDP.
func Encode(m Msg) []byte {
	b, _ := json.Marshal(m)
	return b
}

// Decode lê do UDP; lixo retorna erro (ignorar e seguir).
func Decode(b []byte) (Msg, error) {
	var m Msg
	if err := json.Unmarshal(b, &m); err != nil {
		return Msg{}, err
	}
	if m.Kind == "" {
		return Msg{}, errKind
	}
	return m, nil
}

var errKind = errString("network: sem kind")

type errString string

func (e errString) Error() string { return string(e) }
