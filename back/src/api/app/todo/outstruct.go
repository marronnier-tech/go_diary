package todo

import "../../domain"

type allTodoArray struct {
	TodoObj domain.TodoObjInfo    `json:"TodoObj"`
	User    domain.UserSimpleInfo `json:"User"`
}

type userTodoArray struct {
	User    domain.UserSimpleInfo `json:"User"`
	TodoObj []domain.TodoObjInfo  `json:"TodoObj"`
}
