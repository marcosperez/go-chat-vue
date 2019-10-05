package common

// Channels por los cuales se comunican con el supervisor de chats
type Channels struct {
	ChatsChannel        chan ChatMessage
	SuscriptionsChannel chan SubscriptionMessage
}

// ChatMessage data de un mensaje de tipo chat
type ChatMessage struct {
	ChatID  string `json:"chatID"`
	Message string `json:"message"`
	UserID  string `json:"userID"`
}

// SendMessage que se envian al cliente
type SendMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// ReceiverMessage datos recibidos por WS
type ReceiverMessage struct {
	Type   string      `json:"type"`
	UserID string      `json:"userID"` // TODO: Buscar forma mas elegante de identificar una conexion
	Data   interface{} `json:"data"`
}

// SubscriptionMessage data de un mensaje de tipo suscription
type SubscriptionMessage struct {
	ChatID string `json:"chatID"`
}
