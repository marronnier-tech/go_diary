package todo

import (
	"time"

	"github.com/lib/pq"
)

/* type OutGetAll struct {
	Todo  []todoArray `json:"TodoArray"`
	limit int         `json:"Limit"`
}

type todoArray struct {
	Content  string `json:"Content"`
	UserID   int    `json:"UserID"`
	UserName string `json:"UserName"`
} */

type allTodoArray struct {
	TodoObj todoObjInfo `json:"TodoObj"`
	User    outUserInfo `json:"User"`
}

type todoObjInfo struct {
	TodoID       int         `json:"TodoID"`
	Content      string      `json:"Content"`
	CreatedAt    time.Time   `json:"CreatedAt"`
	LastAchieved pq.NullTime `json:"LastAchieved"`
}

type outUserInfo struct {
	UserID   int    `json:"UserID"`
	UserName string `json:"UserName"`
	UserHN   string `json:"UserHN"`
	UserImg  string `json:"UserImg"`
}

type userTodoArray struct {
	User    outUserInfo   `json:"User"`
	TodoObj []todoObjInfo `json:"TodoObj"`
}
