FROM golang:1.25-alpine

RUN set -ex \
    && apk upgrade && apk update \
    && apk add make

WORKDIR /opt/www/app

ENTRYPOINT ["tail", "-f", "/dev/null"]
