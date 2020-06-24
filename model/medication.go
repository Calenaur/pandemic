package model

type Medication struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Research_cost  int    `json:"research_cost"`
	Maximum_traits int    `json:"maximum_traits"`
	Rarity         int    `json:"rarity"`
	Tier           int    `json:"tier"`
}
