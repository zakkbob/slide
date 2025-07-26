package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/slide/pkg"
)

func TestNewGame(t *testing.T) {
	solution := []string{
		"1", "2", "3", "4", "5", "6",
	}

	expected := "-23\n456\n"

	game := pkg.NewGame(solution, 3, 2, "-")

	assert.Equal(t, expected, game.String())
}

func TestGameUpDown(t *testing.T) {
	solution := []string{
		"1", "2", "3", "4", "5", "6",
	}

	downExpected := "-23\n456\n"
	upExpected := "423\n-56\n"

	game := pkg.NewGame(solution, 3, 2, "-")

	game.Up()
	assert.Equal(t, upExpected, game.String())

	game.Up()
	assert.Equal(t, upExpected, game.String())

	game.Down()
	assert.Equal(t, downExpected, game.String())

	game.Down()
	assert.Equal(t, downExpected, game.String())
}

func TestGameLeftRight(t *testing.T) {
	solution := []string{
		"1", "2", "3", "4", "5", "6",
	}

	rightExpected := "-2\n34\n56\n"
	leftExpected := "2-\n34\n56\n"

	game := pkg.NewGame(solution, 2, 3, "-")

	game.Left()
	assert.Equal(t, leftExpected, game.String(), "Left() should move the gap to the right")

	game.Left()
	assert.Equal(t, leftExpected, game.String(), "Left() should have no effect when the gap is on the right edge")

	game.Right()
	assert.Equal(t, rightExpected, game.String(), "Right() should move the gap to the left")

	game.Right()
	assert.Equal(t, rightExpected, game.String(), "Right() should have no effect when the gap is on the left edge")
}
