package model

type Researcher struct {
	ID              int `json:"id"`
	Tier            int `json:"tier"`
	ResearcherSpeed int `json:"researcher_speed"`
	Salary          int `json:"salary"`
	MaximumTraits   int `json:"maximum_traits"`
	Rarity          int `json:"rarity"`
}

type ResearcherTrait struct {
	ID          int    `json:"id"`
	Tier        int    `json:"tier"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      int    `json:"rarity"`
}

type ResearcherTrait struct {
	Tier        int    `json:"tier"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      int    `json:"rarity"`
}
