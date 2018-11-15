package slack

import (
	"fmt"
	"github.com/nlopes/slack"
)

// Slack defines sending of alerts via Slack
type Slack struct {
	token string
	channelID string
	api *slack.Api
}

// New provides init of the Slack
func New(token string) *Slack {
	api := slack.New(token)
	return &Slack{
		api:api,
		token:token,
	}
}

// Send provides sending message to Slack
func (s *Slack) Send() error {
	channelID, timestamp, err := s.PostMessage(s.channelID, slack.MsgOptionText("Some text", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("unable to send message to Slack: %v", err)
	}
	
	return nil
}

