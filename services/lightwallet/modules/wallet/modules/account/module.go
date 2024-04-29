package account

import (
	"github.com/dnbsd/jsonrpc"
	"gitlab.com/moneropay/go-monero/walletrpc"
)

func New(client *walletrpc.Client) jsonrpc.Module {
	m := module{
		client: client,
	}
	return jsonrpc.Module{
		Name: "account",
		Methods: []jsonrpc.Method{
			{
				Name:    "create",
				Handler: m.Create,
			},
			{
				Name:    "listAddresses",
				Handler: m.ListAddresses,
			},
			{
				Name:    "createAddress",
				Handler: m.CreateAddress,
			},
			{
				Name:    "listTransactions",
				Handler: m.ListTransactions,
			},
		},
	}
}

type module struct {
	client *walletrpc.Client
}

type CreateResponse struct {
	AccountID uint64 `json:"accountID"`
}

func (m *module) Create(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	// TODO: !!!!!!!!!!!!!!!!!
	// TODO: validate params!
	// TODO: !!!!!!!!!!!!!!!!!

	label, err := params.String("label")
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.CreateAccount(c.Context(), &walletrpc.CreateAccountRequest{
		Label: label,
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	return c.Result(CreateResponse{
		AccountID: resp.AccountIndex,
	})
}

type ListAddressesResponse struct {
	Address string `json:"address"`
	Label   string `json:"label"`
	Balance uint64 `json:"balance"`
	Used    bool   `json:"used"`
}

func (m *module) ListAddresses(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	// TODO: !!!!!!!!!!!!!!!!!
	// TODO: validate params!
	// TODO: !!!!!!!!!!!!!!!!!

	// TODO: accountID must be greater than 0!!!

	accountID, err := params.Number("accountID")
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.GetAddress(c.Context(), &walletrpc.GetAddressRequest{
		// TODO: replace with Uint64 method
		AccountIndex: uint64(accountID.Uint()),
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	var res []ListAddressesResponse
	for _, addr := range resp.Addresses {
		res = append(res, ListAddressesResponse{
			Address: addr.Address,
			Label:   addr.Label,
			Balance: addr.Balance,
			Used:    addr.Used,
		})
	}

	return c.Result(res)
}

type CreateAddressResponse struct {
	Address string `json:"address"`
}

func (m *module) CreateAddress(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	// TODO: !!!!!!!!!!!!!!!!!
	// TODO: validate params!
	// TODO: !!!!!!!!!!!!!!!!!

	// TODO: accountID must be greater than 0!!!

	accountID, err := params.Number("accountID")
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.CreateAddress(c.Context(), &walletrpc.CreateAddressRequest{
		// TODO: replace with Uint64 method
		AccountIndex: uint64(accountID.Uint()),
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	return c.Result(CreateAddressResponse{
		Address: resp.Address,
	})
}

type ListTransactionsResponse struct {
	Incoming []ListTransactionsTransactionResponse `json:"incoming"`
	Outgoing []ListTransactionsTransactionResponse `json:"outgoing"`
	Pending  []ListTransactionsTransactionResponse `json:"pending"`
	Failed   []ListTransactionsTransactionResponse `json:"failed"`
	Pool     []ListTransactionsTransactionResponse `json:"pool"`
}

type ListTransactionsTransactionResponse struct {
	Address       string `json:"address"`
	Amount        uint64 `json:"amount"`
	Confirmations uint64 `json:"confirmations"`
	Fee           uint64 `json:"fee"`
	Height        uint64 `json:"height"`
	Timestamp     uint64 `json:"timestamp"`
	TxID          string `json:"txid"`
	UnlockTime    uint64 `json:"unlock_time"`
}

func (m *module) ListTransactions(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	// TODO: !!!!!!!!!!!!!!!!!
	// TODO: validate params!
	// TODO: !!!!!!!!!!!!!!!!!

	accountID, err := params.Number("accountID")
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.GetTransfers(c.Context(), &walletrpc.GetTransfersRequest{
		In:      true,
		Out:     true,
		Pending: true,
		Failed:  true,
		Pool:    true,
		// TODO: replace with Uint64 method
		AccountIndex: uint64(accountID.Uint()),
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	var res = ListTransactionsResponse{
		Incoming: make([]ListTransactionsTransactionResponse, 0, len(resp.In)),
		Outgoing: make([]ListTransactionsTransactionResponse, 0, len(resp.Out)),
		Pending:  make([]ListTransactionsTransactionResponse, 0, len(resp.Pending)),
		Failed:   make([]ListTransactionsTransactionResponse, 0, len(resp.Failed)),
		Pool:     make([]ListTransactionsTransactionResponse, 0, len(resp.Pool)),
	}
	for _, tx := range resp.In {
		res.Incoming = append(res.Incoming, ListTransactionsTransactionResponse{
			Address:       tx.Address,
			Amount:        tx.Amount,
			Confirmations: tx.Confirmations,
			Fee:           tx.Fee,
			Height:        tx.Height,
			Timestamp:     tx.Timestamp,
			TxID:          tx.Txid,
			UnlockTime:    tx.UnlockTime,
		})
	}
	for _, tx := range resp.Out {
		res.Outgoing = append(res.Outgoing, ListTransactionsTransactionResponse{
			Address:       tx.Address,
			Amount:        tx.Amount,
			Confirmations: tx.Confirmations,
			Fee:           tx.Fee,
			Height:        tx.Height,
			Timestamp:     tx.Timestamp,
			TxID:          tx.Txid,
			UnlockTime:    tx.UnlockTime,
		})
	}
	for _, tx := range resp.Pending {
		res.Pending = append(res.Pending, ListTransactionsTransactionResponse{
			Address:       tx.Address,
			Amount:        tx.Amount,
			Confirmations: tx.Confirmations,
			Fee:           tx.Fee,
			Height:        tx.Height,
			Timestamp:     tx.Timestamp,
			TxID:          tx.Txid,
			UnlockTime:    tx.UnlockTime,
		})
	}
	for _, tx := range resp.Failed {
		res.Failed = append(res.Failed, ListTransactionsTransactionResponse{
			Address:       tx.Address,
			Amount:        tx.Amount,
			Confirmations: tx.Confirmations,
			Fee:           tx.Fee,
			Height:        tx.Height,
			Timestamp:     tx.Timestamp,
			TxID:          tx.Txid,
			UnlockTime:    tx.UnlockTime,
		})
	}
	for _, tx := range resp.Pool {
		res.Pool = append(res.Pool, ListTransactionsTransactionResponse{
			Address:       tx.Address,
			Amount:        tx.Amount,
			Confirmations: tx.Confirmations,
			Fee:           tx.Fee,
			Height:        tx.Height,
			Timestamp:     tx.Timestamp,
			TxID:          tx.Txid,
			UnlockTime:    tx.UnlockTime,
		})
	}

	return c.Result(res)
}
