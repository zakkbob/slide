package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func (s *slackGame) msgOption() slack.MsgOption {
	gameText := slack.NewTextBlockObject("plain_text", s.game.String(), true, false)
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

	return slack.MsgOptionBlocks(msg.Blocks.BlockSet...)

}

func (s *slackGame) startGame() error {
	_, timestamp, err := s.client.PostMessage(s.channelID, s.msgOption())
	if err != nil {
		return fmt.Errorf("failed to start game: %v", err)
	}

	s.logger.Info("Made a post!", "timestamp", timestamp)
	s.timestamp = timestamp
	return nil
}

func (s *slackGame) updateGame() error {
	a, b, c, err := s.client.UpdateMessage(s.channelID, s.timestamp, s.msgOption())
	if err != nil {
		return fmt.Errorf("failed to update game: %v", err)
	}

	s.logger.Info("Updated!", "a", a, "b", b, "c", c)
	return nil

}
