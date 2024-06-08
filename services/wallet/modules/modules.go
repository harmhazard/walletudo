package modules

import (
	"github.com/dnbsd/jsonrpc"
	"github.com/dnbsd/xmrmux/services/wallet/modules/rpc"
	"github.com/dnbsd/xmrmux/services/wallet/modules/wallet"
	"gitlab.com/moneropay/go-monero/walletrpc"
)

func New(client *walletrpc.Client) []jsonrpc.Module {
	return []jsonrpc.Module{
		rpc.New(),
		wallet.New(client),
	}
}
