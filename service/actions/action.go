package actions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Action represents a single runnable action.
type Action interface {
	Process(s *discordgo.Session, m *discordgo.MessageCreate) (string, error)
}

// CustomAction represents a custom action.
type CustomAction struct {
	Definition ActionDefinition
	Action     Action
}

// ActionDefinition represents the definition for a custom action.
type ActionDefinition struct {
	Name    string
	Type    string
	Trigger string
	Config  map[string]interface{}
}

var customActions = map[string]func(config map[string]interface{}) (Action, error){
	"apiRequest": NewAPIRequest,
	"simpleChat": NewSimpleChat,
}

// NewAction creates a new custom action.
func NewAction(def ActionDefinition) (*CustomAction, error) {
	customActionFunc, ok := customActions[def.Type]
	if !ok {
		return nil, fmt.Errorf("invalid custom action type: %s", def.Type)
	}

	action, err := customActionFunc(def.Config)
	if err != nil {
		return nil, err
	}

	return &CustomAction{
		Definition: def,
		Action:     action,
	}, nil
}
