package internal

import (
	"math"
	"math/rand"
	"regexp"
	"slices"
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

func GameFromString(s string) Game {
	height := strings.Count(s, "\n")

	regex := regexp.MustCompile(":.*?:")
	solution := regex.FindAllString(s, -1)

	length := len(solution)

	// Okay, so turns out i forgot about the solution thing. So this is a hacky fix for now. Might need a database :(
	tiles := make([]int, length)

	for i := range length {
		tiles[i] = i
	}

	return Game{
		solution: solution,
		tiles:    tiles,
		height:   height,
		length:   length,
		width:    length / height,
		gapVal:   ":blank:",
		gap:      slices.Index(solution, ":blank:"),
	}
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
		gapVal:   gapVal,
	}
}

func (g *Game) gapX() int {
	return g.gap % g.width
}

func (g *Game) gapY() int {
	return int(math.Floor(float64(g.gap) / float64(g.width)))
}

// Moves the gap DOWN, swapping with whatever is there
// Opposite directions will feel more natural
func (g *Game) Up() {
	if g.gapY() >= g.height-1 {
		return
	}

	newGap := g.gap + g.width

	g.tiles[g.gap], g.tiles[newGap] = g.tiles[newGap], g.tiles[g.gap]
	g.gap = newGap
}

// Moves the gap UP, swapping with whatever is there
// Opposite directions will feel more natural
func (g *Game) Down() {
	if g.gapY() <= 0 {
		return
	}

	newGap := g.gap - g.width

	g.tiles[g.gap], g.tiles[newGap] = g.tiles[newGap], g.tiles[g.gap]
	g.gap = newGap
}

// Moves the gap RIGHT, swapping with whatever is there
// Opposite directions will feel more natural
func (g *Game) Left() {
	if g.gapX() >= g.width-1 {
		return
	}

	newGap := g.gap + 1

	g.tiles[g.gap], g.tiles[newGap] = g.tiles[newGap], g.tiles[g.gap]
	g.gap = newGap
}

// Moves the gap LEFT, swapping with whatever is there
// Opposite directions will feel more natural
func (g *Game) Right() {
	if g.gapX() <= 0 {
		return
	}

	newGap := g.gap - 1

	g.tiles[g.gap], g.tiles[newGap] = g.tiles[newGap], g.tiles[g.gap]
	g.gap = newGap
}

func (g *Game) Gap() int {
	return g.gap
}

func CountInversions(a []int) int {
	count := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i+1] < a[i] {
			count++
		}
	}
	return count
}

func (g *Game) Randomise() {
	g.tiles = rand.Perm(g.length)

	// Some permutations aren't solvable!
	// This fix only works for 3x2 board.
	// TODO: make work for nxk boards
	for CountInversions(g.tiles)%2 == 1 {
		g.tiles = rand.Perm(g.length)
	}

	g.gap = g.length - 1
}

func (g *Game) Tile(i int) string {
	if i == g.gap {
		return g.gapVal
	}
	return g.solution[g.tiles[i]]
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

	for i := range g.length {
		_, err := builder.WriteString(g.Tile(i))
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
