package model

type Friend struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	Tier    int    `json:"tier"`
}
