package todolist

import (
	"time"

	"github.com/lib/pq"
)

type TodoList struct {
	ID          int
	Title       string
	Description string
	Tags        pq.StringArray
	Done        bool
	Created_At  time.Time
	Updated_at  time.Time
}

type UsersTodoLists struct {
	ID      int
	User_ID int
	List_ID int
}
