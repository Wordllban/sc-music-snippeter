package telegram

import (
	"log"
	"sc-music-snippeter/lib/logger"
	"sc-music-snippeter/processor"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Client struct {
	bot *tg.BotAPI
}

func New(token string) Client {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		logger.LogError("Telegram: Failed to create client", err)
	}

	bot.Debug = true
	log.Printf("Telegram: Authorized on account %s", bot.Self.UserName)

	return Client{
		bot: bot,
	}
}

func (c *Client) Updates() {
	update := tg.NewUpdate(0)
	update.Timeout = 60

	updates, err := c.bot.GetUpdatesChan(update)
	if err != nil {
		logger.LogError("Telegram: Failed to read updates", err)
	}

	for upd := range updates {
		if upd.Message != nil && upd.Message.Text != "" {
			audioCutName :=  processor.UrlProcessor(upd.Message.Text)
			// Send the audio back to Telegram channel
			audioMsg := tg.NewAudioUpload(upd.Message.Chat.ID, audioCutName)
			c.bot.Send(audioMsg)
			logger.Log("Telegram: message sent")
		}
	}
}

func (c *Client) SendMessage(message tg.Chattable) {
	_, err := c.bot.Send(message)
	if err != nil {
		logger.LogError("Telegram: Unable to send message", err)
	}
}