package actions_test

import (
	"github.com/bwmarrin/discordgo"
)

// MockUser is a mock Discord user.
var MockUser *discordgo.User = &discordgo.User{
	ID: "user_id",
}

var MockMessageCreate *discordgo.MessageCreate = &discordgo.MessageCreate{Message: &discordgo.Message{Author: MockUser}}
