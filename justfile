@_help:
    just --list

# install cli application
@install:
    go install .

# test buy reward
@reward:
  rwenv -ie .env go run cmd/twitch/twitch.go rewards

