package actor

// hat.go — chapéus da loja: catálogo (nome e preço) e o que o
// pato veste. O desenho fica em hat_draw.go.

type Hat int

const (
	HatNone Hat = iota
	HatCap
	HatStraw
	HatCrown
	HatCount
)

var hatDefs = [HatCount]struct {
	name  string
	price int
}{
	{"none", 0},
	{"cap", 30},
	{"straw", 80},
	{"crown", 150},
}

// HatValid diz se h é um chapéu do catálogo.
func HatValid(h int) bool { return h >= 0 && h < int(HatCount) }

// Name devolve a chave curta (i18n: hat.<name>).
func (h Hat) Name() string {
	if !HatValid(int(h)) {
		return hatDefs[0].name
	}
	return hatDefs[h].name
}

// Price é o preço em moedas (0 = de graça).
func (h Hat) Price() int {
	if !HatValid(int(h)) {
		return 0
	}
	return hatDefs[h].price
}

// Hat é o chapéu equipado no pato.
func (d *Duck) Hat() Hat { return d.hat }

// SetHat troca o chapéu (só desenho; a posse é do save).
func (d *Duck) SetHat(h Hat) {
	if !HatValid(int(h)) {
		h = HatNone
	}
	d.hat = h
}
