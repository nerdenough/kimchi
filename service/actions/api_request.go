package actions

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/mitchellh/mapstructure"
	"github.com/nerdenough/kimchi/pkg/replace"
)

// APIRequest represents an api request action.
type APIRequest struct {
	client *http.Client
	config *APIRequestConfig
}

// APIRequestConfig represents the config for an api request action.
type APIRequestConfig struct {
	URL         string
	Method      string
	ContentType string
	Responses   []string
}

// NewAPIRequest creates a new api request action.
func NewAPIRequest(config map[string]interface{}) (Action, error) {
	var actionConfig APIRequestConfig
	err := mapstructure.Decode(config, &actionConfig)
	if err != nil {
		return nil, fmt.Errorf("error decoding api request config: %+v", config)
	}

	return APIRequest{
		client: &http.Client{},
		config: &actionConfig,
	}, nil
}

// Process executes the action.
func (a APIRequest) Process(s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	req, err := http.NewRequest(a.config.Method, a.config.URL, nil)
	if err != nil {
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	str := a.config.Responses[rand.Intn(len(a.config.Responses))]
	str = replace.DiscordTokens(str, *m.Message)
	str = replace.APIRequest(str, a.config.ContentType, body)

	return str, nil
}
