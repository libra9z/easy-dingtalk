package message

import (
	"github.com/libra9z/easy-dingtalk/oauth2"
	"github.com/libra9z/easy-dingtalk/utils"
)

type Message interface {
	utils.DingIdReduceAble
	SendToConversation(userid string, cid int, msg *MessageRequest) (receiver []string, err error)

	CorpconversationaSyncsendV2(useridList []string, deptIdList []string, toAllUser bool, msg *MessageRequest) (taskId int, err error)
}

func NewMessage(oauth2 oauth2.Oauth2) (r Message) {
	return &inner{
		oauth2: oauth2,
	}
}
