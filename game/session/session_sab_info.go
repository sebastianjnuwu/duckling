package session

// session_sab_info.go — leitura das sabotagens: cooldown e quanto
// tempo ainda resta de cada efeito.

// SabotageCooldown devolve os segundos que faltam pra sabotar de novo.
func (s *Session) SabotageCooldown() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.host {
		return s.sabCd.Seconds()
	}
	return float64(s.sabMS) / 1000
}

// LampLeft devolve os segundos restantes da lamparina apagada.
func (s *Session) LampLeft() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.host {
		return s.lampLeft.Seconds()
	}
	return float64(s.lampMS) / 1000
}

// GateLeft devolve os segundos restantes das porteiras fechadas.
func (s *Session) GateLeft() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.host {
		return s.gateLeft.Seconds()
	}
	return float64(s.gateMS) / 1000
}
