package model

type Researcher struct {
	Tier             int `json:"tier"`
	Researcher_speed int `json:"researcher_speed"`
	Salary           int `json:"salary"`
	Maximum_traits   int `json:"maximum_traits"`
	Rarity           int `json:"rarity"`
}

type ResearcherTrait struct {
	Tier        int    `json:"tier"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      int    `json:"rarity"`
}
