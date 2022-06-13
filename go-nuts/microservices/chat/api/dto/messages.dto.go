package dto

type SendMessageRequest struct {
	ChatName string
	From     string
	To       string
	Message  string
}
