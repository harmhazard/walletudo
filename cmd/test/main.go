package main

import (
	"context"
	"encoding/json"
	"github.com/dnbsd/jsonrpc"
	"github.com/dnbsd/xmrmux/services/lightwallet/modules"
	"gitlab.com/moneropay/go-monero/walletrpc"
	"os"
	"strconv"
	"strings"
)

func dumpObject(v any) string {
	b, err := json.MarshalIndent(v, "", strings.Repeat(" ", 4))
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	client := walletrpc.New(walletrpc.Config{
		Address: "http://127.0.0.1:18089/json_rpc",
	})
	b := jsonrpc.Builder{
		Modules: modules.New(client),
	}
	s, err := b.Build()
	if err != nil {
		panic(err)
	}

	var req jsonrpc.Request
	switch os.Args[1] {
	case "admin.wallet.init":
		params := jsonrpc.Object{
			"filename": os.Args[2],
		}
		req = jsonrpc.NewRequest("admin.wallet.init", params)

	case "admin.wallet.setDaemon":
		params := jsonrpc.Object{
			"address": os.Args[2],
		}
		req = jsonrpc.NewRequest("admin.wallet.setDaemon", params)

	case "wallet.create":
		params := jsonrpc.Object{
			"label": os.Args[2],
		}
		req = jsonrpc.NewRequest("wallet.create", params)

	case "wallet.list":
		params := jsonrpc.Object{}
		req = jsonrpc.NewRequest("wallet.list", params)

	case "wallet.listAddresses":
		walletID, _ := strconv.Atoi(os.Args[2])
		params := jsonrpc.Object{
			"walletID": walletID,
		}
		req = jsonrpc.NewRequest("wallet.listAddresses", params)

	case "wallet.createAddress":
		walletID, _ := strconv.Atoi(os.Args[2])
		params := jsonrpc.Object{
			"walletID": walletID,
		}
		req = jsonrpc.NewRequest("wallet.createAddress", params)

	default:
		panic("unknown cmd")
	}

	resp := s.Call(context.Background(), req)
	println(dumpObject(req))
	println(dumpObject(resp))
}
