package slack

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/saromanov/alerting/structs"
)

// Slack defines sending of alerts via Slack
type Slack struct {
	token     string
	channelID string
	api       *slack.Api
}

// New provides init of the Slack
func New(token string) *Slack {
	api := slack.New(token)
	return &Slack{
		api:   api,
		token: token,
	}
}

// Send provides sending message to Slack
func (s *Slack) Send(m *structs.Message) error {
	channelID, timestamp, err := s.PostMessage(s.channelID, slack.MsgOptionText(m.Text, false), nil)
	if err != nil {
		return fmt.Errorf("unable to send message to Slack: %v", err)
	}

	return nil
}
