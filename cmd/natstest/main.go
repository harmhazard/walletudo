package main

import (
	"context"
	"github.com/dnbsd/xmrmux/services/wallet"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	s, err := wallet.New(wallet.Arguments{
		Logger:          logger,
		Name:            "wallet",
		Subject:         "wallet.1.rpc",
		Servers:         []string{"localhost"},
		WalletRpcServer: "http://localhost:18082/json_rpc",
	})
	if err != nil {
		panic(err)
	}

	err = s.Start(ctx)
	if err != nil {
		panic(err)
	}
}
