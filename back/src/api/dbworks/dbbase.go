package dbworks

import (
	"fmt"

	"./def"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() (*sql.DB, *gorm.DB) {
	connect := fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s",
		def.GetEnvUser(), def.GetEnvPass(), def.Protocol, def.Database, def.Char, def.ParseTime, def.Loc,
	)

	sqlDB, err := sql.Open(def.DBAdmin, connect)

	if err != nil {
		panic(err.Error())
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return sqlDB, gormDB
}
