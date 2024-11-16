package sber

import gigachat "github.com/saintbyte/gigachat_api"

type GigaChatApi struct {
	chat *gigachat.Gigachat
}

func New() *GigaChatApi {
	return &GigaChatApi{
		chat: gigachat.NewGigachat(),
	}
}

func (api *GigaChatApi) DoPrompt(prompt string) (string, error) {
	return api.chat.Ask(prompt)
}
