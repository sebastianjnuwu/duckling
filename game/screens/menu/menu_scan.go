package menu

import (
	"time"

	"duckling/engine/i18n"
	"duckling/engine/network"
	"duckling/game/session"
)

// menu_scan.go — varredura da LAN: escuta numa goroutine e
// guarda as salas achadas; a UI lê a lista pronta (sem travar).

// perPage é quantas salas cabem por página na tabelinha.
const perPage = 4

func (l *RoomState) pages() int {
	return max(1, (len(l.roomList())+perPage-1)/perPage)
}

// startHost cria a sala na LAN e entra no jogo como dono.
func (l *RoomState) startHost() {
	pass := ""
	if l.priv {
		pass = l.pass
	}
	s, err := session.Host(l.name, l.name, pass, l.min, l.max, l.imp)
	if err != nil {
		l.toast.Show(i18n.T("net.fail"))
		return
	}
	l.scanStop()
	session.Set(s)
	l.next("play")
}

func (l *RoomState) scanStart() {
	if l.scan != nil {
		return
	}
	l.scan = make(chan struct{})
	go l.scanLoop(l.scan)
}

func (l *RoomState) scanStop() {
	if l.scan != nil {
		close(l.scan)
		l.scan = nil
	}
}

func (l *RoomState) scanLoop(stop chan struct{}) {
	for {
		select {
		case <-stop:
			return
		default:
		}
		hs, err := network.Listen(700 * time.Millisecond)
		if err != nil {
			time.Sleep(400 * time.Millisecond)
			continue
		}
		l.setRooms(hs)
	}
}

func (l *RoomState) setRooms(hs []network.Host) {
	out := make([]room, 0, len(hs))
	for _, h := range hs {
		out = append(out, room{n: h.N, max: h.Max, priv: h.Priv, ip: h.IP, name: h.Room})
	}
	l.mu.Lock()
	l.rooms = out
	l.mu.Unlock()
}

// roomList devolve uma cópia da lista (segura entre goroutines).
func (l *RoomState) roomList() []room {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]room(nil), l.rooms...)
}
