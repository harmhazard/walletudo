# Service discovery

**Problem:** A client connects to a NATS server and does not know how many or if any Walletudo servers are
connected.

**Solution:** By "broadcasting" NATS request message to subject `wallets.discover` a client can
discover all currently connected Walletudo servers. The message can be of any value or empty.
