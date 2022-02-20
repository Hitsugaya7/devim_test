package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Hitsugaya/rest-api-project/cmd/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
)

const sourceURL = "file://db/migration"
const databaseName = "postgres"

func NewDb(databaseConfig *config.DatabaseConfig) (db *sql.DB, err error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		databaseConfig.Database.Username,
		databaseConfig.Database.Password,
		databaseConfig.Database.Server,
		databaseConfig.Database.Port,
		databaseConfig.Database.DbName,
		databaseConfig.Database.SslMode)
	db, err = sql.Open("postgres", dataSourceName)
	//db, err := sql.Open("postgres", "postgres://yerdos:123@db:5432/db_name?sslmode=disable")

	//defer db.Close()
	if err != nil {
		return nil, fmt.Errorf("Error Opening DB: %v \n", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error Pinging DB: %v \n", err)
	}

	log.Println("Connected to db!")
	return db, nil
}

func LaunchMigration(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Error Creating Postgres Instance DB: %v \n", err)
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, databaseName, driver)
	if err != nil {
		return fmt.Errorf("Error Creating Migration Instance DB: %v \n", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Error launching migration: %v \n", err)
	}
	log.Println("Migration has been done")
	return nil
}
