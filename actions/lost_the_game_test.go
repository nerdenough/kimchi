package actions

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestLostTheGame(t *testing.T) {
	m := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: MockUser,
		},
	}
	action, _ := NewLostTheGame()
	resp, _ := action.Process(nil, m)
	assert.Equal(t, "I lost the game. Screw you, <@user_id>!", resp)
}
