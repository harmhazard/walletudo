# Monero Wallet Multiplexer

This is a description of an experimental software that I wish to build in the course of [MoneroKon 2024 Hackathon]().

Monero Wallet Multiplexer is a service implementing a simple JSON-RPC API over NATS.

NOTES:

- how to solve authorization? for instance i want to allow other users access to my wallet, but i also want to allow some other user a readonly access to my wallet.
- how to create new users?
- default wallet account is unused.
- delete wallet can only be implemented as a soft delete, possibly by tagging a an account with a predefined tag.
- invoice API should send notification if the invoice has been paid.

## JSON-RPC API

### Light wallet API

#### Requests

##### wallet.init
##### wallet.listAccounts
##### wallet.setDaemon
##### wallet.restore
##### wallet.backup
##### wallet.account.create
##### wallet.account.hide
##### wallet.account.transfer
##### wallet.account.getBalance
##### wallet.account.createAddress
##### wallet.account.listAddresses
##### wallet.account.listTransactions

#### Notifications

##### wallet.create
##### wallet.delete
##### wallet.transfer

## NATS subjects

```
{{serviceID}}.rpc
{{serviceID}}.notifications.wallets.{{walletID}}
```