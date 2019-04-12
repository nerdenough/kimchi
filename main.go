package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/nerdenough/kimchi/actions"

	"github.com/nerdenough/kimchi/message"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Discord bot token should be provided as an environment variable.
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Printf("Missing BOT_TOKEN environment variable.\n")
		return
	}

	// Create a new Discord session for the bot
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		fmt.Printf("Error creating discord session: %s\n", err.Error())
		return
	}

	// Handlers
	discord.AddHandler(messageCreateHandler)

	// Open the connection
	err = discord.Open()
	if err != nil {
		fmt.Printf("Error opening connection: %s\n", err.Error())
		return
	}

	// Keep bot running until interrupt
	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore messages not intended for the bot
	if !strings.HasPrefix(m.Content, "!kimchi") {
		return
	}

	// Determine the action type
	typ, err := message.DecodeActionType(m, actions.PatternMatchType)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "I'm sorry, I don't understand that command.")
		return
	}

	// Create the action
	action, err := actions.NewAction(typ)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Sorry, something went wrong.")
		return
	}

	// Process the action
	resp, err := action.Process(s, m)
	if err != nil {
		fmt.Printf("Error processing action: %s\n", err.Error())
		s.ChannelMessageSend(m.ChannelID, "Sorry, something went wrong.")
		return
	}

	// Ensure we have a response for the bot to say
	if resp != "" {
		s.ChannelMessageSend(m.ChannelID, resp)
	}
}
