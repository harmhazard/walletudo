# Walletudo

This is a description of an experimental software that I wish to build in the course of [MoneroKon 2024 Hackathon]().

Monero Wallet Multiplexer is a service implementing a simple JSON-RPC API over NATS.

NOTES:

- how to solve authorization? for instance i want to allow other users access to my wallet, but i also want to allow some other user a readonly access to my wallet.
- how to create new users?
- default wallet account is unused.
- delete wallet can only be implemented as a soft delete, possibly by tagging a an account with a predefined tag.
- invoice API should send notification if the invoice has been paid.

## JSON-RPC API

### Auth API

Should use `Ed25519` aka `EdDSA`. To authenticate the wallet service issues a random challenge, which must be signed by the client with their private key.
If the key is known, the wallet service issues a temporary JWT token signed with service's private key. The JWT token contains an object with the permissions.
The token is submitted as a header in NATS message.

#### Requests

##### auth.new

Method generates a new authentication challenge. The field `challenge` contains a random string which the client must
sign with their private key. The client must call `auth.verify` method and pass the signed challenge as a parameter.

##### auth.verify

Method verifies the signed challenge string.

##### auth.invite

Method generates a temporary token which grants 

##### auth.keys.add
##### auth.keys.modify
##### auth.keys.delete

### Wallet API

#### Requests

##### wallet.init
##### wallet.backup
##### wallet.restore
##### wallet.listAccounts
##### wallet.setDaemon
##### wallet.account.create
##### wallet.account.hide
##### wallet.account.transfer
##### wallet.account.getBalance
##### wallet.account.createAddress
##### wallet.account.listAddresses
##### wallet.account.listTransactions

#### Notifications

##### wallet.transfer

## NATS subjects

## Service discovery

Client should send a JSON-RPC request `$rpc.discover` to a wildcard subject `wallet.*.rpc`.

## Wallet service

```
wallet.{{serviceID}}.rpc
wallet.{{serviceID}}.notifications.wallets.{{walletID}}
```

## Client workflow

1. Connect a wallet server by prompting a user for a server URL, i.e. https://nats.example.com/.
2. Discover wallet services by sending a request `json-rpc:$rpc.discover` to `nats:wallets.*.rpc` subjects.
3. Wait and collect responses from the wallet services and present them to a user on-screen. If there are no wallet services, display an error.
4. Get a list of existing accounts from each wallet service by sending `json-rpc:wallet.listAccounts` requests to `nats:wallet.{id}.rpc` subject.
5. Let a user choose which account from which wallet service to use. If there are no accounts, allow a user to create one.
