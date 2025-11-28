# Slide

A sliding tiles game in Slack

![gameplay demo](https://raw.githubusercontent.com/zakkbob/slide/main/demos/game.gif)

## Commands

Run `/slide-test` to start a new game

![/start new game command demo](https://raw.githubusercontent.com/zakkbob/slide/main/demos/slash-command.gif)

### Default
Starts a default game using numbers 1 to 9

`/slide-test default <width> <height>`

e.g `/slide-test default 3 3`

- Width and height are optional, they default to 3x2
- The width and height must result in a board with 9 or less tiles (only numbers 1 to 9 used)

### Custom
Starts a custom game, infers dimensions from newlines and number of slack emojis
`/slide-test custom <board>`

e.g `/slide-test custom :one::two::three:`

`:four::five::six:`

- Pasting from #sticker-bot is an easy way to do it

## How to use

- Add the `Slide` bot to your channel (or join `#slide-test` in Slack)
- Start a game!

## Features

- [x] Play multiple games at once
- [ ] Win detection
- [x] Play with custom sized boards
- [x] Play with images
