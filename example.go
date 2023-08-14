package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/dailydismay/tgbotapi-mux/mux/dispatcher"
	"github.com/dailydismay/tgbotapi-mux/mux/filter"
	"github.com/dailydismay/tgbotapi-mux/mux/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type V5Filter = filter.Filter[*tgbotapi.Update]

func IsCommand(command string) V5Filter {
	return func(u *tgbotapi.Update) bool {
		return u.Message != nil && u.Message.IsCommand() && u.Message.Command() == command
	}
}

func main() {
	globalNoopMwStack := handler.CombineMiddleware([]handler.Middleware[*tgbotapi.Update]{
		func(next handler.HandleFn[*tgbotapi.Update]) handler.HandleFn[*tgbotapi.Update] {
			return func(u *tgbotapi.Update) error {
				fmt.Println("noop one")
				return next(u)
			}
		},
		func(next handler.HandleFn[*tgbotapi.Update]) handler.HandleFn[*tgbotapi.Update] {
			return func(u *tgbotapi.Update) error {
				fmt.Println("noop two")
				return next(u)
			}
		},
	})

	v5dispatcher := dispatcher.NewDefault[*tgbotapi.Update]()
	v5dispatcher.Register(IsCommand("start"), handler.WithMiddleware(
		func(u *tgbotapi.Update) error {
			// implement routed logic
			return nil
		},
		globalNoopMwStack,
	))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	bot, _ := tgbotapi.NewBotAPI("")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// !!! warning !!!
	// this is only an example
	// this code is not production-ready
	// but shows how to use dispatcher instance
	for {
		select {
		case <-ctx.Done():
			return
		case update, ok := <-updates:
			if !ok {
				return
			}
			// global error handler
			_ = v5dispatcher.Dispatch(&update)
		}
	}
}
