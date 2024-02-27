# alpine with golang
FROM golang:1.22.0-alpine3.19

# Prepare server stuff
ENV APP_BIN=lrbooks
ARG SERVER_DIR=/home/.server
WORKDIR $SERVER_DIR
COPY ./${APP_BIN} .

ENV GIN_MODE=release

CMD ./${APP_BIN}