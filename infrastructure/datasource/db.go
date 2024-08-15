package datasource

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"todo-app/core/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func NewDatabase() *sql.DB {
	db, _ := sql.Open(
		config.DB_TYPE,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DB_HOST,
			config.DB_PORT,
			config.DB_USER,
			config.DB_PASSWORD,
			config.DB_DATABASE_NAME))

	err := db.Ping()
	if err != nil {
		log.Fatalf("can not connect: %v", err)
	}

	db.SetMaxOpenConns(config.DB_MAX_CONNECTION_NUM)
	db.SetConnMaxIdleTime(time.Duration(config.DB_MAX_IDLE_TIME_MINUTE) * time.Minute)

	return db
}

func Migrate(sourceURL string) {
	m, err := migrate.New(
		sourceURL,
		fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
			config.DB_TYPE,
			config.DB_USER,
			config.DB_PASSWORD,
			config.DB_HOST,
			config.DB_PORT,
			config.DB_DATABASE_NAME))
	if err != nil {
		log.Fatal(err)
	}
	m.Up()
}
