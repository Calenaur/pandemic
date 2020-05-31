package model

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Session     string `json:"session"`
	Manufacture int64  `json:"manufacture"`
}
