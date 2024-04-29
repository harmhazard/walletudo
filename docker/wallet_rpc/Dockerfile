FROM debian:bookworm-slim as builder

RUN apt update && apt install wget bzip2 -y

RUN wget -q -O monero.tar.bz2 https://downloads.getmonero.org/cli/monero-linux-x64-v0.18.3.3.tar.bz2
RUN tar --wildcards -xjf monero.tar.bz2 "*/monero-wallet-rpc"

FROM debian:bookworm-slim as runner

COPY --from=builder */monero-wallet-rpc ./

CMD ["/monero-wallet-rpc", "--log-level", "3", "--wallet-dir", "/wallets", "--disable-rpc-login", "--rpc-bind-ip", "0.0.0.0", "--rpc-bind-port", "18082", "--confirm-external-bind", "--log-file", "/dev/null", "--no-initial-sync"]