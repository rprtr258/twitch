package twitch

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"

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
	Name:     "twitch",
	Commands: []*cli.Command{},
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
