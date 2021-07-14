package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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

		trackID, err := yandex.GetTrackIDFromURL(update.Message.Text)
		if err != nil {
			b.SendWithRetry(tgbotapi.NewMessage(update.Message.Chat.ID, httperrors.WriteErrorAsString(err)))
		}

		result, err := b.service.Translate(trackID)
		if err != nil {
			b.SendWithRetry(tgbotapi.NewMessage(update.Message.Chat.ID, httperrors.WriteErrorAsString(err)))
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
		msg.ReplyToMessageID = update.Message.MessageID
		b.SendWithRetry(msg)
	}

	return nil
}

func (b *Bot) SendWithRetry(message tgbotapi.MessageConfig) {
	if _, err := b.bot.Send(message); err != nil {
		if _, err := b.bot.Send(message); err != nil {
			log.Print(err)
		}
	}
}
