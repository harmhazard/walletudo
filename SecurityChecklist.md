# Security checklist

* In `wallet.account.createTransaction` an OpenAlias address can be specified. Make sure this does not leak the IP address of the server.
* In `wallet.setDaemon` a daemon address can be specified. Warn that this leaks the IP address of the server.
