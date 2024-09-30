package migrations

import (
	"database/sql"
	"embed"
	"log"

	"github.com/alirezazeynali75/exify/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

//go:embed mysql/*.sql
var migrations embed.FS

type dbLogger struct{}

func (dbLogger) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (dbLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (dbLogger) Print(v ...interface{}) {
	log.Print(v...)
}

func (dbLogger) Println(v ...interface{}) {
	log.Println(v...)
}

func (dbLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func MigrateDB(config *config.Mysql) error {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		return err
	}
	defer db.Close()
	goose.SetLogger(dbLogger{})
	goose.SetTableName("migrations")
	goose.SetDialect("mysql")
	goose.SetBaseFS(migrations)
	return goose.Up(db, "mysql", goose.WithAllowMissing())
}

func UndoMigrateDB(config *config.Mysql) error {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		return err
	}
	defer db.Close()
	goose.SetLogger(dbLogger{})
	goose.SetTableName("migrations")
	goose.SetDialect("mysql")
	goose.SetBaseFS(migrations)
	return goose.Down(db, "mysql", goose.WithAllowMissing())
}
