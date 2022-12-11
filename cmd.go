package twitch

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
}

var Cmd = &Z.Cmd{
	Name:    "twitch",
	Summary: "various utilities for twitch chat",
	Commands: []*Z.Cmd{
		help.Cmd, // vars.Cmd, conf.Cmd,
		{
			Name:  "login",
			Usage: "login using twitch OAUTH token",
			Commands: []*Z.Cmd{
				help.Cmd, vars.Cmd,
				// conf
				// &cli.StringFlag{Name: "nick", Usage: "user nickname", Required: true},
				// &cli.StringFlag{Name: "token", Usage: "oauth token", Required: true},
			},
			Call: func(x *Z.Cmd, args ...string) error {
				// nick, err := x.Caller.C("nick")
				// if err != nil {
				// 	return err
				// }
				// if nick == "" {
				// 	return errors.New("nick is not set")
				// }
				// x.Caller.Set("nick", nick)

				// token, err := x.Caller.C("token")
				// if err != nil {
				// 	return err
				// }
				// if token == "" {
				// 	return errors.New("token is not set")
				// }
				// x.Caller.Set("token", token)

				return nil
			},
		},
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
}

// func readStdin() (string, error) {
// 	input := []string{}
// 	buf := bufio.NewScanner(os.Stdin)
// 	for buf.Scan() {
// 		input = append(input, buf.Text())
// 	}
// 	if err := buf.Err(); err != nil {
// 		return "", err
// 	}

// 	return strings.Join(input, " "), nil
// }
