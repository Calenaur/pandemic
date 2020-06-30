package response

type fmcMessage struct {
	To               string      `json:"to,omitempty"`
	RegisterationIDs []string    `json:"registeration_i_ds,omitempty"`
	Data             interface{} `json:"data,omitempty"`
}
