package actions

import "github.com/bwmarrin/discordgo"

// EventRemove represents the event remove action.
type EventRemove struct{}

// NewEventRemove creates a new event remove action.
func NewEventRemove() (Action, error) {
	return EventRemove{}, nil
}

// Process executes the event remove action.
func (action EventRemove) Process(m *discordgo.MessageCreate) error {
	return nil
}
