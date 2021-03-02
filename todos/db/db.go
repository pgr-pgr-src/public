package db

import (
    "todos/models"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Initialize() {
    // DBのコネクションを接続する
    db, _ = sql.Open("mysql", root:pass@/todo.db)
    
    // Task構造体(Model)を元にマイグレーションを実行する
    db.AutoMigrate(&models.Task{})
}

func Get() *sql.DB {
    return db
}

// DBのコネクションを切断する
func Close() {
    db.Close()
}