package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/google/uuid"
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

				var message string
				if ctx.Args().Len() > 0 {
					message = strings.Join(ctx.Args().Slice(), " ")
				} else {
					var err error
					message, err = readStdin()
					if err != nil {
						return err
					}
				}

				// s = socket.socket()
				// s.connect((HOST, PORT))
				// s.send("PASS {}\r\n".format(PASS).encode("utf-8"))
				// s.send("NICK {}\r\n".format(NICK).encode("utf-8"))
				// s.send("JOIN {}\r\n".format(CHAN).encode("utf-8"))

				// track = input()
				// print(track)
				// sock.send(f"PRIVMSG {CHAN} :{msg}\r\n".encode("utf-8"))

				// s.close()

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
		{
			Name:  "rewards",
			Usage: "buy/list rewards",
			Action: func(ctx *cli.Context) error {
				main2()
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

func main2() error {
	client := &http.Client{}

	data, err := json.Marshal([]map[string]any{{
		"operationName": "RedeemCustomReward",
		"extensions": map[string]any{"persistedQuery": map[string]any{
			"version":    1,
			"sha256Hash": "d56249a7adb4978898ea3412e196688d4ac3cea1c0c2dfd65561d229ea5dcc42",
		}},
		"variables": map[string]any{"input": map[string]any{
			"channelID": "70930005",
			"cost":      1,
			// "textInput":     time.Now().Format(time.ANSIC),
			"prompt":        nil,
			"rewardID":      "50590c98-595b-49e2-a997-e22641bbfda0",
			"title":         "нас рать",
			"transactionID": uuid.New().String(),
		}},
	}})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"https://gql.twitch.tv/gql#origin=twilight",
		bytes.NewReader(data),
	)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Referer":          "https://www.twitch.tv/",
		"Authorization":    "OAuth " + os.Getenv("HEADER_AUTHORIZATION"),
		"Client-Id":        os.Getenv("HEADER_CLIENT_ID"),
		"Client-Integrity": os.Getenv("HEADER_CLIENT_INTEGRITY"),
		"X-Device-Id":      os.Getenv("HEADER_DEVICE_ID"),
	}
	for name, value := range headers {
		req.Header.Set(name, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", bodyText)
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
