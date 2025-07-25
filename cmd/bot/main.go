package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/slack-go/slack"
	"github.com/zakkbob/slide/internal"
)

type slackGame struct {
	debug     bool
	channelID string
	timestamp string
	client    *slack.Client
	logger    slog.Logger
	game      internal.Game
}

func main() {
	debug := flag.Bool("debug", false, "Debug Mode")
	channelID := flag.String("channel", "", "Channel ID")
	apikey := flag.String("apikey", "", "API Key")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if *channelID == "" {
		logger.Error("You need to specify a Channel ID using --channel")
		os.Exit(1)
	}

	if *apikey == "" {
		logger.Error("You need to specify an API Key using --apikey")
		os.Exit(1)
	}

	client := slack.New(*apikey, slack.OptionDebug(*debug))

	solution := []string{
		":one:", ":two:", ":three:", ":four:", ":five:", ":six:",
	}

	game := internal.NewGame(solution, 3, 2, ":blank:")
	//game.Randomise()

	logger.Info(game.String())

	app := slackGame{
		debug:     *debug,
		channelID: *channelID,
		logger:    *logger,
		client:    client,
		game:      game,
	}

	app.game.Randomise()

	err := app.startGame()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	http.HandleFunc("POST /action", app.handleAction(""))
	http.ListenAndServe(":4300", nil)

}
