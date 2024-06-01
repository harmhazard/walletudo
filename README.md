# Walletudo

Walletudo is a Monero wallet server which can be used as a backend for light wallets or payment processors.
The service provides a simple JSON-RPC protocol over NATS for most common wallet interactions.

## Deployment architecture

Clients communicate with Walletudo server through NATS server using [Request-Reply](https://docs.nats.io/nats-concepts/core-nats/reqreply) pattern.
If deployed correctly, the used architecture lowers the attack surface as a computer running the Walletudo server is
on a private network and does not bind any ports. The monero-wallet-rpc, the most sensitive component of the system, is available through the localhost.
The only component that must be addressable is the NATS server itself.

![diagram](https://private-user-images.githubusercontent.com/130958180/335804743-3db8f1e0-2ecb-4d85-84d5-8179bfbad11c.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTcyNDkzMjcsIm5iZiI6MTcxNzI0OTAyNywicGF0aCI6Ii8xMzA5NTgxODAvMzM1ODA0NzQzLTNkYjhmMWUwLTJlY2ItNGQ4NS04NGQ1LTgxNzliZmJhZDExYy5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjQwNjAxJTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI0MDYwMVQxMzM3MDdaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT0wM2NkMGJlMmEwM2FlODE1OWU5ZTQzZDMyNzc4YzQ2ODdkYjlkY2UxYTYzOTEzNTgwMDJjNDJkNDQ3YjhjZDcwJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZhY3Rvcl9pZD0wJmtleV9pZD0wJnJlcG9faWQ9MCJ9.6sHujbFrfVsBV3rG5ewWnMxJlboiqsCcw_2RvVKRBvw "Walletudo architecture")

## Documentation

[Wallet JSON-RPC API](WalletApi.md)
