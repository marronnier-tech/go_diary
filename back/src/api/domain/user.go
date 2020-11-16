package domain

type UserSimpleInfo struct {
	UserID   int    `gorm:"column:id" json:"UserID"`
	UserName string `gorm:"column:name" json:"UserName"`
	UserHN   string `gorm:"column:handle_name" json:"UserHN"`
	UserImg  string `gorm:"column:img" json:"UserImg"`
}
