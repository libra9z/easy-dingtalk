//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package dingtalk

import (
	"github.com/google/wire"
	"github.com/libra9z/easy-dingtalk/calendar"
	calendar_v2 "github.com/libra9z/easy-dingtalk/calendar/v2"
	"github.com/libra9z/easy-dingtalk/contact"
	"github.com/libra9z/easy-dingtalk/meeting"
	"github.com/libra9z/easy-dingtalk/message"
	"github.com/libra9z/easy-dingtalk/oauth2"
	"github.com/libra9z/easy-dingtalk/utils"
)

func NewDingtalk(opt utils.DingtalkOptions) (Dingtalk, func(), error) {
	panic(wire.Build(calendar_v2.NewCalendar, meeting.NewMeeting, message.NewMessage, calendar.NewCalendar, oauth2.NewOuath2, contact.NewContact, newDingtalk))
}
