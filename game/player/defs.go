package actor

// defs.go — elenco: cada personagem é 1 linha de dados.
// Novo personagem = 1 entrada aqui + png em assets/sprites.
// Cada ação/direção é 1 tira horizontal (down/up/left/right,
// fly_down/up/left/right, jump_down/up/left/right).

type Def struct {
	Prefix string
	Scale  float64
	Speed  float64
	Height float64
}

var Cast = map[string]Def{
	"duck": {Prefix: "sprites/duck/yellow/128px", Scale: 0.25, Speed: 2.6, Height: 24},
}
