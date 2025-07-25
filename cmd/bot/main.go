package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/slack-go/slack"
)

type application struct {
	debug     bool
	channelID string
	client    *slack.Client
	logger    slog.Logger
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

	app := application{
		debug:     *debug,
		channelID: *channelID,
		logger:    *logger,
		client:    client,
	}

	app.startGame()

}
