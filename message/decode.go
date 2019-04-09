package message

import (
	"errors"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var patternMatchHelpers = map[string]string{
	"time": `\d+(a|p)\.?m`,
	"date": `\s(on|this|next|every).\w+`,
}

// DecodeActionType determines the intended action for a message.
func DecodeActionType(m *discordgo.MessageCreate, types map[string]string) (string, error) {
	message := strings.ToLower(m.Content)

	// Test message against patterns to find the message type
	for typ, pattern := range types {
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

// DecodeTime extracts a relative time from a message.
func DecodeTime(msg string) (string, error) {
	reg := regexp.MustCompile(patternMatchHelpers["time"])
	matches := reg.FindStringSubmatch(msg)
	if len(matches) == 0 {
		return "", errors.New("Unable to decode time")
	}

	reg = regexp.MustCompile("[^a-zA-Z0-9]+")
	time := reg.ReplaceAllString(matches[0], "")
	return time, nil
}
