package model

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessLevel int    `json:"accesslevel"`
	Tier        int    `json:"tier"`
	Balance     int64  `json:balance`
}
