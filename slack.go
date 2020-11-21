package awesomeProject

import "github.com/slack-go/slack"

//go:generate pegomock generate -m --use-experimental-model-gen --package mocks -o mocks/mock_slack_client.go SlackClient

type SlackClient interface {
	GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error)
}
