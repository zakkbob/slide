package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

type application struct {
	debug  bool
	client *slack.Client
	logger slog.Logger
}

func main() {
	debug := flag.Bool("debug", false, "Debug Mode")
	apikey := flag.String("apikey", "", "API Key")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if *apikey == "" {
		logger.Error("You need to specify an API Key using --apikey")
		os.Exit(1)
	}

	client := slack.New(*apikey, slack.OptionDebug(*debug))

	app := application{
		debug:  *debug,
		logger: *logger,
		client: client,
	}

	// solution := []string{
	// 	":one:", ":two:", ":three:", ":four:", ":five:", ":six:",
	// }

	// game := internal.NewGame(solution, 3, 2, ":blank:")

	// logger.Info(game.String())

	// game.Randomise()

	// err := app.startGame("C097GSY3X5G", game)

	// if err != nil {
	// 	logger.Error(err.Error())
	// 	os.Exit(1)
	// }

	http.HandleFunc("POST /action", app.handleAction())
	http.HandleFunc("POST /slash", app.handleSlash())
	http.ListenAndServe(":4300", nil)
}
