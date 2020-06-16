package model

type Disease struct {
	Tier        int    `json:"tier"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      int    `json:"rarity"`
}
