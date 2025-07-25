package internal_test

import (
	"testing"

	"github.com/zakkbob/slide/internal"
)

func TestGameNew(t *testing.T) {
	solution := []string{
		"tile0.0", "tile0.1", "tile0.2", "tile0.3",
		"tile1.0", "tile1.1", "tile1.2", "tile1.3",
		"tile2.0", "tile2.1", "tile2.2", "tile2.3",
	}

	game := internal.NewGame(solution, 4, 3)
	t.Log(game.String())

	game.Randomise()
	t.Log(game.String())

}
