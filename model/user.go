package model

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Session     string `json:"session"`
	SessionDate string `json:"date"`
	Manufacture int64  `json:manufacture`
}
