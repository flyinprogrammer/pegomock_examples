package awesomeProject_test

import (
	"awesomeProject"
	"awesomeProject/mocks"
	"github.com/petergtz/pegomock"
	"github.com/slack-go/slack"
	"testing"
)

func TestPostMessage(t *testing.T)  {
	pegomock.RegisterMockTestingT(t)
	mock := mocks.NewMockSlackClient()
	channelID := "something"
	text := "secret message"
	awesomeProject.MakeSomething(mock, channelID, text)
	mock.VerifyWasCalledOnce().PostMessage(channelID, slack.MsgOptionText(text, false))
}
