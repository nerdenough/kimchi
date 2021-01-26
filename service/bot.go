package service

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nerdenough/kimchi/service/actions"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	actions []actions.CustomAction
}

func NewBot(actions []actions.CustomAction) *Bot {
	return &Bot{
		actions: actions,
	}
}

func (b *Bot) GetActionFromMessage(msg string) (*actions.Action, error) {
	for _, customAction := range b.actions {
		matched, err := regexp.Match(customAction.Definition.Trigger, []byte(msg))
		if err != nil {
			return nil, nil
		}
		if matched {
			return &customAction.Action, nil
		}
	}

	return nil, nil
}

func (b *Bot) MessageReceived(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := strings.ToLower(m.Content)
	action, err := b.GetActionFromMessage(msg)
	if err != nil {
		log.Errorf("error getting action from message: %s", err)
	}
	if action == nil {
		// don't do anything
		return
	}

	a := *action
	resp, err := a.Process(s, m)
	if err != nil {
		log.Errorf("error processing action: %s", err)
		s.ChannelMessageSend(m.ChannelID, "Oops! Something went wrong.")
		return
	}

	if resp != "" {
		s.ChannelMessageSend(m.ChannelID, resp)
	}
}
