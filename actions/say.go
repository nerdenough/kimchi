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
func (action Say) Process(s *discordgo.Session, m *discordgo.MessageCreate) error {
	reg := regexp.MustCompile(PatternMatchType["say"])
	msg := reg.ReplaceAllString(m.Content, "")
	if msg != "" {
		s.ChannelMessageSend(m.ChannelID, msg)
	}
	return nil
}
