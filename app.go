package main

import (
	"context"
	"database/sql"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	a.CloseConnection()
}

func (a *App) Tables() []string {

	tables, _ := ShowTables(a.db)
	return tables
}

func (a *App) Table(tableName string) TableData {

	rows, _ := SelectAll(tableName, a.db)
	return rows
}

func (a *App) Connect(dbPassword string, dbPort string, dbName string) DBConnection {
	db, _ := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s)/%s", dbPassword, dbPort, dbName))
	connected := db.Ping()
	if connected != nil {
		return DBConnection{Success: false, Err: connected.Error()}
	}
	a.db = db
	return DBConnection{Success: true}
}

func (a *App) CloseConnection() {
	if a.db != nil { //?
		a.db.Close()
	}
}
