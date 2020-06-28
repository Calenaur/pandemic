package model

type Event struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      int    `json:"rarity"`
	Tier        int    `json:"tier"`
}
