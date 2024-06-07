# Security checklist

* In `wallet.init` path is provided as a parameter, make sure a user cannot write outside the designated directory.
* In `wallet.account.createTransaction` an OpenAlias address can be specified. Make sure this does not leak the IP address of the server.
