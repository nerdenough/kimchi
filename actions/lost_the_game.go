package actions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// LostTheGame represents the lost the game action.
type LostTheGame struct{}

// NewLostTheGame creates a new lost the game action.
func NewLostTheGame() (Action, error) {
	return LostTheGame{}, nil
}

// Process executes the lost the game action.
func (action LostTheGame) Process(s *discordgo.Session, m *discordgo.MessageCreate) error {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("I lost the game. Screw you, <@%s>!", m.Author.ID))
	return nil
}
