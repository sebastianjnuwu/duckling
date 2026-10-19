package session

import (
	"duckling/engine/network"
	"duckling/game/match"
)

// session_chat.go — frases prontas da reunião: o host só aceita as
// chaves da lista e de quem está vivo (fantasma é mudo).

// SayKeys são as frases permitidas (i18n: meet.say.N).
var SayKeys = []string{
	"meet.say.1", "meet.say.2", "meet.say.3", "meet.say.4",
	"meet.say.5", "meet.say.6", "meet.say.7", "meet.say.8",
	"meet.say.9", "meet.say.10", "meet.say.11", "meet.say.12",
}

// SayKeyValid diz se a frase está no catálogo.
func SayKeyValid(k string) bool {
	for _, s := range SayKeys {
		if s == k {
			return true
		}
	}
	return false
}

// Say registra a frase de by (host, reunião, vivo).
func (s *Session) Say(by int, key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.phase != match.Meeting || s.ghosts[by] || !SayKeyValid(key) {
		return false
	}
	s.chats = append(s.chats, network.Chat{By: by, Key: key})
	if len(s.chats) > chatKeep {
		s.chats = s.chats[len(s.chats)-chatKeep:]
	}
	return true
}

// Chats devolve as falas (host: direto; cliente: do snapshot).
func (s *Session) Chats() []network.Chat {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]network.Chat(nil), s.chats...)
}

// SendSay manda a frase pro host (cliente).
func (s *Session) SendSay(key string) {
	if !SayKeyValid(key) {
		return
	}
	s.sendThrice(network.Msg{Kind: "say", Say: key})
}
