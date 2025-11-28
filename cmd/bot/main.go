package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/slack-go/slack"
	"github.com/zakkbob/slide/pkg"
)

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

	app := pkg.Application{
		Debug:  *debug,
		Client: client,
		Logger: logger,
	}

	http.HandleFunc("POST /action", app.HandleAction())
	http.HandleFunc("POST /slash", app.HandleSlash())
	http.ListenAndServe(":4300", nil)
}
