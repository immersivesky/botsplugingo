package updates

type Message struct {
	Chat	Chat	`json:"chat,omitempty"`
	User	User	`json:"user,omitempty"`
	Text	string	`json:"text,omitempty"`
}

type Chat struct {
	ID		int		`json:"id,omitempty"`
}

type User struct {
	ID		int		`json:"id,omitempty"`
}