package main

import (
	"flag"
	"gitlab.com/faemproject/backend/core/shared/jobqueue"
	"gitlab.com/faemproject/backend/core/shared/logs"
	"gitlab.com/faemproject/backend/core/shared/os"
	"gitlab.com/faemproject/backend/core/shared/web"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/config"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/handler"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/server"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/telegrambot"
	"log"
	"time"
)

const (
	defaultConfigPath     = "deployment/devlocal/courierbot.toml"
	serverShutdownTimeout = 30 * time.Second
	maxRequestsAllowed    = 1000
)

func main() {
	//Parse flags
	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	cfg, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	if err := logs.SetLogLevel(cfg.Application.LogLevel); err != nil {
		log.Fatalf("Failed to set log level: %v", err)
	}
	if err := logs.SetLogFormat(cfg.Application.LogFormat); err != nil {
		log.Fatalf("Failed to set log format: %v", err)
	}
	logger := logs.Eloger
	tlgBot := telegrambot.BotClient{
		Token: cfg.Application.TelegramToken,
	}
	if err = tlgBot.Init(); err != nil {
		logger.Errorf("failed to init the telegram: %v", err)
	}
	hdlr := handler.Handler{
		Config:      cfg,
		Telegram:    &tlgBot,
		TelegramBot: tlgBot.Bot,
		Jobs:        jobqueue.NewJobQueues(),
	}
	//Telegram Subscriber
	tlgSub := telegrambot.Subscriber{
		Bot:     tlgBot.Bot,
		Handler: &hdlr,
	}
	if err = tlgSub.Init(); err != nil {
		logger.Errorf("failed to handling telegram messages: %v", err)
	}
	router := web.NewRouter()
	rest := server.Rest{
		Router:  router,
		Handler: &hdlr,
	}
	rest.Route()

	// Start an http server and remember to shut it down
	go web.Start(router, cfg.Application.Port)
	defer web.Stop(router, serverShutdownTimeout)
	<-os.NotifyAboutExit()
}
