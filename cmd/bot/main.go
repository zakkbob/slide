package main

import (
	"flag"
	"log/slog"
	"math/rand"
	"os"
	"time"

	"github.com/slack-go/slack"
	"github.com/zakkbob/slide/internal"
)

type slackGame struct {
	debug     bool
	channelID string
	timestamp string
	client    *slack.Client
	logger    slog.Logger
	game      internal.Game
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

	solution := []string{":elephpant-1-3-5b2d:", ":elephpant-2-3-cc88:", ":elephpant-3-3-593d:", ":elephpant-4-3-58bf:", ":elephpant-5-3-e074:", ":elephpant-6-3-a200:", ":elephpant-7-3-a241:", ":elephpant-8-3-a377:", ":elephpant-9-3-29f6:", ":elephpant-10-3-403e:", ":elephpant-11-3-3be5:", ":elephpant-12-3-88d3:", ":elephpant-13-3-19c7:", ":elephpant-14-3-78ce:", ":elephpant-15-3-b6b6:", ":elephpant-16-3-f02f:", ":elephpant-1-4-84cf:", ":elephpant-2-4-d077:", ":elephpant-3-4-8e77:", ":elephpant-4-4-c74a:", ":elephpant-5-4-fd7d:", ":elephpant-6-4-2fdc:", ":elephpant-7-4-f448:", ":elephpant-8-4-5e99:", ":elephpant-9-4-3d85:", ":elephpant-10-4-3569:", ":elephpant-11-4-e585:", ":elephpant-12-4-fe45:", ":elephpant-13-4-fecc:", ":elephpant-14-4-f934:", ":elephpant-15-4-d361:", ":elephpant-16-4-25be:", ":elephpant-1-5-ec74:", ":elephpant-2-5-ef90:", ":elephpant-3-5-bbf8:", ":elephpant-4-5-2d29:", ":elephpant-5-5-9926:", ":elephpant-6-5-4cad:", ":elephpant-7-5-ea34:", ":elephpant-8-5-6282:", ":elephpant-9-5-7cde:", ":elephpant-10-5-cb75:", ":elephpant-11-5-ff74:", ":elephpant-12-5-15d1:", ":elephpant-13-5-fe17:", ":elephpant-14-5-f791:", ":elephpant-15-5-8d10:", ":elephpant-16-5-e7ea:", ":elephpant-1-6-5c1e:", ":elephpant-2-6-11d3:", ":elephpant-3-6-f5b1:", ":elephpant-4-6-2ed1:", ":elephpant-5-6-92d5:", ":elephpant-6-6-83de:", ":elephpant-7-6-130d:", ":elephpant-8-6-fd11:", ":elephpant-9-6-2db6:", ":elephpant-10-6-e3ff:", ":elephpant-11-6-c399:", ":elephpant-12-6-6355:", ":elephpant-13-6-507d:", ":elephpant-14-6-e973:", ":elephpant-15-6-772b:", ":elephpant-16-6-1865:", ":elephpant-1-7-9841:", ":elephpant-2-7-560a:", ":elephpant-3-7-9224:", ":elephpant-4-7-9136:", ":elephpant-5-7-35d8:", ":elephpant-6-7-3144:", ":elephpant-7-7-7b4a:", ":elephpant-8-7-5e22:", ":elephpant-9-7-62fd:", ":elephpant-10-7-dbb3:", ":elephpant-11-7-f3e2:", ":elephpant-12-7-d623:", ":elephpant-13-7-f3cc:", ":elephpant-14-7-f8d7:", ":elephpant-15-7-dd96:", ":elephpant-16-7-00c1:", ":elephpant-1-8-4776:", ":elephpant-2-8-9270:", ":elephpant-3-8-7184:", ":elephpant-4-8-a619:", ":elephpant-5-8-f317:", ":elephpant-6-8-4e4a:", ":elephpant-7-8-1054:", ":elephpant-8-8-2171:", ":elephpant-9-8-9279:", ":elephpant-10-8-548e:", ":elephpant-11-8-39f8:", ":elephpant-12-8-0c8a:", ":elephpant-13-8-c702:", ":elephpant-14-8-98c5:", ":elephpant-15-8-db0c:", ":elephpant-16-8-9a5c:", ":elephpant-1-9-b236:", ":elephpant-2-9-7c93:", ":elephpant-3-9-83b8:", ":elephpant-4-9-7dff:", ":elephpant-5-9-3043:", ":elephpant-6-9-6642:", ":elephpant-7-9-200f:", ":elephpant-8-9-2453:", ":elephpant-9-9-cfe7:", ":elephpant-10-9-ca38:", ":elephpant-11-9-461c:", ":elephpant-12-9-42a1:", ":elephpant-13-9-11aa:", ":elephpant-14-9-caaa:", ":elephpant-15-9-c26a:", ":elephpant-16-9-c491:", ":elephpant-1-10-a823:", ":elephpant-2-10-ad00:", ":elephpant-3-10-9793:", ":elephpant-4-10-b19c:", ":elephpant-5-10-7443:", ":elephpant-6-10-0656:", ":elephpant-7-10-e164:", ":elephpant-8-10-4058:", ":elephpant-9-10-f63a:", ":elephpant-10-10-effb:", ":elephpant-11-10-4509:", ":elephpant-12-10-e07e:", ":elephpant-13-10-7fa1:", ":elephpant-14-10-b864:", ":elephpant-15-10-1bd8:", ":elephpant-16-10-1c2b:", ":elephpant-1-11-62d6:", ":elephpant-2-11-5301:", ":elephpant-3-11-3776:", ":elephpant-4-11-0c90:", ":elephpant-5-11-2359:", ":elephpant-6-11-0650:", ":elephpant-7-11-a650:", ":elephpant-8-11-4ce5:", ":elephpant-9-11-4cc0:", ":elephpant-10-11-9be9:", ":elephpant-11-11-8afe:", ":elephpant-12-11-f9eb:", ":elephpant-13-11-b02c:", ":elephpant-14-11-82dc:", ":elephpant-15-11-c2c2:", ":elephpant-16-11-cd41:", ":elephpant-1-12-6104:", ":elephpant-2-12-bf93:", ":elephpant-3-12-79ef:", ":elephpant-4-12-cc64:", ":elephpant-5-12-14e0:", ":elephpant-6-12-96b6:", ":elephpant-7-12-87b0:", ":elephpant-8-12-44ca:", ":elephpant-9-12-6398:", ":elephpant-10-12-f50e:", ":elephpant-11-12-9c93:", ":elephpant-12-12-f8f6:", ":elephpant-13-12-eff3:", ":elephpant-14-12-18cd:", ":elephpant-15-12-1239:", ":elephpant-16-12-e586:", ":elephpant-1-13-b0c0:", ":elephpant-2-13-a744:", ":elephpant-3-13-125d:", ":elephpant-4-13-8ac9:", ":elephpant-5-13-7ed7:", ":elephpant-6-13-bc01:", ":elephpant-7-13-a69e:", ":elephpant-8-13-f7fc:", ":elephpant-9-13-ab43:", ":elephpant-10-13-0ce8:", ":elephpant-11-13-d022:", ":elephpant-12-13-717a:", ":elephpant-13-13-2e14:", ":elephpant-14-13-3aa2:", ":elephpant-15-13-a798:", ":elephpant-16-13-553e:", ":elephpant-1-14-daa3:", ":elephpant-2-14-ccb6:", ":elephpant-3-14-61d7:", ":elephpant-4-14-2597:", ":elephpant-5-14-3d72:", ":elephpant-6-14-e6c5:", ":elephpant-7-14-79e4:", ":elephpant-8-14-5243:", ":elephpant-9-14-d91b:", ":elephpant-10-14-26a1:", ":elephpant-11-14-decb:", ":elephpant-12-14-f5e3:", ":elephpant-13-14-a742:", ":elephpant-14-14-2ce8:", ":elephpant-15-14-3069:", ":elephpant-16-14-9b39:"}
	game := internal.NewGame(solution, 16, 9, ":blank:")
	//game.Randomise()

	logger.Info(game.String())

	app := slackGame{
		debug:     *debug,
		channelID: *channelID,
		logger:    *logger,
		client:    client,
		game:      game,
	}

	err := app.startGame()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	var repeat int
	var direction int

	for range 50 {
		time.Sleep(time.Millisecond * 100)

		repeat = rand.Intn(8) + 1
		direction = rand.Intn(7)
		for range repeat {

			switch direction {
			case 0:
				app.game.Up()
			case 1:
				app.game.Left()
			case 2:
				app.game.Right()
			case 3:
				app.game.Down()
			}
		}

		err = app.updateGame()
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}

}
