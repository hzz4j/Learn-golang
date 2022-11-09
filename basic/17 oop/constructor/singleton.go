package main

import "sync"

var (
	sUser     *User
	sUserOnce sync.Once
)

func GetUserInstance() *User {
	if sUser == nil {
		sUserOnce.Do(func() {
			if sUser == nil {
				sUser = NewDefaultUser()
			}
		})
	}
	return sUser
}
