# Walletudo

Walletudo is a Monero wallet server which can be used as a backend for light wallets or payment processors.
The service provides a simple JSON-RPC protocol over [NATS](https://nats.io) for most common wallet interactions.

This software was developed for [Monerokon](https://monerokon.org) 2024's hackathon that took place on 7th of July 2024 in Prague.

## Deployment architecture

Clients communicate with Walletudo server through NATS server using [Request-Reply](https://docs.nats.io/nats-concepts/core-nats/reqreply) pattern.
If deployed correctly, the used architecture lowers the attack surface as a computer running the Walletudo server is
on a private network and does not bind any ports. The monero-wallet-rpc, the most sensitive component of the system, is available through the localhost.
The only component that must be addressable is the NATS server itself.

![diagram](https://github.com/harmhazard/walletudo/assets/130958180/3db8f1e0-2ecb-4d85-84d5-8179bfbad11c "Walletudo architecture")

## Documentation

[Wallet JSON-RPC API](WalletApi.md)
