package ui

// button_push.go — fechar modal com animação: arma no clique
// e confirma quando o afundar aparece (sem matar o push).

func (b *Btn) Push() bool {
	if b.Update() {
		b.armed = true
	}
	if b.armed && b.Press > 0 && b.Press < 0.7 {
		b.armed = false
		return true
	}
	return false
}
