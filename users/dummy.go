package users

import (
	"github.com/eduardoacuna/self-esteem/notifications"
)

var dummyUserData = []*User{
	&User{
		ID:     "0",
		Email:  "eacuna@nearsoft.com",
		Notify: notifications.EachWeek,
		Day:    notifications.Friday,
	},
	&User{
		ID:     "1",
		Email:  "mvalle@nearsoft.com",
		Notify: notifications.EachWeek,
		Day:    notifications.Monday,
	},
	&User{
		ID:     "2",
		Email:  "izepeda@nearsoft.com",
		Notify: notifications.EachTwoWeeks,
		Day:    notifications.Monday,
	},
}
