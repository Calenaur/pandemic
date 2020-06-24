package model

type UserMedication struct {
	ID             int       `json:"id"`
	Medication     int       `json:"medication"`
	Traits         []int     `json:"traits,omitempty"`
}