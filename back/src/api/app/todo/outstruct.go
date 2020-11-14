package todo

type getAll struct {
	Todo  []TodoArray `json:"TodoArray"`
	limit int         `json:"Limit"`
}

type todoArray struct {
	Content  string `json:"Content"`
	UserID   int    `json:"UserID"`
	UserName string `json:"UserName"`
}

// type GetAll struct {
// 	Todo  []TodoArray `json:"TodoArray"`
// 	limit int         `json:"Limit"`
// 	page  int         `json:"Page"`
// 	order string      `json:"Order"`
// }

// type TodoArray struct {
// 	TodoObj TodoObjInfo `json:"TodoObj"`
// 	User    UserInfo    `json:"User"`
// }

// type TodoObjInfo struct {
// 	TodoID       int    `json:"TodoID"`
// 	Content      string `json:"Content"`
// 	LastAchieved int    `json:"LastAchieved"`
// }

// type UserInfo struct {
// 	UserID   int    `json:"UserID"`
// 	UserName string `json:"UserName"`
// 	UserHN   string `json:"UserHN"`
// 	UserImg  string `json:"UserImg"`
// }
