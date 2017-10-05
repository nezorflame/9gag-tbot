package main

import (
	"log"
	"strings"

	"golang.org/x/net/context"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
)

// Bot describes a Telegram bot
type Bot struct {
	*tg.BotAPI

	config     tg.UpdateConfig
	updateChan <-chan tg.Update
}

func newBot(api *tg.BotAPI, udTimeout int) (b *Bot, err error) {
	b = &Bot{
		BotAPI: api,
		config: tg.NewUpdate(0),
	}
	b.BotAPI.Debug = false
	b.config.Timeout = udTimeout

	if b.updateChan, err = b.GetUpdatesChan(b.config); err != nil {
		err = errors.Wrap(err, "Unable get updates from bot")
	}

	return
}

func (b *Bot) listenMessages(ctx context.Context) {
	for u := range b.updateChan {
		select {
		case <-ctx.Done():
			return
		default:
			if u.UpdateID >= b.config.Offset {
				b.config.Offset = u.UpdateID + 1
			}

			if u.Message == nil {
				continue
			}

			switch u.Message.Command() {
			case "hot":
				log.Printf("User from chat '%d' requested hot posts\n", u.Message.Chat.ID)
				if posts, err := apiClient.GetHotPosts("hot", 10); err == nil {
					for _, p := range posts {
						// skip pinned posts
						if strings.Contains(p.Title, "📌") {
							continue
						}
						// send according to the type
						switch strings.ToLower(p.Type) {
						case "image":
							err = b.sendPhoto(u.Message.Chat.ID, p.Images.Image700.URL, p.Title)
						case "animated", "video":
							err = b.sendDocument(u.Message.Chat.ID, p.Images.Image700.URL, p.Title)
						}
						// check error
						if err != nil {
							log.Println(err)
						}
					}
				} else {
					log.Println("Unable to get hot posts from API:", err)
				}
			case "start":
				log.Printf("User from chat '%d' started dialog with bot\n", u.Message.Chat.ID)
				if err := b.sendMessage(u.Message.Chat.ID, welcomeMessage); err != nil {
					log.Println(err)
				}
			default:
				log.Printf("User from chat '%d' said to bot: '%s'\n", u.Message.Chat.ID, u.Message.Text)
			}
		}
	}
}

func (b *Bot) sendMessage(chatID int64, msg string) error {
	msgConfig := tg.NewMessage(chatID, msg)
	msgConfig.ParseMode = tg.ModeMarkdown
	if _, err := b.Send(msgConfig); err != nil {
		return errors.Wrap(err, "Unable to send message")
	}
	return nil
}

func (b *Bot) sendPhoto(chatID int64, imageURL, caption string) error {
	photoConfig := tg.NewPhotoUpload(chatID, nil)
	photoConfig.Caption = caption
	photoConfig.FileID = imageURL
	photoConfig.UseExisting = true
	if _, err := b.Send(photoConfig); err != nil {
		return errors.Wrap(err, "Unable to send image")
	}

	return nil
}

func (b *Bot) sendDocument(chatID int64, imageURL, caption string) error {
	docConfig := tg.NewDocumentUpload(chatID, nil)
	docConfig.Caption = caption
	docConfig.FileID = imageURL
	docConfig.UseExisting = true
	if _, err := b.Send(docConfig); err != nil {
		return errors.Wrap(err, "Unable to send animated image")
	}

	return nil
}
