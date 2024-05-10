package wallet

import (
	"github.com/dnbsd/jsonrpc"
	"github.com/dnbsd/xmrmux/services/wallet/modules/wallet/modules/account"
	"gitlab.com/moneropay/go-monero/walletrpc"
)

func New(client *walletrpc.Client) jsonrpc.Module {
	m := module{
		client: client,
	}
	return jsonrpc.Module{
		Name: "wallet",
		Submodules: []jsonrpc.Module{
			account.New(client),
		},
		Methods: []jsonrpc.Method{
			{
				Name:    "init",
				Handler: m.Init,
			},
			{
				Name:    "listAccounts",
				Handler: m.ListAccounts,
			},
			{
				Name:    "setDaemon",
				Handler: m.SetDaemon,
			},
		},
	}
}

type module struct {
	client *walletrpc.Client
}

func (m *module) Init(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	// TODO: !!!!!!!!!!!!!!!!!
	// TODO: validate params!
	// TODO: !!!!!!!!!!!!!!!!!

	filename, err := params.String("filename")
	if err != nil {
		return c.Error(err)
	}

	err = m.client.OpenWallet(c.Context(), &walletrpc.OpenWalletRequest{
		Filename: filename,
	})
	if err != nil {
		err = m.client.CreateWallet(c.Context(), &walletrpc.CreateWalletRequest{
			Filename: filename,
			Language: "English",
		})
		if err != nil {
			// TODO: do not return monero rpc errors to the user!
			return c.Error(err)
		}
	}

	return c.Result(struct{}{})
}

func (m *module) SetDaemon(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	// TODO: !!!!!!!!!!!!!!!!!
	// TODO: validate params!
	// TODO: !!!!!!!!!!!!!!!!!

	address, err := params.String("address")
	if err != nil {
		return c.Error(err)
	}

	err = m.client.SetDaemon(c.Context(), &walletrpc.SetDaemonRequest{
		Address: address,
		Trusted: false,
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	return c.Result(struct{}{})
}

type ListResponse struct {
	ID      uint64 `json:"id"`
	Label   string `json:"label"`
	Balance uint64 `json:"balance"`
}

func (m *module) ListAccounts(c *jsonrpc.Context) (any, error) {
	resp, err := m.client.GetAccounts(c.Context(), &walletrpc.GetAccountsRequest{})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	var res []ListResponse
	for i, acc := range resp.SubaddressAccounts {
		// TODO: skip Primary account (i=0)
		res = append(res, ListResponse{
			// TODO: there's a bug in walletrpc, account.AddressLabel is always 0
			ID:      uint64(i),
			Label:   acc.Label,
			Balance: acc.Balance,
		})
	}

	return c.Result(res)
}
