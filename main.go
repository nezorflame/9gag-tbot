package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nezorflame/ninegago"
)

var (
	username, password, botToken string

	apiClient *ninegago.APIClient
)

func init() {
	flag.StringVar(&username, "u", "", "9GAG username")
	flag.StringVar(&password, "p", "", "9GAG password")
	flag.StringVar(&botToken, "t", "", "Telegram bot token")
	flag.Parse()

	if username == "" || password == "" || botToken == "" {
		username, password, botToken = flag.Arg(0), flag.Arg(1), flag.Arg(2)
	}

	if username == "" || password == "" || botToken == "" {
		flag.Usage()
		log.Fatal("Wrong credentials")
	}
}

func main() {
	var (
		tgBotAPI *tg.BotAPI
		bot      *Bot
		err      error
	)

	ctx, exit := context.WithCancel(context.Background())

	apiClient = ninegago.NewAPIClient()
	if err = apiClient.Login(username, password); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Welcome, %s\n", apiClient.User.FullName)

	if tgBotAPI, err = tg.NewBotAPI(botToken); err != nil {
		log.Fatal(err)
	}

	if bot, err = newBot(tgBotAPI, 10); err != nil {
		log.Fatal(err)
	}

	log.Println(bot.Self.UserName, "is alive!")
	go bot.listenMessages(ctx)

	configureShutdown(ctx, exit)
	<-ctx.Done()
}

func configureShutdown(ctx context.Context, exit func()) {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGTERM)
	signal.Notify(sigChannel, os.Interrupt)

	go func() {
		for {
			select {
			case s := <-sigChannel:
				if s == os.Interrupt || s == syscall.SIGTERM {
					exit()
					log.Println("Shutting down")
					return
				}

			case <-ctx.Done():
				log.Println("Shutting down")
				return
			}
		}
	}()
}
