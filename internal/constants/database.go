package constants

import "fmt"

var (
	DATABASE_FOLDER_PATH = fmt.Sprintf(".local/share/%s", APP_NAME)
)

const (
	DATABASE_FILE_NAME  = "database.db"
	DATABASE_DRIVER     = "sqlite3"
	DATABASE_MIGRATIONS = "file://migrations"
)
