package dingtalk

import (
	"github.com/libra9z/easy-dingtalk/calendar"
	calendar_v2 "github.com/libra9z/easy-dingtalk/calendar/v2"
	"github.com/libra9z/easy-dingtalk/contact"
	"github.com/libra9z/easy-dingtalk/meeting"
	"github.com/libra9z/easy-dingtalk/message"
	"github.com/libra9z/easy-dingtalk/oauth2"
	"github.com/libra9z/easy-dingtalk/utils"
)

type Dingtalk interface {
	SetDingDiReduceFn(fn utils.DingIdReduceFn)
	Oauth2() oauth2.Oauth2
	Contact() contact.Contact
	CalendarV2() calendar_v2.Calendar
	Calendar() calendar.Calendar
	Message() message.Message

	Meeting() meeting.Meeting
}
