package twitch

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

var Cmd = &Z.Cmd{
	Name: "twitch",
	Commands: []*Z.Cmd{
		help.Cmd, vars.Cmd, conf.Cmd,
		// {
		// 	Name:  "login",
		// 	Usage: "login using OAUTH token",
		// 	Flags: []cli.Flag{
		// 		&cli.StringFlag{
		// 			Name:     "nick",
		// 			Usage:    "user nickname",
		// 			Required: true,
		// 		},
		// 		&cli.StringFlag{
		// 			Name:     "token",
		// 			Usage:    "oauth token",
		// 			Required: true,
		// 		},
		// 	},
		// 	Action: func(ctx *cli.Context) error {
		// 		nick := ctx.String("nick")
		// 		token := ctx.String("token")

		// 		data, err := json.Marshal(AuthData{
		// 			Nick:  nick,
		// 			Token: token,
		// 		})
		// 		if err != nil {
		// 			return err
		// 		}

		// 		return os.WriteFile(_tokenFile, data, 0644)
		// 	},
		// },
		// {
		// 	Name:  "send",
		// 	Usage: "send message to chat",
		// 	Flags: []cli.Flag{
		// 		&cli.StringFlag{
		// 			Name:     "channel",
		// 			Aliases:  []string{"c"},
		// 			Usage:    "the channel you want to join",
		// 			Required: true,
		// 		},
		// 		&cli.IntFlag{
		// 			Name:    "count",
		// 			Aliases: []string{"n"},
		// 			Usage:   "how many times send message",
		// 			Value:   1,
		// 		},
		// 		&cli.DurationFlag{
		// 			Name:    "delay",
		// 			Aliases: []string{"d"},
		// 			Usage:   "delay between messages",
		// 			Value:   time.Second,
		// 		},
		// 	},
		// 	Action: func(ctx *cli.Context) error {
		// 		count := ctx.Int("count")
		// 		channel := ctx.String("channel")
		// 		delay := ctx.Duration("delay")

		// 		data, err := os.ReadFile(_tokenFile)
		// 		if err != nil {
		// 			return fmt.Errorf("getting token failed, try login again: %w", err)
		// 		}

		// 		var authData AuthData
		// 		if err := json.Unmarshal(data, &authData); err != nil {
		// 			return err
		// 		}

		// 		var message string
		// 		if ctx.Args().Len() > 0 {
		// 			message = strings.Join(ctx.Args().Slice(), " ")
		// 		} else {
		// 			var err error
		// 			message, err = readStdin()
		// 			if err != nil {
		// 				return err
		// 			}
		// 		}

		// 		// s = socket.socket()
		// 		// s.connect((HOST, PORT))
		// 		// s.send("PASS {}\r\n".format(PASS).encode("utf-8"))
		// 		// s.send("NICK {}\r\n".format(NICK).encode("utf-8"))
		// 		// s.send("JOIN {}\r\n".format(CHAN).encode("utf-8"))

		// 		// track = input()
		// 		// print(track)
		// 		// sock.send(f"PRIVMSG {CHAN} :{msg}\r\n".encode("utf-8"))

		// 		// s.close()

		// 		client := twitch.NewClient(authData.Nick, authData.Token)
		// 		go client.Connect()
		// 		for i := 0; i < count; i++ {
		// 			client.Say(channel, string(message))
		// 			fmt.Println(string(message))
		// 			time.Sleep(delay)
		// 		}

		// 		return nil
		// 	},
		// },
	},
	Shortcuts: Z.ArgMap{
		`started`:    {`var`, `started`},
		`duration`:   {`var`, `set`, `duration`},
		`warn`:       {`var`, `set`, `warn`},
		`prefix`:     {`var`, `set`, `prefix`},
		`prefixwarn`: {`var`, `set`, `prefixwarn`},
	},
	Summary: "various utilities for twitch chat",
}
