package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"log/slog"

	"github.com/slack-go/slack"
)

var (
	ErrFailedToCreateGameString = errors.New("Failed to create game string")
)

type Application struct {
	Debug  bool
	Client *slack.Client
	Logger *slog.Logger
}

func (a *Application) HandleSlash() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		args := strings.Split(strings.Trim(s.Text, " "), " ")

		a.Logger.Error("Handling slash command", "command", s.Command, "text", s.Text, "args", args, "argCount", len(args))

		var game Game

		switch s.Command {
		case "/slide-test":
			width, height := 3, 2

			if len(args) != 0 {
				switch args[0] {
				case "default":
					args = args[1:]

					if !(len(args) == 0 || len(args) == 2) {
						a.Logger.Error("Received invalid number of command arguments", "command", s.Command, "args", s.Text, "count", len(args))
						w.WriteHeader(http.StatusBadRequest)
						return
					}

					if len(args) == 2 {
						width, err = strconv.Atoi(args[0])
						if err != nil {
							a.Logger.Error("Failed to parse command arguments", "command", s.Command, "args", s.Text)
							w.WriteHeader(http.StatusBadRequest)
							return
						}
						height, err = strconv.Atoi(args[1])
						if err != nil {
							a.Logger.Error("Failed to parse command arguments", "command", s.Command, "args", s.Text)
							w.WriteHeader(http.StatusBadRequest)
							return
						}
						if !(width > 0 && height > 0 && width*height <= 9) {
							a.Logger.Error("Received invalid sliding puzzle dimensions", "width", width, "height", height)
							w.WriteHeader(http.StatusBadRequest)
							return
						}
					}

					solution := []string{
						":one:", ":two:", ":three:", ":four:", ":five:", ":six:", ":seven:", ":eight:", ":nine:",
					}
					game = NewGame(solution, width, height, ":blank:")
					game.DoRandomMoves(game.length * 10)

				case "custom":
					gameString := strings.Trim(s.Text[len(args[0])+1:], " ")
					game = GameFromString(gameString)
					game.gap = game.length - 1
					game.gapVal = ":blank:"
					game.DoRandomMoves(game.length * 10)
				}
			} else {
				solution := []string{
					":one:", ":two:", ":three:", ":four:", ":five:", ":six:", ":seven:", ":eight:", ":nine:",
				}
				game = NewGame(solution, width, height, ":blank:")
				game.DoRandomMoves(game.length * 10)
			}

		default:
			a.Logger.Error("Received unknown command", "command", s.Command)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		a.startGame(s.ChannelID, game)
	}
}

func (a *Application) HandleAction() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var i slack.InteractionCallback

		err := json.Unmarshal([]byte(r.FormValue("payload")), &i)
		if err != nil {
			a.Logger.Error("Failed to handle action", "error", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		timestamp := i.Container.MessageTs
		channelID := i.Container.ChannelID

		for _, actionStr := range i.ActionCallback.BlockActions {
			json.Unmarshal([]byte(actionStr.ActionID), &a)

			gameStr, err := a.gameString(i)
			if err != nil {
				a.Logger.Error("Failed to handle action", "error", err.Error())
			}

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

			err = a.updateGame(channelID, timestamp, game)
			if err != nil {
				a.Logger.Error("Failed to update game", "error", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

func (a *Application) gameString(i slack.InteractionCallback) (string, error) {
	for _, block := range i.Message.Blocks.BlockSet {
		if block.ID() == "game" && block.BlockType() == slack.MBTSection {
			sectionBlock, ok := block.(*slack.SectionBlock)
			if ok {
				return sectionBlock.Text.Text, nil
			}
			return "", ErrFailedToCreateGameString
		}
	}
	return "", ErrFailedToCreateGameString
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
	_, timestamp, err := a.Client.PostMessage(channelID, a.msgOption(game.String()))
	if err != nil {
		a.Logger.Info("Failed to start a new game", "channel", channelID, "error", err)
		return fmt.Errorf("failed to start game: %v", err)
	}
	a.Logger.Info("Started a new game", "timestamp", timestamp, "channel", channelID)
	return nil
}

func (a *Application) updateGame(channelID string, timestamp string, game Game) error {
	_, _, _, err := a.Client.UpdateMessage(channelID, timestamp, a.msgOption(game.String()))
	if err != nil {
		return fmt.Errorf("failed to update game: %v", err)
	}
	return nil
}
