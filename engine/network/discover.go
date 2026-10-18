package network

import (
	"net"
	"strconv"
	"time"
)

// discover.go — presença na LAN: host anuncia em broadcast,
// quem procura escuta e lista. Sem servidor externo.

// Host é uma sala achada na rede.
type Host struct {
	Addr string // ip:porta de origem do anúncio
	IP   string // só o IP (é onde a partida conecta)
	Room string
	N    int
	Max  int
	Priv bool
}

// Announce grita presença com contagem fixa até fechar stop.
func Announce(room string, n, max int, priv bool, stop <-chan struct{}) error {
	return AnnounceLoop(room, func() int { return n }, max, priv, stop)
}

// AnnounceLoop é o Announce com contagem viva (sala que enche).
func AnnounceLoop(room string, count func() int, max int, priv bool, stop <-chan struct{}) error {
	return announceTo("255.255.255.255", room, count, max, priv, stop)
}

// AnnounceTo é o Announce com destino parametrizado (teste: 127.0.0.1).
func AnnounceTo(dst, room string, n, max int, priv bool, stop <-chan struct{}) error {
	return announceTo(dst, room, func() int { return n }, max, priv, stop)
}

func announceTo(dst, room string, count func() int, max int, priv bool, stop <-chan struct{}) error {
	addr, err := net.ResolveUDPAddr("udp", dst+":"+strconv.Itoa(Port))
	if err != nil {
		return err
	}
	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer c.Close()
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
	pass := ""
	if priv {
		pass = "1"
	}
	for {
		select {
		case <-stop:
			return nil
		case <-t.C:
			_, _ = c.Write(Encode(Msg{Ver: Version, Kind: "hello", Room: room, N: count(), Max: max, Pass: pass}))
		}
	}
}

// Listen escuta por dur e devolve os hosts distintos.
func Listen(dur time.Duration) ([]Host, error) {
	c, err := net.ListenPacket("udp", ":"+strconv.Itoa(Port))
	if err != nil {
		return nil, err
	}
	defer c.Close()
	_ = c.SetDeadline(time.Now().Add(dur))
	seen := map[string]Host{}
	buf := make([]byte, 1024)
	for {
		n, addr, err := c.ReadFrom(buf)
		if err != nil {
			break
		}
		m, err := Decode(buf[:n])
		if err != nil || m.Kind != "hello" {
			continue
		}
		ip := ""
		if ua, ok := addr.(*net.UDPAddr); ok {
			ip = ua.IP.String()
		}
		seen[addr.String()] = Host{Addr: addr.String(), IP: ip, Room: m.Room, N: m.N, Max: m.Max, Priv: m.Pass != ""}
	}
	out := make([]Host, 0, len(seen))
	for _, h := range seen {
		out = append(out, h)
	}
	return out, nil
}
