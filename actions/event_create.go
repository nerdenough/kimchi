package actions

import (
	"github.com/bwmarrin/discordgo"
)

// EventCreate represents the event create action.
type EventCreate struct{}

// NewEventCreate creates a new event create action.
func NewEventCreate() (Action, error) {
	return EventCreate{}, nil
}

// Process executes the event create action.
func (action EventCreate) Process(s *discordgo.Session, m *discordgo.MessageCreate) error {
	s.ChannelMessageSend(m.ChannelID, "Creating a new event")
	return nil
}
