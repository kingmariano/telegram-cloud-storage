package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/NicoNex/echotron/v3"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func getChannelId(bot *tgbotapi.BotAPI, channelUsername string) (int64, error) {
	var channelID int64
	//sending message to telegram channel
	msgConfig := tgbotapi.NewMessageToChannel(channelUsername, fmt.Sprintf("hello from %s", channelUsername))
	msg, err := bot.Send(msgConfig)
	if err != nil {

		return 0, err
	}
	channelID = msg.Chat.ID
	return channelID, nil
}

type BotApiConfig struct {
	BotToken  string
	ChannelID int64
	Bot       *tgbotapi.BotAPI
	BotUploader echotron.API
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("error loading bot token")
	}
	channelUsername := os.Getenv("TELEGRAM_CHANNEL_USERNAME")
	if channelUsername == "" {
		log.Fatal("error loading channel username")

	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("error loading port")
	}
	uploaderbot := echotron.NewAPI(botToken)
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Println(err)
		return
	}
	channelID, err := getChannelId(bot, channelUsername)
	if err != nil {
		log.Printf("error getting channel ID %v", err)
		return
	}
	botcfg := BotApiConfig{
		BotToken:  botToken,
		Bot:       bot,
		ChannelID: channelID,
		BotUploader: uploaderbot,
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/readiness", botcfg.HandleReadiness)
	router.Post("/upload", botcfg.HandleTelegramUpload)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Server started at port %s", port)
	log.Fatal(server.ListenAndServe())
}
