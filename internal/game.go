package internal

import (
	"math/rand"
	"strings"
)

type Game struct {
	solution []string // Puzzle solution, including blank
	tiles    []int    // Current positions of puzzle. indexes of the solution
	width    int
	height   int
	length   int    // Length of arrays - width*height
	gapVal   string // String to replace the gap position with
	gap      int    // Position of gap
}

func NewGame(solution []string, width int, height int, gapVal string) Game {
	length := width * height
	tiles := make([]int, length)

	for i := range length {
		tiles[i] = i
	}

	return Game{
		solution: solution,
		width:    width,
		height:   height,
		tiles:    tiles,
		length:   length,
		gap:      0,
	}
}

func (g *Game) Randomise() {
	g.tiles = rand.Perm(g.length)
	g.gap = rand.Intn(g.length)
}

func (g *Game) Tile(i int) string {
	if i == g.gap {
		return g.gapVal
	}
	return g.solution[i]
}

func (g *Game) Board() []string {
	board := make([]string, g.length)

	for i, tile := range g.tiles {
		board[i] = g.Tile(tile)
	}

	return board
}

func (g *Game) String() string {
	builder := strings.Builder{}

	for i, tile := range g.tiles {
		_, err := builder.WriteString(g.Tile(tile))
		if err != nil {
			panic(err)
		}
		if i%g.width == g.width-1 {
			_, err := builder.WriteRune('\n')
			if err != nil {
				panic(err)
			}
		}

	}

	return builder.String()

}
