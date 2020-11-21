package domain

type TodoObjInfo struct {
	TodoID       int    `json:"TodoID"`
	Content      string `json:"Content"`
	CreatedAt    string `json:"CreatedAt"`
	LastAchieved string `json:"LastAchieved"`
}
