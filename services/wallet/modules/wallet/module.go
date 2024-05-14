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

type InitRequest struct {
	Filename string
}

func NewInitRequest(params jsonrpc.Object) (InitRequest, error) {
	filename, err := params.String("filename")
	if err != nil {
		return InitRequest{}, err
	}

	if filename == "" {
		err := jsonrpc.NewErrorParamObjectValue("filename", "empty value")
		return InitRequest{}, err
	}

	return InitRequest{
		Filename: filename,
	}, nil
}

type InitResponse struct{}

func (m *module) Init(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	req, err := NewInitRequest(params)
	if err != nil {
		return c.Error(err)
	}

	err = m.client.OpenWallet(c.Context(), &walletrpc.OpenWalletRequest{
		Filename: req.Filename,
	})
	if err != nil {
		err = m.client.CreateWallet(c.Context(), &walletrpc.CreateWalletRequest{
			Filename: req.Filename,
			Language: "English",
		})
		if err != nil {
			// TODO: do not return monero rpc errors to the user!
			return c.Error(err)
		}
	}

	return c.Result(InitResponse{})
}

type SetDaemonRequest struct {
	Address string
}

func NewSetDaemonRequest(params jsonrpc.Object) (SetDaemonRequest, error) {
	address, err := params.String("address")
	if err != nil {
		return SetDaemonRequest{}, err
	}

	if address == "" {
		err := jsonrpc.NewErrorParamObjectValue("address", "empty value")
		return SetDaemonRequest{}, err
	}

	return SetDaemonRequest{
		Address: address,
	}, nil
}

type SetDaemonResponse struct{}

func (m *module) SetDaemon(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	req, err := NewSetDaemonRequest(params)
	if err != nil {
		return c.Error(err)
	}

	err = m.client.SetDaemon(c.Context(), &walletrpc.SetDaemonRequest{
		Address: req.Address,
		Trusted: false,
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	return c.Result(SetDaemonResponse{})
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
