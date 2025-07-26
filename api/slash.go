package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/slack-go/slack"
	"github.com/zakkbob/slide/pkg"
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

	logger.Info("test")

	if apikey == "" {
		w.WriteHeader(http.StatusInternalServerError)
	}

	client := slack.New(apikey, slack.OptionDebug(debug))

	app := pkg.NewApplication(debug, client, logger)

	app.HandleSlash()(w, r)
}
