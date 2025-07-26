package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log/slog"

	"github.com/slack-go/slack"
)

type Application struct {
	debug  bool
	client *slack.Client
	logger *slog.Logger
}

func NewApplication(debug bool, client *slack.Client, logger *slog.Logger) Application {
	return Application {
		debug: debug,
		client: client,
		logger: logger,
	}
}

func (a *Application) gameString(i slack.InteractionCallback) string {
	for _, block := range i.Message.Blocks.BlockSet {
		if block.ID() == "game" && block.BlockType() == slack.MBTSection {
			sectionBlock, ok := block.(*slack.SectionBlock)
			if ok {
				return sectionBlock.Text.Text
			}
			panic("not okay!!")
		}
	}
	panic("aghhh")
}

func (a *Application) HandleSlash() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := slack.SlashCommandParse(r)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch s.Command {
		case "/slide-test":
			solution := []string{
				":one:", ":two:", ":three:", ":four:", ":five:", ":six:",
			}
			game := NewGame(solution, 3, 2, ":blank:")
			game.Randomise()

			a.startGame(s.ChannelID, game)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a *Application) HandleAction() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var i slack.InteractionCallback

		err := json.Unmarshal([]byte(r.FormValue("payload")), &i)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		timestamp := i.Container.MessageTs
		channelID := i.Container.ChannelID

		for _, actionStr := range i.ActionCallback.BlockActions {
			json.Unmarshal([]byte(actionStr.ActionID), &a)

			gameStr := a.gameString(i)

			game := GameFromString(gameStr)

			switch actionStr.ActionID {
			case "left":
				game.Left()
			case "right":
				game.Right()
			case "up":
				game.Up()
			case "down":
				game.Down()
			}

			err := a.updateGame(channelID, timestamp, game)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}

func (a *Application) msgOption(game string) slack.MsgOption {
	gameText := slack.NewTextBlockObject("plain_text", game, true, false)
	gameSection := slack.NewSectionBlock(gameText, nil, nil, slack.SectionBlockOptionBlockID("game"))

	upBtnText := slack.NewTextBlockObject("plain_text", ":upvote:", true, false)
	upBtn := slack.NewButtonBlockElement("up", "click_me", upBtnText)
	leftBtnText := slack.NewTextBlockObject("plain_text", ":leftvote:", true, false)
	leftBtn := slack.NewButtonBlockElement("left", "click_me", leftBtnText)
	rightBtnText := slack.NewTextBlockObject("plain_text", ":rightvote:", true, false)
	rightBtn := slack.NewButtonBlockElement("right", "click_me", rightBtnText)
	downBtnText := slack.NewTextBlockObject("plain_text", ":downvote-red:", true, false)
	downBtn := slack.NewButtonBlockElement("down", "click_me", downBtnText)

	arrowBlock := slack.NewActionBlock("", leftBtn, downBtn, upBtn, rightBtn)

	msg := slack.NewBlockMessage(
		gameSection,
		arrowBlock,
	)

	return slack.MsgOptionBlocks(msg.Blocks.BlockSet...)
}

func (a *Application) startGame(channelID string, game Game) error {
	_, timestamp, err := a.client.PostMessage(channelID, a.msgOption(game.String()))
	if err != nil {
		return fmt.Errorf("failed to start game: %v", err)
	}

	a.logger.Info("Made a post!", "timestamp", timestamp)
	return nil
}

func (a *Application) updateGame(channelID string, timestamp string, game Game) error {
	_, _, _, err := a.client.UpdateMessage(channelID, timestamp, a.msgOption(game.String()))
	if err != nil {
		return fmt.Errorf("failed to update game: %v", err)
	}

	return nil
}
