package infra

import (
	"fmt"

	"./def"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() (*gorm.DB, error) {
	connect := fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s",
		def.GetEnvUser(), def.GetEnvPass(), def.Protocol, def.Database, def.Charset, def.ParseTime, def.Loc)
	gormdb, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return gormdb, err
}
