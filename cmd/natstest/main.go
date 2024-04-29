package main

import (
	"context"
	"github.com/dnbsd/xmrmux/services/lightwallet"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	s, err := lightwallet.New(lightwallet.Arguments{
		Logger:          logger,
		Name:            "lightwallet-test",
		Servers:         []string{"localhost"},
		WalletRpcServer: "http://localhost:18082/json_rpc",
	})
	if err != nil {
		panic(err)
	}

	err = s.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
