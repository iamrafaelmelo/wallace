FROM golang:1.25-alpine

RUN set -ex \
    && apk upgrade && apk update \
    && apk add git make \
    && git clone https://github.com/pressly/goose \
    && cd goose \
    && go mod tidy \
    && go build \
        -ldflags="-s -w" \
        -tags='no_clickhouse no_mssql no_mysql no_sqlite3 no_vertica  no_ydb' \
        -o goose ./cmd/goose \
    && ./goose --version \
    && mv ./goose /usr/local/bin

WORKDIR /opt/www/app

ENTRYPOINT ["tail", "-f", "/dev/null"]
