package models

// ChannelMessage data de un mensaje de tipo channel
type ChannelMessage struct {
	ChannelID string `json:"channelID"`
	Message   string `json:"message"`
	UserID    string `json:"userID"`
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
	ChannelID string `json:"channelID"`
	UserID    string `json:"userID"` // TODO: Buscar forma mas elegante de identificar una conexion
}
