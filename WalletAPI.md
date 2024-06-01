# Wallet JSON-RPC API specification

## Requests

### wallet.init

Create a new wallet or open an existing.

#### Parameters

| Name        | Type    | Required                                          | Description                                       |
|-------------|---------|---------------------------------------------------|---------------------------------------------------|
| `filename` | `string` | yes        | Filename of the wallet that will be created/open. |

#### Returns

Empty Object.

#### Errors

TODO

#### Example

```shell
$ nats request -s wss://user:password@example.com --reply-timeout 10s "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.init","params":{"filename":"mywallet"}}'
11:37:33 Sending request on "wallets.demo1.rpc"
11:37:37 Received with rtt 4.43974031s
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":{}}
```

### wallet.backup [NOT IMPLEMENTED]

### wallet.restore [NOT IMPLEMENTED]

### wallet.listAccounts

List wallet accounts.

#### Parameters

The method takes no parameters.

#### Returns

Array of Objects:

| Name       | Type                       | Description      |
|-------------|----------------------------|------------------|
| `id`        | `integer`                  | Account ID.      |
| `label`     | `string`                   | Account label.   |
| `balance`   | `integer` | Account balance. |

#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.listAccounts","params":{}}'
11:44:33 Sending request on "wallets.demo1.rpc"
11:44:33 Received with rtt 125.197234ms
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":[{"id":1,"label":"Foo","balance":0},{"id":2,"label":"Bar","balance":0},{"id":3,"label":"Bum","balance":0},{"id":4,"label":"Fourth","balance":0},{"id":5,"label":"W6859","balance":0},{"id":6,"label":"W13874","balance":0}]}
```

### wallet.setDaemon

Connect the wallet to Monero Daemon (monerod) available on provided address.

#### Parameters

| Name      | Type    | Required                                          | Description               |
|-----------|---------|---------------------------------------------------|---------------------------|
| `address` | `string` | yes        | Address of Monero daemon. |

#### Returns

Empty Object.

#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.setDaemon","params":{"address":"http://xmr-node.cakewallet.com:18081"}}'
12:15:54 Sending request on "wallets.demo1.rpc"
12:15:56 Received with rtt 1.575149614s
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":{}}
```

### wallet.account.create

Create a new wallet account.

#### Parameters

| Name    | Type    | Required                                          | Description    |
|---------|---------|---------------------------------------------------|----------------|
| `label` | `string` | yes        | Account label. |

#### Returns

Object:

| Name        | Type                       | Description      |
|-------------|----------------------------|------------------|
| `accountID` | `integer`                  | Account ID.      |

#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.account.create","params":{"label":"Test"}}' 
13:00:52 Sending request on "wallets.demo1.rpc"
13:00:53 Received with rtt 715.763102ms
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":{"accountID":7}}
```

### wallet.account.hide [NOT IMPLEMENTED]

### wallet.account.transfer [NOT IMPLEMENTED]

### wallet.account.getBalance [NOT IMPLEMENTED]

### wallet.account.createAddress

Create a new address associated with a specified account.

#### Parameters

| Name        | Type      | Required                                          | Description                                           |
|-------------|-----------|---------------------------------------------------|-------------------------------------------------------|
| `accountID` | `integer` | yes        | Account ID for which a new address will be generated. |

#### Returns

Object:

| Name      | Type     | Description                      |
|-----------|----------|----------------------------------|
| `address` | `string` | Newly generated account address. |

#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.account.createAddress","params":{"accountID":7}}' 
13:09:55 Sending request on "wallets.demo1.rpc"
13:09:55 Received with rtt 127.304505ms
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":{"address":"89tnHwwTN3e6ffbG31C9tUdFhWFjZ5SX6XMz59j5BYmxaPfyXb2qanoSqpLeBrvPRSZT2kwCVTU2hentk7y9jBmKSt9HGVu"}}
```

### wallet.account.listAddresses

### wallet.account.listTransactions

## Notifications

### wallet.transfer