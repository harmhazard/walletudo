package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

// --nats-server
// --nats-rpc-subject
// --wallet-rpc-server
// --log-level
// --version
// --help

type Config struct {
	natsServer      []string
	natsRpcSubject  string
	walletRpcServer string
	logLevel        *slog.Level
}

type Flag struct {
	Name  string
	Usage string
	Func  func(string) error
}

func (c *Config) Flags() []Flag {
	return []Flag{
		{
			Name:  "nats-server",
			Usage: "",
			Func:  c.setNatsServer,
		},
		{
			Name:  "nats-rpc-subject",
			Usage: "",
			Func:  c.setNatsRpcSubject,
		},
		{
			Name:  "wallet-rpc-server",
			Usage: "",
			Func:  c.setWalletRpcServer,
		},
		{
			Name:  "log-level",
			Usage: "",
			Func:  c.setLogLevel,
		},
	}
}

func (c *Config) setNatsServer(s string) error {
	c.natsServer = append(c.natsServer, s)
	return nil
}

func (c *Config) setNatsRpcSubject(s string) error {
	c.natsRpcSubject = s
	return nil
}

func (c *Config) setWalletRpcServer(s string) error {
	c.walletRpcServer = s
	return nil
}

func (c *Config) setLogLevel(s string) error {
	// TODO: parse the string
	v := slog.LevelDebug
	c.logLevel = &v
	return nil
}

func (c *Config) NatsServer() []string {
	defaultValue := []string{"localhost"}
	if len(c.natsServer) == 0 {
		return defaultValue
	}
	return c.natsServer
}

func (c *Config) NatsRpcSubject() string {
	defaultValue := "walletudo.rpc"
	if c.natsRpcSubject == "" {
		return defaultValue
	}
	return c.natsRpcSubject
}

func (c *Config) WalletRpcServer() string {
	defaultValue := "http://localhost:18082/json_rpc"
	if c.walletRpcServer == "" {
		return defaultValue
	}
	return c.walletRpcServer
}

func (c *Config) LogLevel() slog.Level {
	defaultValue := slog.LevelDebug
	if c.logLevel == nil {
		return defaultValue
	}
	return *c.logLevel
}

func (c *Config) Usage() {
	fmt.Printf("Usage: %s <OPTION>\n", os.Args[0])
	for _, flg := range c.Flags() {
		fmt.Printf("  --%s  %s\n", flg.Name, flg.Usage)
	}
	fmt.Printf("  -h, --help  print this help message\n")
}

func NewConfig() *Config {
	c := &Config{}
	for _, flg := range c.Flags() {
		flag.Func(flg.Name, flg.Usage, flg.Func)
	}
	flag.Usage = c.Usage
	flag.Parse()
	return c
}
