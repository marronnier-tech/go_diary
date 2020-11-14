package def

import (
	"os"
)

const (
	DBAdmin   = "mysql"
	Protocol  = "tcp(127.0.0.1:3306)"
	Database  = "daily_todo"
	Charset   = "utf8mb4"
	ParseTime = "True"
	Loc       = "Local"
)

func GetEnvUser() string {
	return os.Getenv("user")
}

func GetEnvPass() string {
	return os.Getenv("envpass")
}
