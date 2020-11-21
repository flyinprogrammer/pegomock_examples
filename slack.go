package awesomeProject

import "github.com/slack-go/slack"

//go:generate pegomock generate -m --use-experimental-model-gen --package mocks -o mocks/mock_slack_client.go SlackClient

type SlackClient interface {
	PostMessage(channelID string, options ...slack.MsgOption) (respChannel string, respTimestamp string, err error)
}

func MakeSomething(client SlackClient, channelID string, text string) error {
	_, _, err := client.PostMessage(channelID, slack.MsgOptionText(text, false))
	return err
}
