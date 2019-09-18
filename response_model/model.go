package response_model

type Model struct {
	MessageOK struct {
		Code    string `json:"code"`
		Content string `json:"content"`
	} `json:"messsage_ok"`
	MeesageError struct {
		Code    string `json:"code"`
		Content string `json:"content"`
	} `json:"message_error"`
	Data interface{} `json:"data"`
}
