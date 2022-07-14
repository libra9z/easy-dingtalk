package contact

import (
	"github.com/libra9z/easy-dingtalk/oauth2"
	"github.com/libra9z/easy-dingtalk/utils"
)

type Contact interface {
	utils.DingIdReduceAble
	GetUserInfo(userid string) (res *GetUserInfoResponseResult, err error)
}

func NewContact(oauth2 oauth2.Oauth2) (r Contact) {
	return &inner{
		oauth2: oauth2,
	}
}
