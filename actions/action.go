package actions

import (
	"github.com/bwmarrin/discordgo"
)

// These keys are the available action types for the bot.
const (
	KeyEventCreate = "event.create"
	KeyEventRemove = "event.remove"
	KeyEventList   = "event.list"
)

// PatternMatchType are the regex tests used to associate a message with an action.
var PatternMatchType = map[string]string{
	KeyEventCreate: "(create|make|new|add|schedule).+?event",
	KeyEventRemove: "(remove|delete|cancel).+?event",
	KeyEventList:   "(list|show|what|when).+?event",
}

// ActionGenerator represents a function that can generate an action from a message.
type ActionGenerator func() (Action, error)

var generatorLookup = map[string]ActionGenerator{
	KeyEventCreate: NewEventCreate,
	KeyEventRemove: NewEventRemove,
	KeyEventList:   NewEventList,
}

// Action represents a single runnable action.
type Action interface {
	Process(s *discordgo.Session, m *discordgo.MessageCreate) error
}

// NewAction returns a new action based on the specified type.
func NewAction(typ string) (Action, error) {
	return generatorLookup[typ]()
}
