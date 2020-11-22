package align

import (
	"fmt"

	"gorm.io/gorm"
)

func ListOrder(db *gorm.DB, table string, joined bool, param string) (ordered *gorm.DB) {
	if joined {
		ordered = db.Order(fmt.Sprintf("%s.%s desc", table, param))
		return
	}

	ordered = db.Order(fmt.Sprintf("%s desc", param))
	return

}
