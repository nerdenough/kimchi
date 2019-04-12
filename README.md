# Kimchi

> A useful Discord bot

## Development Setup

### Prerequisites

Ensure you have [Go][0] set up on your machine.

This project uses [Go Modules][1], in order for you environment to support this,
you will need to have run commands with the `GO111MODULE=on` environment
variable. You can set this globally via your bash configuration with the
following:

```bash
export GO111MODULE=on
```

Clone the repository into your Go workspace and install the dependencies.

```bash
# Clone repo
git clone git@github.com:nerdenough/kimchi.git

# Install dependencies
go get
```

### Discord

You will need to create a new Discord Bot in order to run Kimchi. To do this,
head on over to the [Discord Developer Portal][2] and create a new application.

In the "OAuth2" section of the settings, select `bot` as the scope for the
application, and `Send Messages` as the permission (the required permissions
may change in the future). Open the generated URL and invite the bot to your
Discord server.

### Building

You can build Kimchi with the following command:

```bash
go build
```

### Running Kimchi

In order to run Kimchi, you will need your Bot Token. You can find this on the
`Bot` tab of your Discord application. You can then run Kimchi with the
following command:

```bash
BOT_TOKEN=[your bot token] ./kimchi
```

[0]: https://golang.org/doc/install
[1]: https://github.com/golang/go/wiki/Modules
[2]: https://discordapp.com/developers/applications/
