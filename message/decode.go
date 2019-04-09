package message

import (
	"errors"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nerdenough/kimchi/actions"
)

// These patterns are the regex tests used to associate a message with an action.
var patternMatchType = map[string]string{
	actions.KeyEventCreate: "(create|make|new|add|schedule).+?event",
	actions.KeyEventRemove: "(remove|delete|cancel).+?event",
	actions.KeyEventList:   "(list|show|what|when).+?event",
}

// DecodeActionType determines the intended action for a message.
func DecodeActionType(m *discordgo.MessageCreate) (string, error) {
	message := strings.ToLower(m.Content)

	// Test message against patterns to find the message type
	for typ, pattern := range patternMatchType {
		matched, err := regexp.Match(pattern, []byte(message))
		if err != nil {
			return "", err
		}
		if matched {
			return typ, nil
		}
	}

	// Unknown message type
	return "", errors.New("Unknown message type")
}
