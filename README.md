# Kimchi

> An easily customisable discord bot

## Getting Started

### Discord

You will need to create a new Discord Bot in order to run Kimchi. To do this,
head on over to the [Discord Developer Portal][1] and create a new application.

In the "OAuth2" section of the settings, select `bot` as the scope for the
application, and `Send Messages` as the permission (the required permissions may
change in the future). Open the generated URL and invite the bot to your Discord
server.

### Bot Config

In order use Kimchi, you need to prvoide a config file. The sample below
specifies the bot token (you can find yours [here][0]), as well as a couple of
basic chat actions that the bot can respond to. You can add and customise these
to your liking.

The file should be saved as `bot.config.json` in the root directory for the
project.

```json
{
  "token": "your-bot-token",
  "logLevel": "info",
  "customActions": [
    {
      "name": "I lost the game",
      "type": "simpleChat",
      "trigger": "the game",
      "config": {
        "responses": ["_Sigh..._ Thanks, {author}. I lost the game."]
      }
    },
    {
      "name": "motd",
      "type": "simpleChat",
      "trigger": "^!motd",
      "config": {
        "responses": ["Sometimes people are just stupid :man_shrugging:"]
      }
    }
  ]
}
```

### Running Locally

See "Development Setup" below.

### Running with Docker Compose

You can easily start up your bot by using the `nerdenough/kimchi` docker image.
You will also need to mount your config, as shown in the example below.

```yml
version: "3"

services:
  kimchi:
    image: nerdenough/kimchi
    volumes:
      - ./bot.config.json:/go/src/app/bot.config.json
```

## Development Setup

### Prerequisites

Ensure you have [Go][0] set up on your machine.

Clone the repository into your Go workspace and install the dependencies.

```bash
# Clone repo
git clone git@github.com:nerdenough/kimchi.git

# Install dependencies
go get
```

### Running Kimchi

In order to run Kimchi, you will need your Bot Token. You can find this on the
`Bot` tab of your Discord application. You can then run Kimchi with the
following command:

```bash
go run main.go
```

### Tests

```
go test ./...
```

[0]: https://golang.org/doc/install
[1]: https://discordapp.com/developers/applications/
