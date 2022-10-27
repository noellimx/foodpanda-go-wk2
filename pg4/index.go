package pg4

import "fmt"

func Games() {
	fmt.Printf("[Games]\n")
	g1 := newGame("Minecraft", 5)
	g2 := newGame("World of Warcraft", 19)
	g3 := newGame("Elite Dangerous", 54)

	gs := []*game{g1, g2, g3}

	for idx, g := range gs {
		fmt.Printf("%-10d %s\n", idx+1, g.AsString())
	}
}
