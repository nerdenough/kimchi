package replace

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// DiscordTokens replaces a string with discord tokens.
func DiscordTokens(str string, m discordgo.Message) string {
	discordTokens := map[string]string{
		"author": fmt.Sprintf("<@%s>", m.Author.ID),
	}

	for key, token := range discordTokens {
		str = strings.ReplaceAll(str, fmt.Sprintf("{%s}", key), token)
	}

	return str
}
