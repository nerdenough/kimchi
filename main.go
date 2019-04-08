package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// These keys are the available action types for the bot.
const (
	KeyEventCreate = "event.create"
	KeyEventRemove = "event.remove"
	KeyEventList   = "event.list"
)

// These patterns are the regex tests used to associate a message with an action.
var patternMatchType = map[string]string{
	KeyEventCreate: "(create|make|new|add).+?event",
	KeyEventRemove: "(remove|delete|cancel).+?event",
	KeyEventList:   "(list|show|what|when).+?event",
}

func main() {
	// Discord bot token should be provided as an environment variable.
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Printf("Missing BOT_TOKEN environment variable.\n")
		return
	}

	// Create a new Discord session for the bot.
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		fmt.Printf("Error creating discord session: %s\n", err.Error())
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Printf("Error opening connection: %s\n", err.Error())
		return
	}

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func decodeMessageType(m *discordgo.MessageCreate) (string, error) {
	message := strings.ToLower(m.Content)

	// Test message against patterns to find the message type
	for typ, pattern := range patternMatchType {
		matched, err := regexp.Match(pattern, []byte(message))
		if err != nil {
			return "", err
		}
		if matched {
			return typ, nil
		}
	}

	// Unknown message type
	return "", errors.New("Unknown message type")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore messages not intended for the bot
	if !strings.HasPrefix(m.Content, "!kimchi") {
		return
	}

	typ, err := decodeMessageType(m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "I'm sorry, I don't understand that command.")
	}

	if typ == "event.create" {
		s.ChannelMessageSend(m.ChannelID, "Okay, creating event.")
	}

	if typ == "event.remove" {
		s.ChannelMessageSend(m.ChannelID, "Okay, removing event.")
	}

	if typ == "event.list" {
		s.ChannelMessageSend(m.ChannelID, "Okay, here are your events:")
	}
}
