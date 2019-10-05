package models

// ChannelMessage generados por los usuarios
// type ChannelMessage struct {
// 	messageText string
// 	date        time.Time
// }

// Channel Objeto que contiene un channel particular
type Channel struct {
	ID       string
	Users    map[string]string
	Messages []ChannelMessage
}
