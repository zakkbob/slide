package main

import (
	"os"

	"github.com/slack-go/slack"
)

func (a *application) startGame() {
	gameText := slack.NewTextBlockObject("plain_text", a.game.String(), true, false)
	gameSection := slack.NewSectionBlock(gameText, nil, nil)

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
		gameSection,
		arrowBlock,
	)

	_, b, err := a.client.PostMessage(a.channelID, slack.MsgOptionBlocks(msg.Blocks.BlockSet...))
	if err != nil {
		a.logger.Error(err.Error())
		os.Exit(1)
	}

	a.logger.Info("Made a post!", "timestamp", b)

}
