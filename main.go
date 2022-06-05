package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/modanisatech/marketplace/service-template/internal/config"
	"gitlab.com/modanisatech/marketplace/service-template/internal/foo"
	"gitlab.com/modanisatech/marketplace/service-template/pkg/server"
	"gitlab.com/modanisatech/marketplace/shared/kafka"
	"gitlab.com/modanisatech/marketplace/shared/log"
	"gitlab.com/modanisatech/marketplace/shared/unleash"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	logger := log.New()
	defer func() {
		err := logger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}()

	appEnv := os.Getenv("APP_ENV")
	conf, err := config.New(".config", appEnv)
	if err != nil {
		return err
	}
	conf.Print()

	_ = unleash.Init(conf.Appname, appEnv)

	if unleash.IsEnabled("golang-feature") {
		fmt.Println("Feature flag enabled")
	} else {
		fmt.Println("Feature flag disabled")
	}

	// if your flag have userId strategy, use this
	if unleash.IsEnabled("golang-feature", unleash.WithUserID("user12")) {
		fmt.Println("Feature flag enabled")
	} else {
		fmt.Println("Feature flag disabled")
	}

	kafkaProducer, err := kafka.NewProducer(conf.Kafka.URL)
	if err != nil {
		return err
	}
	kafkaConsumer := kafka.NewConsumer([]string{conf.Kafka.URL}, conf.Kafka.GroupID, logger)
	kafkaRouter := kafka.NewRouter(kafkaConsumer, logger)

	fooConsumer := foo.Consumer{Producer: kafkaProducer, Logger: logger}
	// every event received on repeat topic will be routed to RepeatEventHandler
	kafkaRouter.AddRoute("repeat", fooConsumer.RepeatEventHandler)

	handlers := []server.Handler{
		foo.Handler{},
	}

	s := server.New(conf.Server.Port, handlers, logger)

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	go s.Run()

	kafkaRouter.Start()
	<-shutdownChan

	s.Stop()
	kafkaRouter.Stop()
	kafkaProducer.Stop()

	return nil
}
