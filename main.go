package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mitchellh/mapstructure"
	"github.com/nerdenough/kimchi/config"
	"github.com/nerdenough/kimchi/service"
	"github.com/nerdenough/kimchi/service/actions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	config.Configure()
}

func buildActions(configs []actions.ActionDefinition) ([]actions.CustomAction, error) {
	customActions := []actions.CustomAction{}

	for _, config := range configs {
		action, err := actions.NewAction(config)
		if err != nil {
			return nil, err
		}
		customActions = append(customActions, *action)
	}

	return customActions, nil
}

func main() {
	token := viper.GetString("token")

	// Create a new Discord session for the bot
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		panic(fmt.Errorf("error creating discord session: %s", err))
	}

	var actionDefinitions []actions.ActionDefinition
	mapstructure.Decode(viper.Get("customActions"), &actionDefinitions)
	customActions, err := buildActions(actionDefinitions)
	if err != nil {
		panic(fmt.Errorf("error parsing custom action definitions: %s", err))
	}

	bot := service.NewBot(customActions)
	discord.AddHandler(bot.MessageReceived)

	err = discord.Open()
	defer discord.Close()
	if err != nil {
		panic(fmt.Errorf("error opening connection: %s", err))
	}

	// Keep bot running until interrupt
	log.Info("bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Info("shutting down")
}
