package network

import (
	"encoding/json"
)

// msg.go — mensagens UDP (JSON pequeno, debugável no olho).
// Mesmo formato serve pra LAN hoje e pro servidor externo depois.

const Port = 48217

type Msg struct {
	Kind string `json:"kind"` // hello, join, deny, state
	Room string `json:"room,omitempty"`
	Name string `json:"name,omitempty"`
	Pass string `json:"pass,omitempty"`
	N    int    `json:"n,omitempty"`
	Max  int    `json:"max,omitempty"`
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
