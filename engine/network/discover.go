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
	Addr string
	Room string
	N    int
	Max  int
	Priv bool
}

// Announce grita presença até fechar stop.
func Announce(room string, n, max int, priv bool, stop <-chan struct{}) error {
	return AnnounceTo("255.255.255.255", room, n, max, priv, stop)
}

// AnnounceTo é o Announce com destino parametrizado (teste usa 127.0.0.1).
func AnnounceTo(dst, room string, n, max int, priv bool, stop <-chan struct{}) error {
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
			_, _ = c.Write(Encode(Msg{Kind: "hello", Room: room, N: n, Max: max, Pass: pass}))
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
		seen[addr.String()] = Host{Addr: addr.String(), Room: m.Room, N: m.N, Max: m.Max, Priv: m.Pass != ""}
	}
	out := make([]Host, 0, len(seen))
	for _, h := range seen {
		out = append(out, h)
	}
	return out, nil
}
