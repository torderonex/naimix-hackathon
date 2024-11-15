package service

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/webapi/sber"
	"log/slog"
)

func Test(str string) (string, error) {
	gigachat := sber.New()
	ll, err := gigachat.DoPrompt("Сколько рыбы в море?")
	slog.Info(ll)
	return ll, err
}
