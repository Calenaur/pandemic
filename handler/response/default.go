package response

type Default struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	code    string `json:"code"`
	message string `json:"message"`
}

