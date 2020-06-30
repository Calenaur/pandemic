package model

type Disease struct {
	ID          int    `json:"id"`
	Tier        int    `json:"tier"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      int    `json:"rarity"`
}

type DiseaseMedication struct {
	Diseases      string `json:'disease'`
	Medications   string `json:'medication'`
	Effectiveness string `json:'effectiveness'`
}
