package model

type Disease struct {
	Tier        int    `json:"tier"`
	Name        string `json:"name"`
	Description string `json:"accesslevel"`
	Rarity      int    `json:"rarity"`
}
