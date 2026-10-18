package network

import (
	"net"
	"strconv"
	"time"
)

// endpoint.go — socket UDP cru da partida: o host escuta no
// GamePort e o cliente disca pro host. A leitura roda numa
// goroutine e entrega pacotes num canal (nunca trava o jogo).

// Packet é um datagrama recebido + quem mandou (ip:porta).
type Packet struct {
	From string
	Data []byte
}

type Endpoint struct {
	conn *net.UDPConn
	host string // cliente: destino fixo; host: ""
	in   chan Packet
	done chan struct{}
}

// ListenHost abre a partida no GamePort (papel de host).
func ListenHost() (*Endpoint, error) {
	ua, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(GamePort))
	if err != nil {
		return nil, err
	}
	c, err := net.ListenUDP("udp", ua)
	if err != nil {
		return nil, err
	}
	return newEndpoint(c, ""), nil
}

// DialHost disca pro host da LAN (papel de cliente).
func DialHost(addr string) (*Endpoint, error) {
	ua, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	c, err := net.DialUDP("udp", nil, ua)
	if err != nil {
		return nil, err
	}
	return newEndpoint(c, addr), nil
}

func newEndpoint(c *net.UDPConn, host string) *Endpoint {
	e := &Endpoint{conn: c, host: host, in: make(chan Packet, 256), done: make(chan struct{})}
	go e.read()
	return e
}

// C é o canal de pacotes recebidos.
func (e *Endpoint) C() <-chan Packet { return e.in }

func (e *Endpoint) read() {
	buf := make([]byte, 4096)
	for {
		_ = e.conn.SetReadDeadline(time.Now().Add(time.Second))
		n, addr, err := e.conn.ReadFromUDP(buf)
		if err != nil {
			select {
			case <-e.done:
				return
			default:
				continue // timeout ou pacote ruim: tenta de novo
			}
		}
		pkt := Packet{From: addr.String(), Data: append([]byte(nil), buf[:n]...)}
		select {
		case e.in <- pkt:
		default: // fila cheia: joga fora o mais novo
		}
	}
}

// Close encerra a leitura e o socket.
func (e *Endpoint) Close() {
	select {
	case <-e.done:
	default:
		close(e.done)
	}
	_ = e.conn.Close()
}
