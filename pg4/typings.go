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

func Games() {
	fmt.Printf("[Games]\n")
	g1 := newGame("Minecraft", 5)
	g2 := newGame("World of Warcraft", 19)
	g3 := newGame("Elite Dangerous", 54)

	gs := []*game{g1, g2, g3}

	for idx, g := range gs {
		fmt.Printf("%-10d %s\n", idx, g.AsString())
	}

}
