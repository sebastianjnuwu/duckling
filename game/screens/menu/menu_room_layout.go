package menu

// menu_room_layout.go — grade do modal: 1 lugar pra mudar.
// Fileiras de 62px a partir de 175; rótulo à esquerda,
// controle à direita; rodapé em 560.

const (
	rpX, rpY, rpW, rpH = 280.0, 100.0, 720.0, 540.0
	rlX                = 340.0
	rcX                = 640.0
	rTop               = 175.0
	rGap               = 62.0
	rBtnY              = 560.0
)

// rowY devolve o topo da fileira i (0 = nome).
func rowY(i int) float64 { return rTop + float64(i)*rGap }
