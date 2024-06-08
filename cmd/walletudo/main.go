package main

import (
	"context"
	"github.com/dnbsd/xmrmux/services/wallet"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	conf := NewConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: conf.LogLevel() == slog.LevelDebug,
		Level:     conf.LogLevel(),
	}))

	s := New(Arguments{
		Logger: logger.With(slog.String("service", "walletudo")),
		WalletArgs: wallet.Arguments{
			Logger:          logger.With(slog.String("service", "wallet")),
			Name:            "walletudo-wallet",
			Subject:         conf.NatsRpcSubject(),
			Servers:         conf.NatsServer(),
			WalletRpcServer: conf.WalletRpcServer(),
			WalletName:      conf.WalletName(),
		},
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	err := s.Start(ctx)
	if err != nil {
		os.Exit(1)
	}
}
