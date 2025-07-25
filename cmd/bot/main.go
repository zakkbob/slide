package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/slack-go/slack"
)

func main() {
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

	client := slack.New(*apikey, slack.OptionDebug(true))
	a, b, err := client.PostMessage(*channelID, slack.MsgOptionText("test", false))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	upBtnText := slack.NewTextBlockObject("plain_text", ":upvote:", true, false)
	upBtn := slack.NewButtonBlockElement("", "click_me", upBtnText)
	leftBtnText := slack.NewTextBlockObject("plain_text", ":leftvote:", true, false)
	leftBtn := slack.NewButtonBlockElement("", "click_me", leftBtnText)
	rightBtnText := slack.NewTextBlockObject("plain_text", ":rightvote:", true, false)
	rightBtn := slack.NewButtonBlockElement("", "click_me", rightBtnText)
	downBtnText := slack.NewTextBlockObject("plain_text", ":downvote-red:", true, false)
	downBtn := slack.NewButtonBlockElement("", "click_me", downBtnText)

	arrowBlock := slack.NewActionBlock("", leftBtn, downBtn, upBtn, rightBtn)

	msg := slack.NewBlockMessage(
		arrowBlock,
	)

	a, b, err = client.PostMessage(*channelID, slack.MsgOptionBlocks(msg.Blocks.BlockSet...))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Made a post!", "a", a, "b", b)

}
