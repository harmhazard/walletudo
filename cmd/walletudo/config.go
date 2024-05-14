package main

import "log/slog"

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

func (c *Config) SetNatsServer(s string) error {
	c.natsServer = append(c.natsServer, s)
	return nil
}

func (c *Config) SetNatsRpcSubject(s string) error {
	c.natsRpcSubject = s
	return nil
}

func (c *Config) SetWalletRpcServer(s string) error {
	c.walletRpcServer = s
	return nil
}

func (c *Config) SetLogLevel(s string) error {
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

func NewConfig() *Config {
	return &Config{}
}
