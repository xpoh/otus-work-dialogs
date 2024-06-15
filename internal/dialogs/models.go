package dialogs

type DialogMessage struct {
	// Идентификатор пользователя
	From string `json:"from"`

	// Идентификатор пользователя
	To string `json:"to"`

	// Текст сообщения
	Text string `json:"text"`
}
