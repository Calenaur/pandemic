package model

type Disease struct {
	ID          int    `json:"id"`
	Tier        int    `json:"tier"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      int    `json:"rarity"`
}

type DiseaseMedication struct {
	Diseases      int `json:"disease"`
	Medications   int `json:"medication"`
	Effectiveness int `json:"effectiveness"`
}
