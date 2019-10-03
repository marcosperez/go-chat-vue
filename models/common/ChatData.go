package common

// ChatData data de un mensaje de tipo chat
type ChatData struct {
	ChatID  string `json:"chatID"`
	Message string `json:"message"`
	UserID  string `json:"userID"`
}

// SendMessage que se envian al cliente
type SendMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
