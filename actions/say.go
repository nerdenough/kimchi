package actions

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

// Say represents the say action.
type Say struct{}

// NewSay creates a new say action.
func NewSay() (Action, error) {
	return Say{}, nil
}

// Process executes the say action.
func (action Say) Process(s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	reg, err := regexp.Compile(PatternMatchType["misc.say"])
	if err != nil {
		return "", err
	}

	return reg.ReplaceAllString(m.Content, ""), nil
}
