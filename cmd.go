package twitch

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
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
		help.Cmd, vars.Cmd, conf.Cmd,
		{
			Name:  "send",
			Usage: "send message to chat",
			Call: func(x *Z.Cmd, args ...string) error {
				nick, err := x.Caller.Get("nick")
				if err != nil {
					return err
				}
				if nick == "" {
					return fmt.Errorf("%s var is not defined", "nick")
				}

				token, err := x.Caller.Get("token")
				if err != nil {
					return err
				}
				if token == "" {
					return fmt.Errorf("%s var is not defined", "token")
				}

				// count: "how many times send message"
				count := 1
				countStr, err := x.Caller.Get("count")
				if err == nil && countStr != "" {
					countInt, err := strconv.Atoi(countStr)
					if err == nil {
						count = countInt
					}
				}

				// channel: "the channel you want to join"
				channel := nick
				channelStr, err := x.Caller.Get("channel")
				if err == nil && channelStr != "" {
					channel = channelStr
				}

				// delay: "delay between messages"
				delay := 1 * time.Second
				delayStr, err := x.Caller.Get("delay")
				if err == nil && delayStr != "" {
					delayDuration, err := time.ParseDuration(delayStr)
					if err == nil {
						delay = delayDuration
					}
				}

				var message string
				if len(args) > 0 {
					message = strings.Join(args, " ")
				} else {
					var err error
					message, err = readStdin()
					if err != nil {
						return err
					}
				}

				fmt.Println(
					count,
					delay,
				)

				conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
				if err != nil {
					return err
				}
				defer conn.Close()

				if _, err := conn.Write([]byte(fmt.Sprintf(
					"PASS %s\r\n"+
						"NICK %s\r\n"+
						"JOIN #%s\r\n"+
						"PRIVMSG #%s :%s\r\n",
					token,
					nick,
					channel,
					channel, message,
				))); err != nil {
					return err
				}

				return nil
			},
		},
	},
}

func readStdin() (string, error) {
	var input string
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		input = buf.Text()
	}
	if err := buf.Err(); err != nil {
		return "", err
	}

	return input, nil
}
