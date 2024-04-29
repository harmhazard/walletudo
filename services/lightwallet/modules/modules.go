package modules

import (
	"github.com/dnbsd/jsonrpc"
	"github.com/dnbsd/xmrmux/services/lightwallet/modules/wallet"
	"gitlab.com/moneropay/go-monero/walletrpc"
)

func New(client *walletrpc.Client) []jsonrpc.Module {
	return []jsonrpc.Module{
		wallet.New(client),
	}
}
