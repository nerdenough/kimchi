package actions_test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/magiconair/properties/assert"
	"github.com/nerdenough/kimchi/service/actions"
)

func TestSimpleChat(t *testing.T) {
	testCases := []struct {
		name          string
		config        map[string]interface{}
		messageCreate *discordgo.MessageCreate
		expected      string
	}{
		{
			name:          "Basic response",
			config:        map[string]interface{}{"responses": []string{"hello"}},
			messageCreate: MockMessageCreate,
			expected:      "hello",
		},
		{
			name:          "Response with tokens",
			config:        map[string]interface{}{"responses": []string{"hello {author}"}},
			messageCreate: MockMessageCreate,
			expected:      "hello <@user_id>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			action, _ := actions.NewSimpleChat(tc.config)
			actual, _ := action.Process(nil, tc.messageCreate)
			assert.Equal(tt, actual, tc.expected)
		})
	}
}
