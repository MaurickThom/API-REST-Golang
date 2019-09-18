package response_model

type Model struct {
	MessageOK    MessageOK    `json:"messsage_ok"`
	MessageError MessageError `json:"message_error"`
	Data         interface{}  `json:"data"`
}

type MessageError struct {
	Code    string `json:"code"`
	Content string `json:"content"`
}

type MessageOK struct {
	Code    string `json:"code"`
	Content string `json:"content"`
}
