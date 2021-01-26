package actions

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/mitchellh/mapstructure"
	"github.com/nerdenough/kimchi/pkg/replace"
)

// SimpleChat represents a simple chat action.
type SimpleChat struct {
	config *SimpleChatConfig
}

// SimpleChatConfig represents the config for a simple chat action.
type SimpleChatConfig struct {
	Responses []string
}

// NewSimpleChat creates a new simple chat action.
func NewSimpleChat(config map[string]interface{}) (Action, error) {
	var actionConfig SimpleChatConfig
	err := mapstructure.Decode(config, &actionConfig)
	if err != nil {
		return nil, fmt.Errorf("error decoding simple chat config: %+v", config)
	}

	return SimpleChat{
		config: &actionConfig,
	}, nil
}

// Process executes the action.
func (a SimpleChat) Process(s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	resp := a.config.Responses[rand.Intn(len(a.config.Responses))]
	resp = replace.DiscordTokens(resp, *m.Message)

	return resp, nil
}
