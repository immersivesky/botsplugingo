package updates

type Update struct {
	Event	string	`json:"event,omitempty"`
	Connect	string	`json:"connect,omitempty"`
	Message	Message	`json:"message,omitempty"`
}