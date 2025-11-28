package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/slack-go/slack"
	"github.com/zakkbob/slide/internal"
)

func SlashHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Only POST allowed")
		return
	}

	debug := false
	apikey := os.Getenv("SLACK_API_KEY")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if apikey == "" {
		logger.Error("Received slash command, but SLACK_API_KEY is not set")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Info("Received slash command", "host", r.Host)

	app := internal.Application{
		Debug:  debug,
		Client: slack.New(apikey, slack.OptionDebug(debug)),
		Logger: logger,
	}

	app.HandleSlash()(w, r)
}
