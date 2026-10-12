package actor

// defs.go — elenco: cada personagem é 1 linha de dados.
// Novo personagem = 1 entrada aqui + png em assets/sprites.
// Fileiras da sheet do pato: lado 0/1/3, frente 2/4;
// o voo alterna Fly (asa baixa) com FlapUp (asa erguida).

type Def struct {
	Prefix string
	Scale  float64
	Speed  float64
	Height float64
}

var Cast = map[string]Def{
	"duck": {Prefix: "sprites/duck/yellow/128px", Scale: 0.25, Speed: 3, Height: 24},
}
