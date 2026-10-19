package play

// play_meet_geo.go — geometria da reunião: painel, lista (até 8),
// log de falas e os 12 botões de frase (2 colunas).

const (
	mtX, mtY, mtW, mtH                 = 140.0, 60.0, 1000.0, 600.0
	mtRowX, mtRowY, mtRowW, mtRowH     = 170.0, 176.0, 420.0, 42.0
	mtRowGap                           = 4.0
	mtLogX, mtLogY, mtLogStep          = 620.0, 176.0, 28.0
	mtSayX, mtSayY                     = 620.0, 424.0
	mtSayW, mtSayH, mtSayGX, mtSayGY   = 246.0, 32.0, 8.0, 6.0
	mtSkipX, mtSkipY, mtSkipW, mtSkipH = 170.0, 562.0, 200.0, 42.0
	mtRows, mtLogLines                 = 8, 8
)

// meetRowRect devolve o retângulo da linha i da lista (1 ou 2 colunas).
func meetRowRect(i, n int) (x, y, w, h float64) {
	if n <= mtRows {
		return mtRowX, mtRowY + float64(i)*(mtRowH+mtRowGap), mtRowW, mtRowH
	}
	col, row := i/mtRows, i%mtRows
	return mtRowX + float64(col)*(mtRowW/2+8),
		mtRowY + float64(row)*(mtRowH+mtRowGap), mtRowW/2 - 8, mtRowH
}

// meetSayRect devolve o retângulo da frase i (2 colunas).
func meetSayRect(i int) (x, y, w, h float64) {
	col, row := i%2, i/2
	return mtSayX + float64(col)*(mtSayW+mtSayGX),
		mtSayY + float64(row)*(mtSayH+mtSayGY), mtSayW, mtSayH
}

// inRect diz se o ponto caiu no retângulo.
func inRect(x, y, rx, ry, rw, rh float64) bool {
	return x >= rx && x <= rx+rw && y >= ry && y <= ry+rh
}
