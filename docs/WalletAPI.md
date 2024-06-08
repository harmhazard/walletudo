# Wallet JSON-RPC API specification

## Requests

### wallet.backup [NOT IMPLEMENTED]

### wallet.restore [NOT IMPLEMENTED]

### wallet.listAccounts

List wallet accounts.

#### Parameters

Empty Object

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

### wallet.relayTransaction

Relay previously created transaction to the network.

#### Parameters

Object

| Name         | Type      | Required                                          | Description                        |
|--------------|-----------|---------------------------------------------------|------------------------------------|
| `txMetadata` | `string`  | yes        | Transaction metadata that must be relayed to the network by calling `wallet.relayTransaction`. |

#### Returns

Object

| Name     | Type     | Description       |
|----------|----------|-------------------|
| `txHash` | `string` | Transaction hash. |


#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.relayTransaction","params":{"txMetadata":"01020001020010b7eb8b2bcac8e704c0b281029e188c9f01bf8d09c9b416d27ae30dcfbd03a13ddc4d12b660b969fe07a7cdb0e73d83827e4083676f8cd12bb63bb49f26e082fa3b763dd47f63cc06700200032b0dcfabe9e0088c2c6964b146fc8b25ac1e4a7cf1c6c04595fb7cbb6bcde8bbf0000385617dd7b3ed6156dfa1f8dda2765cda45f96e9e0bedad91b16f6880bc7c0557a12c019fa8e14578612ce2eb48308f51a07a3d499aee026a193a785d5a0dcfdbd61d59020901c0a195c2e1bebb9106e0f2cc0e34d9495d3aff0ae2773ac659b64374f9b843d685b708cbc6df5321ece05e0951364375b71f0762323985c9307ed13a1075c7cd89e066dfd45d900d9531e4f47938eedac2fd4f4fbb73e98c42fb50c7e501c432358d18a68421d0e5f3e8b3ac05da4fd950adcb64bd1e36305dc66a1bc7a7a3bdf9966cc644287e96eee6cdf48d916e01ee3a4d6430d452e19d4122c6bd58b28dfc94749b92def004152b54fd770ee61d8b2679b65b7dbdef28a57eb86592e06fa3d8739b5d5fc13d8512c4bffaacc0f3a73cc378463263c4711619e1b607cb175715e1b83be5c308ffc0c836c7dd13549ed5bf76215b8b94538e3172e50a10f8a3fc25cfad76fed663a9912dd511e1f1f40e9f1a39a02b8f28e0da2ac80d07c2e1439e9c0fedb7a371ecb3acb4e254157bd14410e67f68517bd4c97bb36859534a524517d54f3693c3944fbdea0e74949f6eb3ab4c7b731c41ba14d7e1e7b2e7b7c8585242a4bf58f38fe8d3346a99582ba89d9393270db8e3d8b44c93df4c0c46c5d3375293d3d91962c38284d2a334eb474b7ce2cd74b2c15c6105d6d7120721388e80315972a051365a03d3e25f3e64a08f69f079dc7465346023f3f1b62b548b8b99466477e2e3fc20b8b73f84b4ce97be637a1632aba06eda11369af17ff2e6bcfb0d6b9001f32ca4f33beae26f49e8d9f4a6e6dbe0545cc4936fab800785f0c1b2e5314007b2e919c2e059f32565f143b6171ec04139e19b1b6f8146788a2639fda3e69117e31c0fed457a3b0fbd077d1baca714c987f3e0ee3e497727ab443b00735542f38d91a7af33b1b25d39a360a34460a7bb8434a436cfd01892018497722f4c68eca929c0e499bef740fc881354a95627b3f36a48b0ebc60cc7a370b1f05c6bc4dd11edf5cc2678295059811c5c42fbf61fc0d454f1dbc3ad4c3ff3b7ec319ee2717312343028fc119b90471d2b23962ab0f64b401e753569ecccbdf39ca907413abd965d8db1f45aa35cf1d6628e3551185108d70679f15866fb743699edcbc75e0d2028546238986af22fe5248156f08bd7ce400a1b4d41098d270da48d0d6a327bb68e0f5cd977791e7c568831899915f97a20a7068ba4036909509c0671f7913e8ff23670d4962a46105f8ce05e4cf26a1ecd46b042850ab74f4759a9a01a82aab1d042422e57121526033c7f72f5410609bb4e8b25ec01854d4b5096067fc670331381e9f0cf54e9b24b0a2fc5efb41bbff037c48d730b3c6ebc136262bd02a18b07dee97303364790c1b9e534bc6fc3015895191ef40f1432b056249345ce09c77244a09678c2b8b30040706eeef1e1ebcdb476164b02001567b5720795d5d973c6a76c0e8e5b6ed33a67205bd22b16dd06345852fc0cfcba197c6a0ea429b64935c33ef1bfe573af946134edcc099a01cbdf6ed43604566019075d8f890dc976f6e03ae027b4e851d7d27bdf948d6e8a4cf375cce1061ef326d3a9e2f129bbdd26343f912d36512d8b78b1106072f6d5d0989216790bbf8d0a3ba31310a30377124182646fb28432396f7c8440367a4ac478d186cb03d3c8fbc18959aef6eee811eb2567ca64fe65393272381cc60079762b43e9000e657a22b9dfe46b7798a32c3722ebcff236291938898631c40b109d0b163c1702de37161bdd7dbf973af17df98b199da1c96c4d4d79db3aa530ac1f8bb4ac4d02783141b375062ce5f0b5762e78ffbec47daaa8e9caf236fb3a0c6b8a6b6f3c023e56957a53eb8d027608a23e3e1f236d4ed36e83ba08b0e4a38ff5a2a7e0bf024eceb79a07f419191e34d55bdf2a77391d0de980848f9a531644924289c765e11d9c1bd92582f4526a425c6fd0540e5f8474ae918755ec8a2cde3ac8a391a52d00000000000000006039d301000000000000809a88ac029355b230df54cc5e7fb97877b7aa3961c986355c927a99a84ae049e500f3c46dbb54278caf3fe7690df46b1497fc37f7aaeb42a05f0e105f1c78b6aeeb42b0b501000100433c613763646230653733643833383237653430383336373666386364313262623633626234396632366530383266613362373633646434376636336363303637303e200d41c8fe431660e25268520c456fbc7075af7321814b0f4b095a162b8ccf6e0700015f38326932446e6a784337554c6471675a6a5451764d54324e6832615343386157374155687a6158447244376441323932454a396b4d627a67423862627674416b444c5750693953577264624563686553314863776b4d35344468553454316ea08d0606b840df4556db756415db25cedc0e0836e3f57940a26438a9da10b84a951835ecc0f59b7fed55ea35c68fb7e2780fafb875a754a20bedf300b48ab68541b3700100011002b7eb8b2bc9ce3b6206e0cf6af517f6693842c564d52b2f4f47c6734399ba58a515967ab53418d0e26f6f661895689c219c7eba6a80bff483b2074920fae663285346d92c0281b4f32f4861c5912a2801cdea7eb6c5ea15bec7ec60544ea854c3c2977ad4d26687cc9cc64f0d27f1b7b8615a6ae33ca16364195656bbfedf990da924bf536a6de1819802c1e6f431f779b6d49bac79faaacc902e8413532f831b5b2fc3e1371e12ab5e7c126367e50822eb5d166ee79c428d6a08caa8c78c80ccc17c6dc19f2763299bd98e14a2d202dffef431ccbb2a3a71dca857c9c76981c243135b2d592445bbfdd1ad160f981069f89cf2152fadb725ee803bc2a8b6fc64f7441b9b085c0c9086e44c135566782cfa206c02eb9df6317cc8f6260bf1b8753c03a108fc8fd317cfec67f4eb11ed2974655651e2ef4a4d27b8a92ea703543d41e5edc25e16336ae77c9d4b7a919fadbf953aa9754b335d02aaabff31961eb2ee684ce011377fca44c70be7ef4630b652c7bb150d0667fcf17ab1b14e446eb630be29ef560f146afea32ac296b10561c49322bd3a947ddd0036a9978b02f3df95325d4543f7c6f5fd430d96e3d6fee114fef70a0936eeab44b9c35fa5dd3eee2fe6568c5139ee3586ed96d2aaad02fb61d57dfa6e846854244895a53878d3ae4dc402c5da96326bea777f523245d0941a66b42b9e88ec35acf24fc3fd00d7c9b753b666c56e96c3ac1392084d569e464317f2539e0680cb44d12b75c4c29154ccd5eec4b1cf1702a8e896321917e2bbfa33d73fd17df840bb93c383a06233c464c6463900b215bef24fc3032600ea7eaa70f1da740e7f7870e07d5efae8f6b8c15b5683dca2e994ca1d77e002f7a59a32bafe10f1f9f2830231532e3ca263d28da62a25033fbfacd6abd44eecff835c38691791c2cc5dc228fc55e6eb48f5bff45e4eabe1de7f51a3b2f17df1eed81d160298e39a32d37a5e8f79730586dc52c7bcebb08006193b4ce6797c756be391e92226f3665427ccfa4f9d2c19914498a4805c365c6788cdbf440f6433c2f279ac60a5f0982b02f4b09b32f44cf23ac4b35f7daa710cf00e25b40f4b06ff03640b5f23d288d56da7daf95a10fd886639e43290a2942b721582baf42a14f4b163a3523728941959602061b10286b19b3200331b860da48c67922365b48fd0537d22f9cdc308f8e8959dfa32bd4d6203d7fefa97b0c5234fe78819fee14f2c8cd0acb014991963075b7a974c9ab1c8c3a002bc919c3240fa2633a628767d81cf36c75d73077d90c77f94d6e6045372ba8023e063866308aa70dac798e9a8d17531d5e928240f41062cb307f0376fc5cf03878e84ee3b02f5fa9c32792c0cbd86c043817aefff0cc418587839c5e288f732adb4223e53ef0673ee014c2906bd3d445eee6eb4efdbf0bcfdd222422bc5f1d820b5f1dbd5452d0a7c9302f3829d32f07f8322f8df2be88a271e0d4c4cd1cd9ec499e837f656cb117dc68cc81c1eacc2840ef0349479b152cb2e71ef2ee23a0595fb35da63eb6a3c94d79a811eeee601000000000000001cc36e5c0f7ae5b5093e6c2b36db12db54d65d01a06d367da27607d5686d5f3300000000000000000000cd5627000000000116ebc61a2eb2067748e11fcc7c1285e9c0be61bb5bdd70a90aa3b80cca5ea10f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000809a88ac029355b230df54cc5e7fb97877b7aa3961c986355c927a99a84ae049e500f3c46dbb54278caf3fe7690df46b1497fc37f7aaeb42a05f0e105f1c78b6aeeb42b0b501000200809a88ac029355b230df54cc5e7fb97877b7aa3961c986355c927a99a84ae049e500f3c46dbb54278caf3fe7690df46b1497fc37f7aaeb42a05f0e105f1c78b6aeeb42b0b501005f38326932446e6a784337554c6471675a6a5451764d54324e6832615343386157374155687a6158447244376441323932454a396b4d627a67423862627674416b444c5750693953577264624563686553314863776b4d35344468553454316ea08d0606b840df4556db756415db25cedc0e0836e3f57940a26438a9da10b84a951835ecc0f59b7fed55ea35c68fb7e2780fafb875a754a20bedf300b48ab68541b370010001002c019fa8e14578612ce2eb48308f51a07a3d499aee026a193a785d5a0dcfdbd61d59020901c0a195c2e1bebb91000000000000000003000304015f38326932446e6a784337554c6471675a6a5451764d54324e6832615343386157374155687a6158447244376441323932454a396b4d627a67423862627674416b444c5750693953577264624563686553314863776b4d35344468553454316ea08d0606b840df4556db756415db25cedc0e0836e3f57940a26438a9da10b84a951835ecc0f59b7fed55ea35c68fb7e2780fafb875a754a20bedf300b48ab68541b3700100010000000106000000000000000000000000000000000000000000000000000000000000000000"}}' 
13:00:52 Sending request on "wallets.demo1.rpc"
13:00:53 Received with rtt 715.763102ms
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":{}}
```

### wallet.account.create

Create a new wallet account.

#### Parameters

| Name    | Type    | Required                                          | Description    |
|---------|---------|---------------------------------------------------|----------------|
| `label` | `string` | yes        | Account label. |

#### Returns

Object

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

### wallet.account.createTransaction

Creates a new transaction.

#### Parameters

Object

| Name           | Type            | Required                                          | Description                        |
|----------------|-----------------|---------------------------------------------------|------------------------------------|
| `accountID`    | `integer`       | yes        | Account ID from which to transfer. |
| `address` | `string`  | yes        | Address where to send.        |
| `amount`  | `integer` | yes        | Amount to send. |

#### Returns

Object

| Name         | Type      | Description                                                                                            |
|--------------|-----------|--------------------------------------------------------------------------------------------------------|
| `address`    | `string`  | Destination address.                                                                                   |
| `amount`     | `integer` | Transaction amount.                                                                                    |
| `fee`        | `integer` | Transaction fee.                                                                                       |
| `txMetadata` | `string`  | Transaction metadata that must be relayed to the network by calling `wallet.relayTransaction`. |

#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.account.createTransaction","params":{"accountID": 7, "address":"ADDR", "amount": 10000000}}' 
13:00:52 Sending request on "wallets.demo1.rpc"
13:00:53 Received with rtt 715.763102ms
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":{"address":"82i2DnjxC7ULdqgZjTQvMT2Nh2aSC8aW7AUhzaXDrD7dA292EJ9kMbzgB8bbvtAkDLWPi9SWrdbEcheS1HcwkM54DhU4T1n","amount":100000,"fee":30620000,"txMetadata":"01020001020010b7eb8b2bcac8e704c0b281029e188c9f01bf8d09c9b416d27ae30dcfbd03a13ddc4d12b660b969fe07a7cdb0e73d83827e4083676f8cd12bb63bb49f26e082fa3b763dd47f63cc06700200032b0dcfabe9e0088c2c6964b146fc8b25ac1e4a7cf1c6c04595fb7cbb6bcde8bbf0000385617dd7b3ed6156dfa1f8dda2765cda45f96e9e0bedad91b16f6880bc7c0557a12c019fa8e14578612ce2eb48308f51a07a3d499aee026a193a785d5a0dcfdbd61d59020901c0a195c2e1bebb9106e0f2cc0e34d9495d3aff0ae2773ac659b64374f9b843d685b708cbc6df5321ece05e0951364375b71f0762323985c9307ed13a1075c7cd89e066dfd45d900d9531e4f47938eedac2fd4f4fbb73e98c42fb50c7e501c432358d18a68421d0e5f3e8b3ac05da4fd950adcb64bd1e36305dc66a1bc7a7a3bdf9966cc644287e96eee6cdf48d916e01ee3a4d6430d452e19d4122c6bd58b28dfc94749b92def004152b54fd770ee61d8b2679b65b7dbdef28a57eb86592e06fa3d8739b5d5fc13d8512c4bffaacc0f3a73cc378463263c4711619e1b607cb175715e1b83be5c308ffc0c836c7dd13549ed5bf76215b8b94538e3172e50a10f8a3fc25cfad76fed663a9912dd511e1f1f40e9f1a39a02b8f28e0da2ac80d07c2e1439e9c0fedb7a371ecb3acb4e254157bd14410e67f68517bd4c97bb36859534a524517d54f3693c3944fbdea0e74949f6eb3ab4c7b731c41ba14d7e1e7b2e7b7c8585242a4bf58f38fe8d3346a99582ba89d9393270db8e3d8b44c93df4c0c46c5d3375293d3d91962c38284d2a334eb474b7ce2cd74b2c15c6105d6d7120721388e80315972a051365a03d3e25f3e64a08f69f079dc7465346023f3f1b62b548b8b99466477e2e3fc20b8b73f84b4ce97be637a1632aba06eda11369af17ff2e6bcfb0d6b9001f32ca4f33beae26f49e8d9f4a6e6dbe0545cc4936fab800785f0c1b2e5314007b2e919c2e059f32565f143b6171ec04139e19b1b6f8146788a2639fda3e69117e31c0fed457a3b0fbd077d1baca714c987f3e0ee3e497727ab443b00735542f38d91a7af33b1b25d39a360a34460a7bb8434a436cfd01892018497722f4c68eca929c0e499bef740fc881354a95627b3f36a48b0ebc60cc7a370b1f05c6bc4dd11edf5cc2678295059811c5c42fbf61fc0d454f1dbc3ad4c3ff3b7ec319ee2717312343028fc119b90471d2b23962ab0f64b401e753569ecccbdf39ca907413abd965d8db1f45aa35cf1d6628e3551185108d70679f15866fb743699edcbc75e0d2028546238986af22fe5248156f08bd7ce400a1b4d41098d270da48d0d6a327bb68e0f5cd977791e7c568831899915f97a20a7068ba4036909509c0671f7913e8ff23670d4962a46105f8ce05e4cf26a1ecd46b042850ab74f4759a9a01a82aab1d042422e57121526033c7f72f5410609bb4e8b25ec01854d4b5096067fc670331381e9f0cf54e9b24b0a2fc5efb41bbff037c48d730b3c6ebc136262bd02a18b07dee97303364790c1b9e534bc6fc3015895191ef40f1432b056249345ce09c77244a09678c2b8b30040706eeef1e1ebcdb476164b02001567b5720795d5d973c6a76c0e8e5b6ed33a67205bd22b16dd06345852fc0cfcba197c6a0ea429b64935c33ef1bfe573af946134edcc099a01cbdf6ed43604566019075d8f890dc976f6e03ae027b4e851d7d27bdf948d6e8a4cf375cce1061ef326d3a9e2f129bbdd26343f912d36512d8b78b1106072f6d5d0989216790bbf8d0a3ba31310a30377124182646fb28432396f7c8440367a4ac478d186cb03d3c8fbc18959aef6eee811eb2567ca64fe65393272381cc60079762b43e9000e657a22b9dfe46b7798a32c3722ebcff236291938898631c40b109d0b163c1702de37161bdd7dbf973af17df98b199da1c96c4d4d79db3aa530ac1f8bb4ac4d02783141b375062ce5f0b5762e78ffbec47daaa8e9caf236fb3a0c6b8a6b6f3c023e56957a53eb8d027608a23e3e1f236d4ed36e83ba08b0e4a38ff5a2a7e0bf024eceb79a07f419191e34d55bdf2a77391d0de980848f9a531644924289c765e11d9c1bd92582f4526a425c6fd0540e5f8474ae918755ec8a2cde3ac8a391a52d00000000000000006039d301000000000000809a88ac029355b230df54cc5e7fb97877b7aa3961c986355c927a99a84ae049e500f3c46dbb54278caf3fe7690df46b1497fc37f7aaeb42a05f0e105f1c78b6aeeb42b0b501000100433c613763646230653733643833383237653430383336373666386364313262623633626234396632366530383266613362373633646434376636336363303637303e200d41c8fe431660e25268520c456fbc7075af7321814b0f4b095a162b8ccf6e0700015f38326932446e6a784337554c6471675a6a5451764d54324e6832615343386157374155687a6158447244376441323932454a396b4d627a67423862627674416b444c5750693953577264624563686553314863776b4d35344468553454316ea08d0606b840df4556db756415db25cedc0e0836e3f57940a26438a9da10b84a951835ecc0f59b7fed55ea35c68fb7e2780fafb875a754a20bedf300b48ab68541b3700100011002b7eb8b2bc9ce3b6206e0cf6af517f6693842c564d52b2f4f47c6734399ba58a515967ab53418d0e26f6f661895689c219c7eba6a80bff483b2074920fae663285346d92c0281b4f32f4861c5912a2801cdea7eb6c5ea15bec7ec60544ea854c3c2977ad4d26687cc9cc64f0d27f1b7b8615a6ae33ca16364195656bbfedf990da924bf536a6de1819802c1e6f431f779b6d49bac79faaacc902e8413532f831b5b2fc3e1371e12ab5e7c126367e50822eb5d166ee79c428d6a08caa8c78c80ccc17c6dc19f2763299bd98e14a2d202dffef431ccbb2a3a71dca857c9c76981c243135b2d592445bbfdd1ad160f981069f89cf2152fadb725ee803bc2a8b6fc64f7441b9b085c0c9086e44c135566782cfa206c02eb9df6317cc8f6260bf1b8753c03a108fc8fd317cfec67f4eb11ed2974655651e2ef4a4d27b8a92ea703543d41e5edc25e16336ae77c9d4b7a919fadbf953aa9754b335d02aaabff31961eb2ee684ce011377fca44c70be7ef4630b652c7bb150d0667fcf17ab1b14e446eb630be29ef560f146afea32ac296b10561c49322bd3a947ddd0036a9978b02f3df95325d4543f7c6f5fd430d96e3d6fee114fef70a0936eeab44b9c35fa5dd3eee2fe6568c5139ee3586ed96d2aaad02fb61d57dfa6e846854244895a53878d3ae4dc402c5da96326bea777f523245d0941a66b42b9e88ec35acf24fc3fd00d7c9b753b666c56e96c3ac1392084d569e464317f2539e0680cb44d12b75c4c29154ccd5eec4b1cf1702a8e896321917e2bbfa33d73fd17df840bb93c383a06233c464c6463900b215bef24fc3032600ea7eaa70f1da740e7f7870e07d5efae8f6b8c15b5683dca2e994ca1d77e002f7a59a32bafe10f1f9f2830231532e3ca263d28da62a25033fbfacd6abd44eecff835c38691791c2cc5dc228fc55e6eb48f5bff45e4eabe1de7f51a3b2f17df1eed81d160298e39a32d37a5e8f79730586dc52c7bcebb08006193b4ce6797c756be391e92226f3665427ccfa4f9d2c19914498a4805c365c6788cdbf440f6433c2f279ac60a5f0982b02f4b09b32f44cf23ac4b35f7daa710cf00e25b40f4b06ff03640b5f23d288d56da7daf95a10fd886639e43290a2942b721582baf42a14f4b163a3523728941959602061b10286b19b3200331b860da48c67922365b48fd0537d22f9cdc308f8e8959dfa32bd4d6203d7fefa97b0c5234fe78819fee14f2c8cd0acb014991963075b7a974c9ab1c8c3a002bc919c3240fa2633a628767d81cf36c75d73077d90c77f94d6e6045372ba8023e063866308aa70dac798e9a8d17531d5e928240f41062cb307f0376fc5cf03878e84ee3b02f5fa9c32792c0cbd86c043817aefff0cc418587839c5e288f732adb4223e53ef0673ee014c2906bd3d445eee6eb4efdbf0bcfdd222422bc5f1d820b5f1dbd5452d0a7c9302f3829d32f07f8322f8df2be88a271e0d4c4cd1cd9ec499e837f656cb117dc68cc81c1eacc2840ef0349479b152cb2e71ef2ee23a0595fb35da63eb6a3c94d79a811eeee601000000000000001cc36e5c0f7ae5b5093e6c2b36db12db54d65d01a06d367da27607d5686d5f3300000000000000000000cd5627000000000116ebc61a2eb2067748e11fcc7c1285e9c0be61bb5bdd70a90aa3b80cca5ea10f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000809a88ac029355b230df54cc5e7fb97877b7aa3961c986355c927a99a84ae049e500f3c46dbb54278caf3fe7690df46b1497fc37f7aaeb42a05f0e105f1c78b6aeeb42b0b501000200809a88ac029355b230df54cc5e7fb97877b7aa3961c986355c927a99a84ae049e500f3c46dbb54278caf3fe7690df46b1497fc37f7aaeb42a05f0e105f1c78b6aeeb42b0b501005f38326932446e6a784337554c6471675a6a5451764d54324e6832615343386157374155687a6158447244376441323932454a396b4d627a67423862627674416b444c5750693953577264624563686553314863776b4d35344468553454316ea08d0606b840df4556db756415db25cedc0e0836e3f57940a26438a9da10b84a951835ecc0f59b7fed55ea35c68fb7e2780fafb875a754a20bedf300b48ab68541b370010001002c019fa8e14578612ce2eb48308f51a07a3d499aee026a193a785d5a0dcfdbd61d59020901c0a195c2e1bebb91000000000000000003000304015f38326932446e6a784337554c6471675a6a5451764d54324e6832615343386157374155687a6158447244376441323932454a396b4d627a67423862627674416b444c5750693953577264624563686553314863776b4d35344468553454316ea08d0606b840df4556db756415db25cedc0e0836e3f57940a26438a9da10b84a951835ecc0f59b7fed55ea35c68fb7e2780fafb875a754a20bedf300b48ab68541b3700100010000000106000000000000000000000000000000000000000000000000000000000000000000"}}
```

### wallet.account.getBalance

Get account balance.

#### Parameters

| Name        | Type      | Required                                          | Description                              |
|-------------|-----------|---------------------------------------------------|------------------------------------------|
| `accountID` | `integer` | yes        | Account ID for which to get the balance. |

#### Returns

Object:

| Name               | Type      | Description                                       |
|--------------------|-----------|---------------------------------------------------|
| `balance`          | `integer` | Account balance including any unspendable inputs. |
| `unlocked_balance` | `integer` | Spendable account balance.                        |

#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.account.getBalance","params":{"accountID":7}}' 
13:09:55 Sending request on "wallets.demo1.rpc"
13:09:55 Received with rtt 127.304505ms
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":{"address":"89tnHwwTN3e6ffbG31C9tUdFhWFjZ5SX6XMz59j5BYmxaPfyXb2qanoSqpLeBrvPRSZT2kwCVTU2hentk7y9jBmKSt9HGVu"}}
```

### wallet.account.createAddress

Create a new address associated with a provided account.

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

List addresses for a provided account.

#### Parameters

| Name        | Type      | Required                                          | Description |
|-------------|-----------|---------------------------------------------------|-------------|
| `accountID` | `integer` | yes        | Account ID. |

#### Returns

Array of Objects:

| Name      | Type      | Description                                                          |
|-----------|-----------|----------------------------------------------------------------------|
| `address` | `string`  | Address.                                                             |
| `label`   | `string`  | Address label.                                                       |
| `used`    | `boolean` | Set to `true` if the address has received a transaction. |

#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.account.listAddresses","params":{"accountID":7}}'
13:21:08 Sending request on "wallets.demo1.rpc"
13:21:08 Received with rtt 56.078476ms
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":[{"address":"87LLdzd9mNUCjwCTpv8KVKNXv82nbi5kpJ2ujhdQmBQvDeVQfPbM1v5WdcxgCmhM1rF3C7a8tkLe7CoAH8U7KfHA17VGAak","label":"Test","balance":0,"used":false},{"address":"89tnHwwTN3e6ffbG31C9tUdFhWFjZ5SX6XMz59j5BYmxaPfyXb2qanoSqpLeBrvPRSZT2kwCVTU2hentk7y9jBmKSt9HGVu","label":"","used":false}]}
```

### wallet.account.listTransactions

List transactions for a provided account.

#### Parameters

| Name        | Type      | Required                                          | Description |
|-------------|-----------|---------------------------------------------------|-------------|
| `accountID` | `integer` | yes        | Account ID. |

#### Returns

TODO

#### Example

```shell
$ nats -s "wss://user:password@example.com" request "wallets.demo1.rpc" '{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","method":"wallet.account.listTransactions","params":{"accountID":7}}'
13:31:37 Sending request on "wallets.demo1.rpc"
13:31:37 Received with rtt 84.392311ms
{"jsonrpc":"2.0","id":"zzlDZNRdMTqs","result":{"incoming":[],"outgoing":[],"pending":[],"failed":[],"pool":[]}}
```

## Notifications

### wallet.transfer

TODO
