package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"stripe_api/config"
)

func DBConnect() (*sql.DB, error) {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return nil, err
	}

	dataSourceName := fmt.Sprintf("%s:@tcp(%s)/%s", loadConfig.Database.User, loadConfig.Database.Host,
		loadConfig.Database.DBName)
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Failed to open database connection")
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
