package wallet

import (
	"context"
	"github.com/dnbsd/bun.go"
	"github.com/dnbsd/jsonrpc"
	"github.com/dnbsd/xmrmux/services/wallet/modules"
	"github.com/nats-io/nats.go"
	"gitlab.com/moneropay/go-monero/walletrpc"
	"log/slog"
)

type Arguments struct {
	Logger           *slog.Logger
	Name             string
	Subject          string
	SubjectDiscovery string
	Servers          []string
	WalletRpcServer  string
	WalletName       string
}

type Service struct {
	args   Arguments
	client *walletrpc.Client
	rpc    *jsonrpc.Service
}

func New(args Arguments) (*Service, error) {
	client := walletrpc.New(walletrpc.Config{
		Address: args.WalletRpcServer,
	})
	builder := jsonrpc.Builder{
		Modules: modules.New(client),
	}
	rpc, err := builder.Build()
	if err != nil {
		return nil, err
	}
	return &Service{
		args:   args,
		rpc:    rpc,
		client: client,
	}, nil
}

func (s *Service) initWallet(ctx context.Context, walletName string) error {
	err := s.client.OpenWallet(ctx, &walletrpc.OpenWalletRequest{
		Filename: walletName,
	})
	if err != nil {
		err = s.client.CreateWallet(ctx, &walletrpc.CreateWalletRequest{
			Filename: walletName,
			Language: "English",
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Start(ctx context.Context) error {
	s.args.Logger.Info("started service")
	defer s.args.Logger.Info("stopped service")

	err := s.initWallet(ctx, s.args.WalletName)
	if err != nil {
		s.args.Logger.Error("failed to initialize wallet", "error", err)
		return err
	}

	b := bun.New(bun.Arguments{
		Name:    s.args.Name,
		Servers: s.args.Servers,
	})
	b.ConnectedHandler = func(conn *nats.Conn) {
		s.args.Logger.Info("connected to nats server")
	}
	b.ReconnectedHandler = func(conn *nats.Conn) {
		s.args.Logger.Info("reconnected to nats server")
	}
	b.DisconnectedHandler = func(conn *nats.Conn, err error) {
		if err != nil {
			s.args.Logger.Error("disconnected from nats server", "error", err)
			return
		}
		s.args.Logger.Info("disconnected from nats server")
	}
	b.Subscribe(s.args.Subject, s.handleRpc)
	b.Subscribe(s.args.SubjectDiscovery, s.handleDiscovery)
	return b.Start(ctx)
}

func (s *Service) handleRpc(c *bun.Context) error {
	defer func() {
		r := recover()
		if r != nil {
			s.args.Logger.Error("rpc handler error", "err", r)
			return
		}
	}()

	var req jsonrpc.Request
	err := c.BindJSON(&req)
	if err != nil {
		s.args.Logger.Debug("received invalid rpc request", "error", err)
		return err
	}

	if req.IsNotification() {
		s.args.Logger.Debug("received invalid rpc request", "error", "is a notification")
		return nil
	}

	s.args.Logger.Debug("received rpc request", "id", req.ID, "method", req.Method, "params", req.Params)

	resp := s.rpc.Call(c.Context(), req)
	if resp.Error != nil {
		s.args.Logger.Debug("received error rpc response", "id", req.ID, "error", resp.Error)
	} else {
		s.args.Logger.Debug("received rpc response", "id", req.ID, "result", resp.Result)
	}

	return c.JSON(resp)
}

func (s *Service) handleDiscovery(c *bun.Context) error {
	defer func() {
		r := recover()
		if r != nil {
			s.args.Logger.Error("service discovery handler error", "err", r)
			return
		}
	}()

	return c.String(s.args.Subject)
}
