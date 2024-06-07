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
			{
				Name:    "getBalance",
				Handler: m.GetBalance,
			},
			{
				Name:    "createTransaction",
				Handler: m.CreateTransaction,
			},
		},
	}
}

type module struct {
	client *walletrpc.Client
}

type CreateRequest struct {
	Label string
}

func NewCreateRequest(params jsonrpc.Object) (CreateRequest, error) {
	label, err := params.String("label")
	if err != nil {
		return CreateRequest{}, err
	}

	if label == "" {
		err := jsonrpc.NewErrorParamObjectValue("label", "empty value")
		return CreateRequest{}, err
	}

	return CreateRequest{
		Label: label,
	}, nil
}

type CreateResponse struct {
	AccountID uint64 `json:"accountID"`
}

func (m *module) Create(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	req, err := NewCreateRequest(params)
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.CreateAccount(c.Context(), &walletrpc.CreateAccountRequest{
		Label: req.Label,
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	return c.Result(CreateResponse{
		AccountID: resp.AccountIndex,
	})
}

type ListAddressesRequest struct {
	AccountID uint64
}

func NewListAddressesRequest(params jsonrpc.Object) (ListAddressesRequest, error) {
	accountID, err := params.Number("accountID")
	if err != nil {
		return ListAddressesRequest{}, err
	}

	if accountID.Int() <= 0 {
		err := jsonrpc.NewErrorParamObjectValue("accountID", "must be greater than 0")
		return ListAddressesRequest{}, err
	}

	return ListAddressesRequest{
		// TODO: replace with Uint64 method
		AccountID: uint64(accountID.Uint()),
	}, nil
}

type ListAddressesResponse struct {
	Address string `json:"address"`
	Label   string `json:"label"`
	Used    bool   `json:"used"`
}

func (m *module) ListAddresses(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	req, err := NewListAddressesRequest(params)
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.GetAddress(c.Context(), &walletrpc.GetAddressRequest{
		AccountIndex: req.AccountID,
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	res := make([]ListAddressesResponse, 0, len(resp.Addresses))
	for _, addr := range resp.Addresses {
		res = append(res, ListAddressesResponse{
			Address: addr.Address,
			Label:   addr.Label,
			Used:    addr.Used,
		})
	}

	return c.Result(res)
}

type CreateAddressRequest struct {
	AccountID uint64
}

func NewCreateAddressRequest(params jsonrpc.Object) (CreateAddressRequest, error) {
	accountID, err := params.Number("accountID")
	if err != nil {
		return CreateAddressRequest{}, err
	}

	if accountID.Int() <= 0 {
		err := jsonrpc.NewErrorParamObjectValue("accountID", "must be greater than 0")
		return CreateAddressRequest{}, err
	}

	return CreateAddressRequest{
		// TODO: replace with Uint64 method
		AccountID: uint64(accountID.Uint()),
	}, nil
}

type CreateAddressResponse struct {
	Address string `json:"address"`
}

func (m *module) CreateAddress(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	req, err := NewCreateAddressRequest(params)
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.CreateAddress(c.Context(), &walletrpc.CreateAddressRequest{
		AccountIndex: req.AccountID,
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	return c.Result(CreateAddressResponse{
		Address: resp.Address,
	})
}

type ListTransactionsRequest struct {
	AccountID uint64
}

func NewListTransactionsRequest(params jsonrpc.Object) (ListTransactionsRequest, error) {
	accountID, err := params.Number("accountID")
	if err != nil {
		return ListTransactionsRequest{}, err
	}

	if accountID.Int() <= 0 {
		err := jsonrpc.NewErrorParamObjectValue("accountID", "must be greater than 0")
		return ListTransactionsRequest{}, err
	}

	return ListTransactionsRequest{
		// TODO: replace with Uint64 method
		AccountID: uint64(accountID.Uint()),
	}, nil
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

	req, err := NewListTransactionsRequest(params)
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.GetTransfers(c.Context(), &walletrpc.GetTransfersRequest{
		In:           true,
		Out:          true,
		Pending:      true,
		Failed:       true,
		Pool:         true,
		AccountIndex: req.AccountID,
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

type GetBalanceRequest struct {
	AccountID uint64
}

func NewGetBalanceRequest(params jsonrpc.Object) (GetBalanceRequest, error) {
	accountID, err := params.Number("accountID")
	if err != nil {
		return GetBalanceRequest{}, err
	}

	if accountID.Int() <= 0 {
		err := jsonrpc.NewErrorParamObjectValue("accountID", "must be greater than 0")
		return GetBalanceRequest{}, err
	}

	return GetBalanceRequest{
		// TODO: replace with Uint64 method
		AccountID: uint64(accountID.Uint()),
	}, nil
}

type GetBalanceResponse struct {
	Balance         uint64 `json:"balance"`
	UnlockedBalance uint64 `json:"unlocked_balance"`
}

func (m *module) GetBalance(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	req, err := NewGetBalanceRequest(params)
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.GetBalance(c.Context(), &walletrpc.GetBalanceRequest{
		AccountIndex: req.AccountID,
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	return c.Result(GetBalanceResponse{
		Balance:         resp.Balance,
		UnlockedBalance: resp.UnlockedBalance,
	})
}

type CreateTransactionRequest struct {
	AccountID uint64
	Address   string
	Amount    uint64
}

func NewCreateTransactionRequest(params jsonrpc.Object) (CreateTransactionRequest, error) {
	accountID, err := params.Number("accountID")
	if err != nil {
		return CreateTransactionRequest{}, err
	}

	if accountID.Int() <= 0 {
		err := jsonrpc.NewErrorParamObjectValue("accountID", "must be greater than 0")
		return CreateTransactionRequest{}, err
	}

	address, err := params.String("address")
	if err != nil {
		return CreateTransactionRequest{}, err
	}

	if address == "" {
		err := jsonrpc.NewErrorParamObjectValue("address", "must not be empty")
		return CreateTransactionRequest{}, err
	}

	amount, err := params.Number("amount")
	if err != nil {
		return CreateTransactionRequest{}, err
	}

	if amount.Int() <= 0 {
		err := jsonrpc.NewErrorParamObjectValue("amount", "must be greater than 0")
		return CreateTransactionRequest{}, err
	}

	return CreateTransactionRequest{
		AccountID: uint64(accountID.Uint()),
		Address:   address,
		// TODO: replace with Uint64 method
		Amount: uint64(amount.Uint()),
	}, nil
}

type CreateTransactionResponse struct {
	Address    string `json:"address"`
	Amount     uint64 `json:"amount"`
	Fee        uint64 `json:"fee"`
	TxMetadata string `json:"txMetadata"`
}

func (m *module) CreateTransaction(c *jsonrpc.Context) (any, error) {
	params, err := c.ParamsObject()
	if err != nil {
		return c.Error(err)
	}

	req, err := NewCreateTransactionRequest(params)
	if err != nil {
		return c.Error(err)
	}

	resp, err := m.client.Transfer(c.Context(), &walletrpc.TransferRequest{
		Destinations: []walletrpc.Destination{
			{
				Amount:  req.Amount,
				Address: req.Address,
			},
		},
		AccountIndex:  req.AccountID,
		DoNotRelay:    true,
		GetTxMetadata: true,
	})
	if err != nil {
		// TODO: do not return monero rpc errors to the user!
		return c.Error(err)
	}

	return CreateTransactionResponse{
		Address:    req.Address,
		Amount:     resp.Amount,
		Fee:        resp.Fee,
		TxMetadata: resp.TxMetadata,
	}, nil
}
