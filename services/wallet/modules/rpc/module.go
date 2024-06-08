package rpc

import (
	"github.com/dnbsd/jsonrpc"
)

func New() jsonrpc.Module {
	m := module{}
	return jsonrpc.Module{
		Name: "$rpc",
		Methods: []jsonrpc.Method{
			{
				Name:    "discover",
				Handler: m.Discover,
			},
		},
	}
}

type module struct{}

type DiscoverResponse struct {
	Info DiscoverInfoResponse `json:"info"`
}

type DiscoverInfoResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

func (m *module) Discover(c *jsonrpc.Context) (any, error) {
	return c.Result(DiscoverResponse{
		Info: DiscoverInfoResponse{
			Title:       "Walletudo",
			Description: "A Monero wallet server.",
			Version:     "0.0.0",
		},
	})
}
