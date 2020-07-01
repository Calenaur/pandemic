package model

type Token struct {
	Token   string `json:"token"`
	Balance int64  `json:"balance"`
	Tier    int    `json:"tier"`
}
