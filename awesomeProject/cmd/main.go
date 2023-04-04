package main

import (
	"awesomeProject/config"
	"awesomeProject/service"
	"awesomeProject/storage"
	"awesomeProject/transport"
	"awesomeProject/transport/handlers"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	conf := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())

	storage, err := storage.NewStorage(ctx, conf)
	if err != nil {
		return err
	}

	defer cancel()

	gracefulShutdown(cancel)

	service, err := service.NewService(storage)
	if err != nil {
		return err
	}

	handler := handlers.NewHandlers(storage, service)

	srv := transport.NewServer(handler, conf)

	err = srv.Run(ctx)
	if err != nil {
		return err
	}

	return nil
}
func gracefulShutdown(ctx context.CancelFunc) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		log.Println(<-done)
		fmt.Println("Gracefully shutdown")
		ctx()
	}()
}
