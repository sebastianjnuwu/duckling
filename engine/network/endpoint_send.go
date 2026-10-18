package network

import "net"

// endpoint_send.go — saída do endpoint: o cliente fala só com o
// host; o host escolhe o convidado pelo endereço.

// SendTo manda bytes pra um endereço. No cliente (conectado) só
// existe o host; no host, o endereço escolhe o convidado.
func (e *Endpoint) SendTo(addr string, b []byte) error {
	if e.host != "" {
		_, err := e.conn.Write(b)
		return err
	}
	if addr == "" {
		return nil
	}
	ua, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	_, err = e.conn.WriteToUDP(b, ua)
	return err
}

// Send manda pro host fixo (cliente → host).
func (e *Endpoint) Send(b []byte) error { return e.SendTo(e.host, b) }
