package actor

// defs.go — elenco: cada personagem é 1 linha de dados.
// Novo personagem = 1 entrada aqui + png em assets/sprites.
// Fileiras da sheet do pato: lado 0/1/3, frente 2/4.

type Def struct {
	Side, Down, Up, Fly int
	Cell, Cols          int
	Scale               float64
	Speed               float64
}

var Cast = map[string]Def{
	"duck": {Side: 4, Down: 4, Up: 2, Fly: 0, Cell: 128, Cols: 24, Scale: 0.25, Speed: 3},
}
