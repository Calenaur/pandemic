package model

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessLevel int64  `json:"accesslevel"`
	Balance     int64  `json:"balance"`
	Manufacture int64  `json:"manufacture"`
}
