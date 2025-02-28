package pokedex

type Pokedex struct {
	contents map[string]Pokemon
}

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Types          []string
}

func NewPokedex() Pokedex {
	p := Pokedex{
		contents: make(map[string]Pokemon),
	}
	return p
}

func (p *Pokedex) Add(pokemon Pokemon) {
	p.contents[pokemon.Name] = pokemon
}

func (p *Pokedex) Has(name string) bool {
	_, ok := p.contents[name]
	return ok
}
