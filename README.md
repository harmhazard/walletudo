# Walletudo

> [!CAUTION]
> This is very early stage software. Expect to lose your Monero if you use it!!!

Walletudo is a Monero wallet server. The server provides a simple JSON-RPC protocol over [NATS](https://nats.io)
for most common wallet interactions. The main goal of this software is to provide a reasonably secure backend system to aid the development of lightweight
wallets.

This software was developed during [Monerokon](https://monerokon.org) hackathon that took place between 7th and 9th of
June 2024 in Prague.

## Installation

Build from source:

```bash
$ go install github.com/harmhazard/walletudo@latest
```

## Usage

```bash
$ git clone https://github.com/harmhazard/walletudo
$ cd walletudo/
$ docker-compose up --detach # Be cool and use podman!
$ walletudo --nats-server=wss://user:password@connect.walletudo.com --nats-rpc-subject=wallets.demo1.rpc --nats-discovery-subject "wallets.discover"
```

## Deployment architecture

Clients communicate with Walletudo server through NATS server using [Request-Reply](https://docs.nats.io/nats-concepts/core-nats/reqreply) pattern.
If deployed correctly, the used architecture lowers the attack surface as a computer running the Walletudo server is
on a private network and does not bind any ports. The monero-wallet-rpc, the most sensitive component of the system, is available through the localhost.
The only component that must be addressable is the NATS server itself.

![diagram](https://github.com/harmhazard/walletudo/assets/130958180/3db8f1e0-2ecb-4d85-84d5-8179bfbad11c "Walletudo architecture")

## Documentation

[Service Discovery](docs/ServiceDiscovery.md)
[Wallet JSON-RPC API](docs/WalletAPI.md)
