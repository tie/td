// Binary gotdecho provides example of Telegram echo bot.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/tg"
)

func run(ctx context.Context) error {
	logger, _ := zap.NewDevelopment(zap.IncreaseLevel(zapcore.DebugLevel))
	defer func() { _ = logger.Sync() }()

	dispatcher := tg.NewUpdateDispatcher()
	return telegram.BotFromEnvironment(ctx, telegram.Options{
		Logger:        logger,
		UpdateHandler: dispatcher,
	}, func(ctx context.Context, client *telegram.Client) error {
		sender := message.NewSender(client)
		dispatcher.OnNewMessage(func(ctx tg.UpdateContext, u *tg.UpdateNewMessage) error {
			m, ok := u.Message.(*tg.Message)
			if !ok || m.Out {
				return nil
			}

			_, err := sender.Reply(ctx, u).Text(ctx, m.Message)
			return err
		})
		return nil
	}, telegram.RunUntilCanceled)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(2)
	}
}
