package models

import "database/sql"

type Task struct {
    sql.Model
    Text string
}