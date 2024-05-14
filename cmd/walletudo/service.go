package main

import (
	"context"
	"github.com/dnbsd/xmrmux/services/wallet"
	"golang.org/x/sync/errgroup"
	"log/slog"
)

type Arguments struct {
	Logger     *slog.Logger
	WalletArgs wallet.Arguments
}

type Service struct {
	args Arguments
}

func New(args Arguments) *Service {
	return &Service{
		args: args,
	}
}

func (s *Service) Start(ctx context.Context) error {
	s.args.Logger.Info("started service")
	defer s.args.Logger.Info("stopped service")

	group, groupCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		ws, err := wallet.New(wallet.Arguments{
			Logger:          s.args.WalletArgs.Logger,
			Name:            s.args.WalletArgs.Name,
			Subject:         s.args.WalletArgs.Subject,
			Servers:         s.args.WalletArgs.Servers,
			WalletRpcServer: s.args.WalletArgs.WalletRpcServer,
		})
		if err != nil {
			return err
		}
		return ws.Start(groupCtx)
	})

	err := group.Wait()
	if err != nil {
		s.args.Logger.Error("service stopped with an error", "error", err)
		return err
	}

	return nil
}
