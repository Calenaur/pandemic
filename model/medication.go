package model

type Medication struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ResearchCost   int    `json:"research_cost"`
	MaximumTraits  int    `json:"maximum_traits"`
	BaseValue      int    `json:"base_value"`
	Tier           int    `json:"tier"`
}


type MedicationTrait struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Tier           int    `json:"tier"`
}

