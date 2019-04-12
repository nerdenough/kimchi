package actions

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestSay(t *testing.T) {
	m := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "!kimchi say do it for the vine",
		},
	}
	action, _ := NewSay()
	resp, _ := action.Process(nil, m)
	assert.Equal(t, "do it for the vine", resp)
}
