package model

import "database/sql"

type User struct {
	ID          int64         `json:"id"`
	Username    string        `json:"username"`
	AccessLevel int64         `json:"accesslevel"`
	Balance     sql.NullInt64 `json:"balance"`
	Manufacture sql.NullInt64 `json:"manufacture"`
}
