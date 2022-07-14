package dingtalk_test

import (
	"fmt"
	"testing"

	"github.com/libra9z/easy-dingtalk/message"
)

func TestMessageSendToConversation(t *testing.T) {
	var err error
	defer deferErr(&err)
	msg := &message.MessageRequest{
		Msgtype: "text",
		Text: &message.TextMessage{
			Content: "这是一段文本消息",
		},
	}
	receiver, err := client.Message().SendToConversation("user0", 123453556, msg)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", receiver)
}

func TestMessageCorpconversationaSyncsendV2(t *testing.T) {
	var err error
	defer deferErr(&err)
	msg := &message.MessageRequest{
		Msgtype: "text",
		Text: &message.TextMessage{
			Content: "这是一段文本消息",
		},
	}
	taskId, err := client.Message().CorpconversationaSyncsendV2([]string{"user0"}, nil, false, msg)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", taskId)
}

func TestMessageCorpconversationaSyncsendV2Markdown(t *testing.T) {
	var err error
	defer deferErr(&err)
	msg := &message.MessageRequest{
		Msgtype: "markdown",
		Markdown: &message.MarkdownMessage{
			Text:  "这是一段文本消息",
			Title: "标题",
		},
	}
	taskId, err := client.Message().CorpconversationaSyncsendV2([]string{"user0"}, nil, false, msg)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", taskId)
}
