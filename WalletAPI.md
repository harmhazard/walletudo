# Wallet JSON-RPC API specification

## Methods

### wallet.init

Create a new wallet or open an existing.

> **Note**
Creating a new wallet takes quite some time. It's advisable to set a higher request timeout.

#### Parameters

| Name     | Required | Description                                       |
|----------|----------|---------------------------------------------------|
| filename | yes        | Filename of the wallet that will be created/open. |

#### Example

```shell
$ nats request -s wss://user:password@example.com --reply-timeout 10s "wallets.demo.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.init","params":{"filename":"mywallet"}}'
```

### wallet.backup [NOT IMPLEMENTED]

### wallet.restore [NOT IMPLEMENTED]

### wallet.listAccounts

### wallet.setDaemon

### wallet.account.create

### wallet.account.hide

### wallet.account.transfer

### wallet.account.getBalance

### wallet.account.createAddress

### wallet.account.listAddresses

### wallet.account.listTransactions

## Notifications

### wallet.transfer