package todolist

import "time"

type Session struct {
	ID         int
	User_ID    int
	Created_At time.Time
	Expires_at time.Time
}
