package actions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nerdenough/kimchi/message"
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

	// Find the event time
	time, err := message.DecodeTime(m.Content)
	if err != nil {
		return err
	}

	fmt.Printf("Time: %s\n", time)
	// TODO get date and create event

	return nil
}
