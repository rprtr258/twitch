# Twitch utils available from cli/bonzai

## Usage
```bash
go install ./cmd/twitch/
twitch var set nick <nickname>
twitch var set token <oauth:tokentokentokentoken>

# set channel to set to, default is <nickname>
twitch var set channel <channel>
# set how many messages to send, default is 1
twitch var set count <count>
# set delay between messages, default is 1s
twitch var set delay 1m

twitch send hi chat
```
