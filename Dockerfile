# alpine linux
FROM alpine:3.19

# Prepare server stuff
ENV APP_BIN=lrbooks
ARG SERVER_DIR=/home/.server
WORKDIR $SERVER_DIR
COPY ./${APP_BIN} .

ENV GIN_MODE=release

CMD ./${APP_BIN}