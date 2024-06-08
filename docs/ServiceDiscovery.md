# Service discovery

**Problem:** A client connects to a NATS server and does not know how many or if any Walletudo servers are
connected.

**Solution:** By "broadcasting" NATS request messages to wildcard subject `wallets.*.rpc` a client can
discover all currently connected Walletudo servers. The message must carry valid JSON-RPC request.

## JSON-RPC Request

### $rpc.discover

Service discovery method. This is inspired by OpenRPC but not compliant with the specification.

#### Parameters

Empty Object.

#### Returns

Object

| Name   | Type   | Description                            |
|--------|--------|----------------------------------------|
| `info` | `Info` | Object containing service information. |

Info

| Name          | Type     | Description                     |
|---------------|----------|---------------------------------|
| `title`       | `string` | The title of the application.   |
| `description` | `string` | Description of the application. |
| `version`     | `string` | Version of the application.     |

#### Examples

TODO