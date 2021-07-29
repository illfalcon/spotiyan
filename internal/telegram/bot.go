package telegram

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/illfalcon/spotiyan/internal/spoti"
	"github.com/illfalcon/spotiyan/internal/translator"
	"github.com/illfalcon/spotiyan/internal/yandex"
	"github.com/illfalcon/spotiyan/pkg/httperrors"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	service *translator.Service
}

func NewBot(service *translator.Service) *Bot {
	return &Bot{service: service}
}

func (b *Bot) Init() error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		return err
	}

	bot.Debug = true
	b.bot = bot

	return nil
}

func (b *Bot) Listen() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		result, err := b.handleText(update.Message.Text)
		if err != nil {
			b.SendWithRetry(tgbotapi.NewMessage(update.Message.Chat.ID, httperrors.WriteErrorAsString(err)))

			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
		msg.ReplyToMessageID = update.Message.MessageID
		b.SendWithRetry(msg)
	}

	return nil
}

func (b *Bot) handleText(text string) (string, error) {
	if isSpotifyLink(text) {
		spotifyID, err := spoti.GetTrackIDFromURL(text)
		if err != nil {
			return "", err
		}

		return b.service.TranslateSpotifyToYandex(spotifyID)
	}

	trackID, err := yandex.GetTrackIDFromURL(text)
	if err != nil {
		return "", err
	}

	return b.service.TranslateYandexToSpotify(trackID)
}

func isSpotifyLink(text string) bool {
	return strings.Contains(text, "spotify")
}

func (b *Bot) SendWithRetry(message tgbotapi.MessageConfig) {
	if _, err := b.bot.Send(message); err != nil {
		if _, err := b.bot.Send(message); err != nil {
			log.Print(err)
		}
	}
}
