package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/urfave/cli/v2"
)

var (
	_tokenFile = path.Join(os.Getenv("HOME"), ".twitch-cli")
)

type AuthData struct {
	Nick  string
	Token string
}

var app = cli.App{
	Name:  "twitch",
	Usage: "various twitch utilities for twitch chat",
	Commands: []*cli.Command{
		{
			Name:  "login",
			Usage: "login using OAUTH token",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "nick",
					Usage:    "user nickname",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "token",
					Usage:    "oauth token",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				nick := ctx.String("nick")
				token := ctx.String("token")

				data, err := json.Marshal(AuthData{
					Nick:  nick,
					Token: token,
				})
				if err != nil {
					return err
				}

				return os.WriteFile(_tokenFile, data, 0644)
			},
		},
		{
			Name:  "send",
			Usage: "send message to chat",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "channel",
					Aliases:  []string{"c"},
					Usage:    "the channel you want to join",
					Required: true,
				},
				&cli.IntFlag{
					Name:    "count",
					Aliases: []string{"n"},
					Usage:   "how many times send message",
					Value:   1,
				},
				&cli.DurationFlag{
					Name:    "delay",
					Aliases: []string{"d"},
					Usage:   "delay between messages",
					Value:   time.Second,
				},
			},
			Action: func(ctx *cli.Context) error {
				count := ctx.Int("count")
				channel := ctx.String("channel")
				delay := ctx.Duration("delay")

				data, err := os.ReadFile(_tokenFile)
				if err != nil {
					return fmt.Errorf("getting token failed, try login again: %w", err)
				}

				var authData AuthData
				if err := json.Unmarshal(data, &authData); err != nil {
					return err
				}

				message, err := readStdin()
				if err != nil {
					return err
				}

				client := twitch.NewClient(authData.Nick, authData.Token)
				go client.Connect()
				for i := 0; i < count; i++ {
					client.Say(channel, string(message))
					fmt.Println(string(message))
					time.Sleep(delay)
				}

				return nil
			},
		},
	},
}

func readStdin() (string, error) {
	input := []string{}
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		input = append(input, buf.Text())
	}
	if err := buf.Err(); err != nil {
		return "", err
	}

	return strings.Join(input, " "), nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
