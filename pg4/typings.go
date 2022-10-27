package pg4

import "fmt"

type game struct {
	Title string
	Price float64
}

func (g *game) AsString() string {

	return fmt.Sprintf("%s, $%0.5f", g.Title, g.Price)
}

func newGame(t string, p float64) *game {
	return &game{
		Title: t,
		Price: p,
	}
}

