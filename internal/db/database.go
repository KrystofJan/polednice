package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/KrystofJan/tempus/internal/constants"
)

type Database struct {
	Instance *sql.DB
}

func New() (*Database, error) {
	dbPath, err := getDBPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Println("Did not find the database, creating one...")
	}

	db, err := sql.Open(constants.DATABASE_DRIVER, dbPath)
	if err != nil {
		return nil, err
	}

	return &Database{
		Instance: db,
	}, nil
}

func GetConnString() (string, error) {
	dbPath, err := getDBPath()
	if err != nil {
		return "", fmt.Errorf("Could not create the database")
	}
	return fmt.Sprintf("sqlite3://%s", dbPath), nil
}

// TODO: cache path
func getDBPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	databasePath := strings.Split(constants.DATABASE_FOLDER_PATH, "/")
	databasePath = append([]string{home}, databasePath...)
	dbDir := filepath.Join(databasePath...)
	err = os.MkdirAll(dbDir, 0755)
	if err != nil {
		return "", err
	}
	return filepath.Join(dbDir, constants.DATABASE_FILE_NAME), nil
}
