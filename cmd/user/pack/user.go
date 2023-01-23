package pack

import (
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/user/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
)

// User pack user info
func User(u *db.User) *userdemo.User {
	if u == nil {
		return nil
	}

	return &userdemo.User{UserId: int64(u.ID), UserName: u.UserName, Avatar: "test"}
}

// Users pack list of user info
func Users(us []*db.User) []*userdemo.User {
	users := make([]*userdemo.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
