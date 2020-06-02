package model

import "database/sql"

type User struct {
	ID          int64          `json:"id"`
	Username    string         `json:"username"`
	Session     sql.NullString `json:"session"`
	Manufacture sql.NullInt64  `json:"manufacture"`
}
