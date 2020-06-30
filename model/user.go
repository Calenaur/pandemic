package model

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessLevel int64  `json:"accesslevel"`
	Tier        int    `json:"tier"`
	Balance     int    `json:"balance"`
}
