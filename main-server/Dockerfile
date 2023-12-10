FROM golang:1.21 as builder

RUN mkdir -p /build

WORKDIR /build

COPY main-server .

RUN go build -a

# --------------------

FROM alpine as base

COPY --from=builder /build/pocket-trader-backend /usr/bin
COPY main-server.config.toml /usr/bin/config.toml
ENV CONFIG=/usr/bin/config.toml
RUN mkdir -p /usr/bin/db
COPY ./db/stocks.json /usr/bin/db/stocks.json
COPY ./db/predictions.json /usr/bin/db/predictions.json

WORKDIR /usr/bin
ENTRYPOINT ["pocket-trader-backend"]